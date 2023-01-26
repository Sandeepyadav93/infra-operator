/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	condition "github.com/openstack-k8s-operators/lib-common/modules/common/condition"
)

// OpenStackControlPlane Condition Types used by API objects.
const (
	// OpenStackControlPlaneClientReadyCondition Status=True condition which indicates if OpenStackClient is configured and operational
	OpenStackControlPlaneClientReadyCondition condition.Type = "OpenStackControlPlaneClientReady"

	// OpenStackClientReadyCondition Status=True condition which indicates if OpenStackClient is configured and operational
	OpenStackClientReadyCondition condition.Type = "OpenStackClientReady"
)

// OpenStackControlPlane Reasons used by API objects.
const ()

// Common Messages used by API objects.
const (
	//
	// OpenStackControlPlaneReady condition messages
	//

	// OpenStackControlPlaneReadyErrorMessage
	OpenStackControlPlaneReadyErrorMessage = "OpenStackControlPlane error occured %s"

	// OpenStackControlPlaneRabbitMQReadyInitMessage
	OpenStackControlPlaneRabbitMQReadyInitMessage = "OpenStackControlPlane RabbitMQ not started"

	// OpenStackControlPlaneRabbitMQReadyMessage
	OpenStackControlPlaneRabbitMQReadyMessage = "OpenStackControlPlane RabbitMQ completed"

	// OpenStackControlPlaneRabbitMQReadyRunningMessage
	OpenStackControlPlaneRabbitMQReadyRunningMessage = "OpenStackControlPlane RabbitMQ in progress"

	// OpenStackControlPlaneRabbitMQReadyErrorMessage
	OpenStackControlPlaneRabbitMQReadyErrorMessage = "OpenStackControlPlane RabbitMQ error occured %s"

	// OpenStackControlPlaneMariaDBReadyInitMessage
	OpenStackControlPlaneMariaDBReadyInitMessage = "OpenStackControlPlane MariaDB not started"

	// OpenStackControlPlaneMariaDBReadyMessage
	OpenStackControlPlaneMariaDBReadyMessage = "OpenStackControlPlane MariaDB completed"

	// OpenStackControlPlaneMariaDBReadyRunningMessage
	OpenStackControlPlaneMariaDBReadyRunningMessage = "OpenStackControlPlane MariaDB in progress"

	// OpenStackControlPlaneMariaDBReadyErrorMessage
	OpenStackControlPlaneMariaDBReadyErrorMessage = "OpenStackControlPlane MariaDB error occured %s"

	// OpenStackControlPlaneKeystoneAPIReadyInitMessage
	OpenStackControlPlaneKeystoneAPIReadyInitMessage = "OpenStackControlPlane KeystoneAPI not started"

	// OpenStackControlPlaneKeystoneAPIReadyMessage
	OpenStackControlPlaneKeystoneAPIReadyMessage = "OpenStackControlPlane KeystoneAPI completed"

	// OpenStackControlPlaneKeystoneAPIReadyRunningMessage
	OpenStackControlPlaneKeystoneAPIReadyRunningMessage = "OpenStackControlPlane KeystoneAPI in progress"

	// OpenStackControlPlaneKeystoneAPIReadyErrorMessage
	OpenStackControlPlaneKeystoneAPIReadyErrorMessage = "OpenStackControlPlane KeystoneAPI error occured %s"

	// OpenStackControlPlanePlacementAPIReadyInitMessage
	OpenStackControlPlanePlacementAPIReadyInitMessage = "OpenStackControlPlane PlacementAPI not started"

	// OpenStackControlPlanePlacementAPIReadyMessage
	OpenStackControlPlanePlacementAPIReadyMessage = "OpenStackControlPlane PlacementAPI completed"

	// OpenStackControlPlanePlacementAPIReadyRunningMessage
	OpenStackControlPlanePlacementAPIReadyRunningMessage = "OpenStackControlPlane PlacementAPI in progress"

	// OpenStackControlPlanePlacementAPIReadyErrorMessage
	OpenStackControlPlanePlacementAPIReadyErrorMessage = "OpenStackControlPlane PlacementAPI error occured %s"

	// OpenStackControlPlaneGlanceReadyInitMessage
	OpenStackControlPlaneGlanceReadyInitMessage = "OpenStackControlPlane Glance not started"

	// OpenStackControlPlaneGlanceReadyMessage
	OpenStackControlPlaneGlanceReadyMessage = "OpenStackControlPlane Glance completed"

	// OpenStackControlPlaneGlanceReadyRunningMessage
	OpenStackControlPlaneGlanceReadyRunningMessage = "OpenStackControlPlane Glance in progress"

	// OpenStackControlPlaneGlanceReadyErrorMessage
	OpenStackControlPlaneGlanceReadyErrorMessage = "OpenStackControlPlane Glance error occured %s"

	// OpenStackControlPlaneCinderReadyInitMessage
	OpenStackControlPlaneCinderReadyInitMessage = "OpenStackControlPlane Cinder not started"

	// OpenStackControlPlaneCinderReadyMessage
	OpenStackControlPlaneCinderReadyMessage = "OpenStackControlPlane Cinder completed"

	// OpenStackControlPlaneCinderReadyRunningMessage
	OpenStackControlPlaneCinderReadyRunningMessage = "OpenStackControlPlane Cinder in progress"

	// OpenStackControlPlaneCinderReadyErrorMessage
	OpenStackControlPlaneCinderReadyErrorMessage = "OpenStackControlPlane Cinder error occured %s"

	// OpenStackControlPlaneOVNReadyInitMessage
	OpenStackControlPlaneOVNReadyInitMessage = "OpenStackControlPlane OVN not started"

	// OpenStackControlPlaneOVNReadyMessage
	OpenStackControlPlaneOVNReadyMessage = "OpenStackControlPlane OVN completed"

	// OpenStackControlPlaneOVNReadyRunningMessage
	OpenStackControlPlaneOVNReadyRunningMessage = "OpenStackControlPlane OVN in progress"

	// OpenStackControlPlaneOVNReadyErrorMessage
	OpenStackControlPlaneOVNReadyErrorMessage = "OpenStackControlPlane OVN error occured %s"

	// OpenStackControlPlaneOVSReadyInitMessage
	OpenStackControlPlaneOVSReadyInitMessage = "OpenStackControlPlane OVS not started"

	// OpenStackControlPlaneOVSReadyMessage
	OpenStackControlPlaneOVSReadyMessage = "OpenStackControlPlane OVS completed"

	// OpenStackControlPlaneOVSReadyRunningMessage
	OpenStackControlPlaneOVSReadyRunningMessage = "OpenStackControlPlane OVS in progress"

	// OpenStackControlPlaneOVSReadyErrorMessage
	OpenStackControlPlaneOVSReadyErrorMessage = "OpenStackControlPlane OVS error occured %s"

	// OpenStackControlPlaneNeutronReadyInitMessage
	OpenStackControlPlaneNeutronReadyInitMessage = "OpenStackControlPlane Neutron not started"

	// OpenStackControlPlaneNeutronReadyMessage
	OpenStackControlPlaneNeutronReadyMessage = "OpenStackControlPlane Neutron completed"

	// OpenStackControlPlaneNeutronReadyRunningMessage
	OpenStackControlPlaneNeutronReadyRunningMessage = "OpenStackControlPlane Neutron in progress"

	// OpenStackControlPlaneNeutronReadyErrorMessage
	OpenStackControlPlaneNeutronReadyErrorMessage = "OpenStackControlPlane Neutron error occured %s"

	// OpenStackControlPlaneNovaReadyInitMessage
	OpenStackControlPlaneNovaReadyInitMessage = "OpenStackControlPlane Nova not started"

	// OpenStackControlPlaneNovaReadyMessage
	OpenStackControlPlaneNovaReadyMessage = "OpenStackControlPlane Nova completed"

	// OpenStackControlPlaneNovaReadyRunningMessage
	OpenStackControlPlaneNovaReadyRunningMessage = "OpenStackControlPlane Nova in progress"

	// OpenStackControlPlaneNovaReadyErrorMessage
	OpenStackControlPlaneNovaReadyErrorMessage = "OpenStackControlPlane Nova error occured %s"

	// OpenStackControlPlaneClientReadyInitMessage
	OpenStackControlPlaneClientReadyInitMessage = "OpenStackControlPlane Client not started"

	// OpenStackControlPlaneClientReadyMessage
	OpenStackControlPlaneClientReadyMessage = "OpenStackControlPlane Client completed"

	// OpenStackControlPlaneClientReadyRunningMessage
	OpenStackControlPlaneClientReadyRunningMessage = "OpenStackControlPlane Client in progress"

	// OpenStackControlPlaneClientReadyErrorMessage
	OpenStackControlPlaneClientReadyErrorMessage = "OpenStackControlPlane Client error occured %s"

	// OpenStackClientReadyInitMessage
	OpenStackClientReadyInitMessage = "OpenStack Client not started, waiting on keystone API"

	// OpenStackClientKeystoneWaitingMessage
	OpenStackClientKeystoneWaitingMessage = "OpenStack Client keystone API not yet ready"

	// OpenStackClientConfigMapWaitingMessage
	OpenStackClientConfigMapWaitingMessage = "OpenStack Client waiting for keystone configmap"

	// OpenStackClientSecretWaitingMessage
	OpenStackClientSecretWaitingMessage = "OpenStack Client waiting for secret"

	// OpenStackClientInputReady
	OpenStackClientInputReady = "OpenStack Client input ready"

	// OpenStackClientReadyMessage
	OpenStackClientReadyMessage = "OpenStack Client created"

	// OpenStackClientReadyErrorMessage
	OpenStackClientReadyErrorMessage = "OpenStack Client error occured %s"
)
