// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package powerplatform

import original "github.com/Azure/azure-sdk-for-go/services/preview/powerplatform/mgmt/2020-10-30/powerplatform"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	Creating  PrivateEndpointConnectionProvisioningState = original.Creating
	Deleting  PrivateEndpointConnectionProvisioningState = original.Deleting
	Failed    PrivateEndpointConnectionProvisioningState = original.Failed
	Succeeded PrivateEndpointConnectionProvisioningState = original.Succeeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	Approved PrivateEndpointServiceConnectionStatus = original.Approved
	Pending  PrivateEndpointServiceConnectionStatus = original.Pending
	Rejected PrivateEndpointServiceConnectionStatus = original.Rejected
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None           ResourceIdentityType = original.None
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type Status = original.Status

const (
	Disabled      Status = original.Disabled
	Enabled       Status = original.Enabled
	NotConfigured Status = original.NotConfigured
)

type BaseClient = original.BaseClient
type EnterprisePoliciesClient = original.EnterprisePoliciesClient
type EnterprisePolicy = original.EnterprisePolicy
type EnterprisePolicyIdentity = original.EnterprisePolicyIdentity
type EnterprisePolicyList = original.EnterprisePolicyList
type ErrorResponse = original.ErrorResponse
type ErrorResponseBody = original.ErrorResponseBody
type KeyProperties = original.KeyProperties
type KeyVaultProperties = original.KeyVaultProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type Properties = original.Properties
type PropertiesEncryption = original.PropertiesEncryption
type PropertiesLockbox = original.PropertiesLockbox
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type Subnet = original.Subnet
type SubnetEndpointProperty = original.SubnetEndpointProperty
type SubnetListResult = original.SubnetListResult
type SubnetProperties = original.SubnetProperties
type SubnetsClient = original.SubnetsClient
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEnterprisePoliciesClient(subscriptionID string) EnterprisePoliciesClient {
	return original.NewEnterprisePoliciesClient(subscriptionID)
}
func NewEnterprisePoliciesClientWithBaseURI(baseURI string, subscriptionID string) EnterprisePoliciesClient {
	return original.NewEnterprisePoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubnetsClient(subscriptionID string) SubnetsClient {
	return original.NewSubnetsClient(subscriptionID)
}
func NewSubnetsClientWithBaseURI(baseURI string, subscriptionID string) SubnetsClient {
	return original.NewSubnetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
