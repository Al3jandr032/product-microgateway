/*
 * Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.wso2.choreo.connect.tests.testcases.withapim.apikey;

import com.github.dockerjava.zerodep.shaded.org.apache.hc.core5.http.HttpStatus;
import org.json.JSONArray;
import org.json.JSONObject;
import org.testng.Assert;
import org.testng.annotations.BeforeClass;
import org.testng.annotations.Test;
import org.wso2.am.integration.clients.store.api.v1.dto.APIKeyDTO;
import org.wso2.am.integration.test.utils.bean.APILifeCycleAction;
import org.wso2.choreo.connect.tests.apim.ApimBaseTest;
import org.wso2.choreo.connect.tests.apim.dto.Application;
import org.wso2.choreo.connect.tests.apim.utils.PublisherUtils;
import org.wso2.choreo.connect.tests.apim.utils.StoreUtils;
import org.wso2.choreo.connect.tests.context.CCTestException;
import org.wso2.choreo.connect.tests.util.*;

import java.util.HashMap;
import java.util.Map;

public class APIKeyBlockedAPITestCase extends ApimBaseTest {

    private static final String SAMPLE_API_NAME = "APIKeyTestAPI";
    private static final String SAMPLE_API_CONTEXT = "apiKey";
    private static final String SAMPLE_API_VERSION = "1.0.0";
    private static final String APP_NAME = "APIKeyTestApp";
    private String applicationId;
    private String apiId;
    private String endpointURL;
    Map<String, String> headers = new HashMap<>();

    @BeforeClass(description = "Initialise the setup for API key App level test case")
    void start() throws Exception {
        super.initWithSuperTenant();

        String targetDir = Utils.getTargetDirPath();
        String filePath = targetDir + ApictlUtils.OPENAPIS_PATH + "api_key_openAPI.yaml";

        JSONArray securityScheme = new JSONArray();
        securityScheme.put("oauth_basic_auth_api_key_mandatory");
        securityScheme.put("api_key");

        JSONObject apiProperties = new JSONObject();
        apiProperties.put("name", SAMPLE_API_NAME);
        apiProperties.put("context", "/" + SAMPLE_API_CONTEXT);
        apiProperties.put("version", SAMPLE_API_VERSION);
        apiProperties.put("provider", user.getUserName());
        apiProperties.put("securityScheme", securityScheme);
        apiId = PublisherUtils.createAPIUsingOAS(apiProperties, filePath, publisherRestClient);

        publisherRestClient.changeAPILifeCycleStatus(apiId, "Publish");

        // creating the application
        Application app = new Application(APP_NAME, TestConstant.APPLICATION_TIER.UNLIMITED);
        applicationId = StoreUtils.createApplication(app, storeRestClient);

        PublisherUtils.createAPIRevisionAndDeploy(apiId, publisherRestClient);

        StoreUtils.subscribeToAPI(apiId, applicationId, TestConstant.SUBSCRIPTION_TIER.UNLIMITED, storeRestClient);

        endpointURL = Utils.getServiceURLHttps(SAMPLE_API_CONTEXT + "/1.0.0/pet/1");

        // Obtain API keys
        APIKeyDTO apiKeyDTO = StoreUtils.generateAPIKey(applicationId, TestConstant.KEY_TYPE_PRODUCTION,
                storeRestClient);
        String apiKey = apiKeyDTO.getApikey();
        headers.put("apikey", apiKey);

        Utils.delay(TestConstant.DEPLOYMENT_WAIT_TIME, "Could not wait till initial setup completion.");
    }

    @Test(description = "Invoke API which has API Key as the Application Level Security")
    public void testAPIKeyForAppLevel() throws Exception {
        HttpResponse response = HttpClientRequest.doGet(Utils.getServiceURLHttps(endpointURL), headers);
        Assert.assertNotNull(response);
        Assert.assertEquals(response.getResponseCode(),
                com.github.dockerjava.zerodep.shaded.org.apache.hc.core5.http.HttpStatus.SC_OK,
                "Response code mismatched");
    }

    @Test(description = "Invoke blocked API and check 700700 error code is received", dependsOnMethods = "testAPIKeyForAppLevel")
    public void testAPIKeyForBlockedStateAPI() throws CCTestException, InterruptedException {
        PublisherUtils.changeLCStateAPI(apiId, APILifeCycleAction.BLOCK.getAction(), publisherRestClient, false);
        Thread.sleep(3000);
        HttpResponse response = HttpsClientRequest.doGet(endpointURL, headers);
        Assert.assertNotNull(response, "Error occurred while invoking the endpoint " + endpointURL + ". HttpResponse");
        Assert.assertEquals(response.getResponseCode(), HttpStatus.SC_SERVICE_UNAVAILABLE,
                "Expected error code is 503, but received the code : " + response.getResponseCode());
        Assert.assertTrue(response.getData().contains("700700") && response.getData().contains("API blocked"),
                "Response message mismatched. Expected the error code 700700, but Response Data: " + response.getData());
    }

    @Test(description = "Re publish the blocked API and invoke API", dependsOnMethods = "testAPIKeyForBlockedStateAPI")
    public void testAPIKeyForRePublishedAPI() throws CCTestException, InterruptedException {
        PublisherUtils.changeLCStateAPI(apiId, APILifeCycleAction.RE_PUBLISH.getAction(), publisherRestClient, false);
        Thread.sleep(3000);
        HttpResponse response = HttpsClientRequest.retryGetRequestUntilDeployed(endpointURL, headers);
        Assert.assertNotNull(response, "Error occurred while invoking the endpoint " + endpointURL + ". HttpResponse");
        Assert.assertEquals(response.getResponseCode(), HttpStatus.SC_SUCCESS,
                "Valid subscription should be able to invoke the associated API");
    }

    @Test(description = "Send a request after deleting the subscription", dependsOnMethods = "testAPIKeyForRePublishedAPI")
    public void testAPIKeyWhenSubscriptionDeleted() throws CCTestException, InterruptedException {
        StoreUtils.removeAllSubscriptionsForAnApp(applicationId, storeRestClient);
        Thread.sleep(3000);
        HttpResponse response = HttpsClientRequest.doGet(endpointURL, headers);
        Assert.assertNotNull(response, "Error occurred while invoking the endpoint " + endpointURL + ". HttpResponse");
        Assert.assertEquals(response.getResponseCode(), HttpStatus.SC_FORBIDDEN,
                "Expected error code is 503, but received the code : " + response.getResponseCode());
    }
}
