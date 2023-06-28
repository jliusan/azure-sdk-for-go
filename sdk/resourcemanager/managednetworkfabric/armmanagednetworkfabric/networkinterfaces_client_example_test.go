//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_Create_MaximumSet_Gen.json
func ExampleNetworkInterfacesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkInterfacesClient().BeginCreate(ctx, "resourceGroupName", "networkDeviceName", "networkInterfaceName", armmanagednetworkfabric.NetworkInterface{
		Properties: &armmanagednetworkfabric.NetworkInterfaceProperties{
			Annotation: to.Ptr("null"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkInterface = armmanagednetworkfabric.NetworkInterface{
	// 	Name: to.Ptr("networkInterfaceName"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkdevices/networkinterfaces"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkDevices/networkDeviceName/networkInterfaces/networkInterfaceName"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
	// 		CreatedBy: to.Ptr("d1bd24c7-b27f-477e-86dd-939e107873d7"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("email@address.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkInterfaceProperties{
	// 		Annotation: to.Ptr("null"),
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
	// 		ConnectedTo: to.Ptr("null"),
	// 		InterfaceType: to.Ptr(armmanagednetworkfabric.InterfaceTypeManagement),
	// 		IPv4Address: to.Ptr("10.2.2.8"),
	// 		IPv6Address: to.Ptr("10:2:0:0::"),
	// 		PhysicalIdentifier: to.Ptr("Ethernet1"),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_Get_MaximumSet_Gen.json
func ExampleNetworkInterfacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkInterfacesClient().Get(ctx, "resourceGroupName", "networkDeviceName", "networkInterfaceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkInterface = armmanagednetworkfabric.NetworkInterface{
	// 	Name: to.Ptr("networkInterfaceName"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkdevices/networkinterfaces"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkDevices/networkDeviceName/networkInterfaces/networkInterfaceName"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
	// 		CreatedBy: to.Ptr("d1bd24c7-b27f-477e-86dd-939e107873d7"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("email@address.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkInterfaceProperties{
	// 		Annotation: to.Ptr("null"),
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
	// 		ConnectedTo: to.Ptr("null"),
	// 		InterfaceType: to.Ptr(armmanagednetworkfabric.InterfaceTypeManagement),
	// 		IPv4Address: to.Ptr("10.2.2.8"),
	// 		IPv6Address: to.Ptr("10:2:0:0::"),
	// 		PhysicalIdentifier: to.Ptr("Ethernet1"),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_Update_MaximumSet_Gen.json
func ExampleNetworkInterfacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkInterfacesClient().BeginUpdate(ctx, "resourceGroupName", "networkDeviceName", "networkInterfaceName", armmanagednetworkfabric.NetworkInterfacePatch{
		Properties: &armmanagednetworkfabric.NetworkInterfacePatchProperties{
			Annotation: to.Ptr("null"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkInterface = armmanagednetworkfabric.NetworkInterface{
	// 	Name: to.Ptr("networkInterfaceName"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkdevices/networkinterfaces"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkDevices/networkDeviceName/networkInterfaces/networkInterfaceName"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
	// 		CreatedBy: to.Ptr("d1bd24c7-b27f-477e-86dd-939e107873d7"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("email@address.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkInterfaceProperties{
	// 		Annotation: to.Ptr("null"),
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
	// 		ConnectedTo: to.Ptr("null"),
	// 		InterfaceType: to.Ptr(armmanagednetworkfabric.InterfaceTypeManagement),
	// 		IPv4Address: to.Ptr("10.2.2.8"),
	// 		IPv6Address: to.Ptr("10:2:0:0::"),
	// 		PhysicalIdentifier: to.Ptr("Ethernet1"),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_Delete_MaximumSet_Gen.json
func ExampleNetworkInterfacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkInterfacesClient().BeginDelete(ctx, "resourceGroupName", "networkDeviceName", "networkInterfaceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_List_MaximumSet_Gen.json
func ExampleNetworkInterfacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkInterfacesClient().NewListPager("resourceGroupName", "networkDeviceName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.NetworkInterfacesList = armmanagednetworkfabric.NetworkInterfacesList{
		// 	Value: []*armmanagednetworkfabric.NetworkInterface{
		// 		{
		// 			Name: to.Ptr("networkInterfaceName"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkdevices/networkinterfaces"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkDevices/networkDeviceName/networkInterfaces/networkInterfaceName"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
		// 				CreatedBy: to.Ptr("d1bd24c7-b27f-477e-86dd-939e107873d7"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T07:58:04.840Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("email@address.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkInterfaceProperties{
		// 				Annotation: to.Ptr("null"),
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
		// 				ConnectedTo: to.Ptr("null"),
		// 				InterfaceType: to.Ptr(armmanagednetworkfabric.InterfaceTypeManagement),
		// 				IPv4Address: to.Ptr("10.2.2.8"),
		// 				IPv6Address: to.Ptr("10:2:0:0::"),
		// 				PhysicalIdentifier: to.Ptr("Ethernet1"),
		// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_getStatus_MaximumSet_Gen.json
func ExampleNetworkInterfacesClient_BeginGetStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkInterfacesClient().BeginGetStatus(ctx, "resourceGroupName", "networkDeviceName", "networkInterfaceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_updateAdministrativeState_MaximumSet_Gen.json
func ExampleNetworkInterfacesClient_BeginUpdateAdministrativeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkInterfacesClient().BeginUpdateAdministrativeState(ctx, "resourceGroupName", "networkDeviceName", "networkInterfaceName", armmanagednetworkfabric.UpdateAdministrativeState{
		State: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnable),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}