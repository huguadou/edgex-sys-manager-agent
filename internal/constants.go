/*******************************************************************************
 * Copyright 2018 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package internal

const (
	ApiConfigRoute       = "/api/v1/config"
	ApiPingRoute         = "/api/v1/ping"
	BootTimeoutDefault   = 30000
	ClientMonitorDefault = 15000
	ConfigFileName       = "configuration.toml"
	ConfigV2Stem         = "config/V2/"
)

const (
	ServiceKeyPrefix                = "edgex-"
	ConfigSeedServiceKey            = "edgex-config-seed"
	CoreCommandServiceKey           = "edgex-core-command"
	CoreDataServiceKey              = "edgex-core-data"
	CoreMetaDataServiceKey          = "edgex-core-metadata"
	ExportClientServiceKey          = "edgex-export-client"
	ExportDistroServiceKey          = "edgex-export-distro"
	SupportLoggingServiceKey        = "edgex-support-logging"
	SupportNotificationsServiceKey  = "edgex-support-notifications"
	SystemManagementAgentServiceKey = "edgex-sys-mgmt-agent"
	SupportSchedulerServiceKey      = "edgex-support-scheduler"
	MockMqttDriver1ServiceKey       = "edgex-mock-mqtt1"
	MockMqttDriver2ServiceKey       = "edgex-mock-mqtt2"
	MockMqttDriver3ServiceKey       = "edgex-mock-mqtt3"
)
