/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 *  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Package messaging holds the implementation for event listeners functions
package messaging

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
	logger "github.com/wso2/adapter/loggers"
)

// Try to connect to the RabbitMQ server as long as it takes to establish a connection
func connectToRabbitMQ(url string) (*amqp.Connection, error) {
	var err error = nil
	var conn *amqp.Connection
	conn, err = amqp.Dial(url)
	if err == nil {
		return conn, nil
	}
	_, conn, err = connectionRetry("")
	return conn, err
}

// reconnect reconnects to server if the connection or a channel
// is closed unexpectedly.
func (c *Consumer) reconnect(key string) {
	var err error
	shouldReconnect := false
	connClose := <-c.conn.NotifyClose(make(chan *amqp.Error))
	connBlocked := c.conn.NotifyBlocked(make(chan amqp.Blocking))
	chClose := c.channel.NotifyClose(make(chan *amqp.Error))

	if connClose != nil {
		shouldReconnect = true
		logger.LoggerMsg.Errorf("CRITICAL: Connection dropped for %s, reconnecting...", key)
	}

	if connBlocked != nil {
		shouldReconnect = true
		logger.LoggerMsg.Errorf("CRITICAL: Connection blocked for %s, reconnecting...", key)
	}

	if chClose != nil {
		shouldReconnect = true
		logger.LoggerMsg.Errorf("CRITICAL: Channel closed for %s, reconnecting...", key)
	}

	if shouldReconnect {
		c.conn.Close()
		c, rabbitConn, err = connectionRetry(key)
		if err != nil {
			logger.LoggerMsg.Errorf("Cannot establish connection for topic %s", key)
		}
	} else {
		logger.LoggerMsg.Infof("NotifyClose from the connection and channel are %v and %v respectively, NotifyBlocked from the connection is %v",
			connClose, chClose, connBlocked)
	}
}

// connectionRetry
func connectionRetry(key string) (*Consumer, *amqp.Connection, error) {
	var err error = nil
	var i int

	for j := 0; j < len(amqpURIArray); j++ {
		var maxAttempt int = amqpURIArray[j].retryCount
		var retryInterval time.Duration = time.Duration(amqpURIArray[j].connectionDelay) * time.Second

		if maxAttempt == 0 {
			maxAttempt = 5
		}

		if retryInterval == 0 {
			retryInterval = 10 * time.Second
		}
		logger.LoggerMsg.Infof("Retrying to connect with %s in every %d seconds until exceed %d attempts",
			maskURL(amqpURIArray[j].url), amqpURIArray[j].connectionDelay, maxAttempt)

		for i = 1; i <= maxAttempt; i++ {

			rabbitConn, err = amqp.Dial(amqpURIArray[j].url + "/")
			if err == nil {
				if key != "" && len(key) > 0 {
					logger.LoggerMsg.Infof("Reconnected to topic %s", key)
					// startup pull
					c, err := StartConsumer(key)
					return c, rabbitConn, err
				}
				return nil, rabbitConn, nil
			}

			if key != "" && len(key) > 0 {
				logger.LoggerMsg.Infof("Retry attempt %d for the %s to connect with topic %s has failed. Retrying after %d seconds", i,
					maskURL(amqpURIArray[j].url), key, amqpURIArray[j].connectionDelay)
			} else {
				logger.LoggerMsg.Infof("Retry attempt %d for the %s has failed. Retrying after %d seconds", i, maskURL(amqpURIArray[j].url), amqpURIArray[j].connectionDelay)
			}
			time.Sleep(retryInterval)
		}
		if i >= maxAttempt {
			logger.LoggerMsg.Infof("Exceeds maximum connection retry attempts %d for %s", maxAttempt, maskURL(amqpURIArray[j].url))
			return retryExponentially(key, amqpURIArray[j].url, retryInterval)
		}
	}
	return nil, rabbitConn, err
}

func retryExponentially(key string, url string, retryInterval time.Duration) (*Consumer, *amqp.Connection, error) {
	logger.LoggerMsg.Infof("Trying to connect exponentially for %s", maskURL(url))
	var err error = nil
	initInterval := int(retryInterval.Seconds())
	maxInterval := initInterval * 10
	interval := initInterval
	count := 0
	for {
		rabbitConn, err = amqp.Dial(url + "/")
		if err != nil {
			if interval < maxInterval {
				interval = initInterval + (initInterval * count)
				count = count + 1
			}
			if key != "" && len(key) > 0 {
				logger.LoggerMsg.Infof("Retry attempt for the %s to connect with topic %s has failed. Retrying after %d seconds", maskURL(url), key, interval)
			} else {
				logger.LoggerMsg.Infof("Retry attempt for the %s has failed. Retrying after %d seconds", maskURL(url), interval)
			}

			time.Sleep(time.Duration(interval) * time.Second)
		} else {
			if key != "" && len(key) > 0 {
				logger.LoggerMsg.Infof("Reconnected to topic %s", key)
				// startup pull
				c, err := StartConsumer(key)
				return c, rabbitConn, err
			}
		}
	}
}

//extract AMQPURLList from EventListening connection url
func retrieveAMQPURLList() []amqpFailoverURL {
	var connectionURLList []string
	connectionURLList = mgwConfig.ControlPlane.JmsConnectionParameters.EventListeningEndpoints

	amqlURLList := []amqpFailoverURL{}

	for _, conURL := range connectionURLList {
		var delay int = 0
		var retries int = 0
		amqpConnectionURL := strings.Split(conURL, "?")[0]
		u, err := url.Parse(conURL)
		if err != nil {
			logger.LoggerMsg.Errorf("Error occured %s", maskURL(err.Error()))
		} else {
			m, _ := url.ParseQuery(u.RawQuery)
			if m["connectdelay"] != nil {
				connectdelay := m["connectdelay"][0]
				delay, _ = strconv.Atoi(connectdelay[1 : len(connectdelay)-1])
			}

			if m["retries"] != nil {
				retrycount := m["retries"][0]
				retries, _ = strconv.Atoi(retrycount[1 : len(retrycount)-1])
			}

			failoverurlObj := amqpFailoverURL{url: amqpConnectionURL, retryCount: retries,
				connectionDelay: delay}
			amqlURLList = append(amqlURLList, failoverurlObj)
		}
	}
	return amqlURLList
}

func maskURL(url string) string {
	pattern := regexp.MustCompile(`\/\/([a-zA-Z].*)@\b`)
	matches := pattern.FindStringSubmatch(url)
	if len(matches) > 1 {
		return strings.ReplaceAll(url, matches[1], "******")
	}
	return url
}

type amqpFailoverURL struct {
	url             string
	retryCount      int
	connectionDelay int
}
