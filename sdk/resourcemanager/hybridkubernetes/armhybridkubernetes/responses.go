// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

// ConnectedClusterClientCreateOrReplaceResponse contains the response from method ConnectedClusterClient.BeginCreateOrReplace.
type ConnectedClusterClientCreateOrReplaceResponse struct {
	// Represents a connected cluster.
	ConnectedCluster
}

// ConnectedClusterClientDeleteResponse contains the response from method ConnectedClusterClient.BeginDelete.
type ConnectedClusterClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedClusterClientGetResponse contains the response from method ConnectedClusterClient.Get.
type ConnectedClusterClientGetResponse struct {
	// Represents a connected cluster.
	ConnectedCluster
}

// ConnectedClusterClientListByResourceGroupResponse contains the response from method ConnectedClusterClient.NewListByResourceGroupPager.
type ConnectedClusterClientListByResourceGroupResponse struct {
	// The paginated list of connected Clusters
	ConnectedClusterList
}

// ConnectedClusterClientListBySubscriptionResponse contains the response from method ConnectedClusterClient.NewListBySubscriptionPager.
type ConnectedClusterClientListBySubscriptionResponse struct {
	// The paginated list of connected Clusters
	ConnectedClusterList
}

// ConnectedClusterClientListClusterUserCredentialResponse contains the response from method ConnectedClusterClient.ListClusterUserCredential.
type ConnectedClusterClientListClusterUserCredentialResponse struct {
	// The list of credential result response.
	CredentialResults
}

// ConnectedClusterClientUpdateResponse contains the response from method ConnectedClusterClient.Update.
type ConnectedClusterClientUpdateResponse struct {
	// Represents a connected cluster.
	ConnectedCluster
}

// OperationsClientGetResponse contains the response from method OperationsClient.NewGetPager.
type OperationsClientGetResponse struct {
	// The paginated list of connected cluster API operations.
	OperationList
}
