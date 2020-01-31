/*
 *Copyright 2015 Huawei Technologies Co., Ltd. All rights reserved.
 *	   eSDK is licensed under the Apache License, Version 2.0 (the "License");
 *	   you may not use this file except in compliance with the License.
 *	   You may obtain a copy of the License at
 *
 *	       http://www.apache.org/licenses/LICENSE-2.0
 *
 *
 *	   Unless required by applicable law or agreed to in writing, software
 *	   distributed under the License is distributed on an "AS IS" BASIS,
 *	   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *	   See the License for the specific language governing permissions and
 *	   limitations under the License.
 */
package neutronModules

import (
	"github.com/opentelekomcloud/dockermachinedriver/com/modules"
)

// The response of creating a security group rule
type CreateSecurityGroupRuleResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.SecurityGroupRuleCreateInfo
}

// The response of querying a security group rule
type ShowSecurityGroupRuleResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.SecurityGroupRuleCreateInfo
}

// The response of deleting a security group rule
type DeleteSecurityGroupRuleResp struct {
	ResponseCode int
	modules.ErrorInfo
}
