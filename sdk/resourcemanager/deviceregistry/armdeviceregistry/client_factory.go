//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceregistry

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAssetEndpointProfilesClient creates a new instance of AssetEndpointProfilesClient.
func (c *ClientFactory) NewAssetEndpointProfilesClient() *AssetEndpointProfilesClient {
	return &AssetEndpointProfilesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAssetsClient creates a new instance of AssetsClient.
func (c *ClientFactory) NewAssetsClient() *AssetsClient {
	return &AssetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBillingContainersClient creates a new instance of BillingContainersClient.
func (c *ClientFactory) NewBillingContainersClient() *BillingContainersClient {
	return &BillingContainersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationStatusClient creates a new instance of OperationStatusClient.
func (c *ClientFactory) NewOperationStatusClient() *OperationStatusClient {
	return &OperationStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}
