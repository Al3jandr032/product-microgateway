// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

import ballerina/http;
import ballerina/log;


int bal_metric_port = getConfigIntValue(MICRO_GATEWAY_METRICS_PORTS, PORT, 9797);
int jmx_metric_port = getConfigIntValue(MICRO_GATEWAY_METRICS_PORTS, JMX_PORT, 8080);

http:Client balMetricEndpoint = new ("http://localhost:" + bal_metric_port.toString());
http:Client jmxMetricEndpoint = new ("http://localhost:" + jmx_metric_port.toString());


service metric =
@http:ServiceConfig {
    basePath: "/*",
    auth: {
        scopes: ["observability"]
    }
}
service {

    @http:ResourceConfig {
        path: "/balMetric"
    }

    resource function balMetric(http:Caller caller, http:Request req) returns error? {

        var bal_response = balMetricEndpoint->forward("/metrics", <@untainted>req);

        if (bal_response is http:Response) {
            var result = caller->respond(bal_response);
            if (result is error) {
                log:printError("Error sending bal response http:responce", err = result);
            }
        } else {

            http:Response res = new;
            res.statusCode = 500;
            res.setPayload(<string>bal_response.detail()?.message);
            var result = caller->respond(res);
            if (result is error) {
                log:printError("Error sending bal response else http:responce", err = result);
            }
        }
    }

    resource function jmxMetric(http:Caller caller, http:Request req) returns error? {

        var jmx_response = jmxMetricEndpoint->forward("/metrics", <@untainted>req);

        if (jmx_response is http:Response) {
            var result = caller->respond(jmx_response);
            if (result is error) {
                log:printError("Error sending jmx response http:responce", err = result);
            }
        } else {

            http:Response res = new;
            res.statusCode = 500;
            res.setPayload(<string>jmx_response.detail()?.message);
            var result = caller->respond(res);
            if (result is error) {
                log:printError("Error sending jmx response else http:responce", err = result);
            }
        }
    }

};

public function startObservabilityMetrics() {
    if (isMetricsEnabled) {
        ObservabilityMetricListener observabilityMetricListner = new;
        error? err = observabilityMetricListner.__attach(metric, ());
        error? err1 = observabilityMetricListner.__start();
    }
}
