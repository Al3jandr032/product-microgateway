/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  WSO2 Inc. licenses this file to you under the Apache License,
 *  Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing,
 *  software distributed under the License is distributed on an
 *  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 *  KIND, either express or implied.  See the License for the
 *  specific language governing permissions and limitations
 *  under the License.
 */
package org.wso2.micro.gateway.tests.faultResponses;

import io.netty.handler.codec.http.HttpHeaderNames;
import org.json.JSONObject;
import org.testng.Assert;
import org.testng.annotations.AfterClass;
import org.testng.annotations.BeforeClass;
import org.testng.annotations.Test;
import org.wso2.micro.gateway.tests.common.BaseTestCase;
import org.wso2.micro.gateway.tests.common.ResponseConstants;
import org.wso2.micro.gateway.tests.common.model.ApplicationDTO;
import org.wso2.micro.gateway.tests.util.HttpClientRequest;
import org.wso2.micro.gateway.tests.util.TestConstant;
import org.wso2.micro.gateway.tests.util.TokenUtil;

import java.util.HashMap;
import java.util.Map;

public class FaultResponsesWithAllPathsTestCase extends BaseTestCase {
    protected String jwtTokenProd;

    @BeforeClass
    public void start() throws Exception {
        String project = "faultResponsesWithAllPathsProject";
        // Define application info
        ApplicationDTO application = new ApplicationDTO();
        application.setName("faultResponsesWithAllPathsApp");
        application.setTier("Unlimited");
        application.setId((int) (Math.random() * 1000));

        jwtTokenProd = TokenUtil.getBasicJWT(application, new JSONObject(), TestConstant.KEY_TYPE_PRODUCTION, 3600);

        // generate apis with CLI and start the micro gateway server
        super.init(project, new String[]{"faultResponses/fault_responses.yaml"}, null, null,
                "confs/fault-responses-cli-test-config.toml");
    }

    @Test(description = "Test the normal API invocation when fault responses are enabled")
    public void testAPIInvocationWithFaultResponses() throws Exception {
        Map<String, String> headers = new HashMap<>();
        //test endpoint with token
        headers.put(HttpHeaderNames.AUTHORIZATION.toString(), "Bearer " + jwtTokenProd);
        org.wso2.micro.gateway.tests.util.HttpResponse response = HttpClientRequest
                .doGet(getServiceURLHttp("abc"), headers);
        Assert.assertNotNull(response);
        Assert.assertEquals(response.getData(), ResponseConstants.petByIdResponseV1);
        Assert.assertEquals(response.getResponseCode(), 200, "Response code mismatched");
    }

    @Test(description = "Test the normal API invocation when fault responses are enabled")
    public void testMethodNotAllowed() throws Exception {
        Map<String, String> headers = new HashMap<>();
        org.wso2.micro.gateway.tests.util.HttpResponse response = HttpClientRequest
                .doPost(getServiceURLHttp("abc"), "{}", headers);
        Assert.assertNotNull(response);
        Assert.assertEquals(response.getHeaders().get("Allow"), "GET");
        Assert.assertEquals(response.getData(), ResponseConstants.METHOD_NOT_ALLOWED_RESPONSE);
        Assert.assertEquals(response.getResponseCode(), 405, "Response code mismatched");
    }

    @AfterClass
    public void stop() throws Exception {
        //Stop all the mock servers
        super.finalize();
    }
}
