//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/BastionShareableLinkCreate.json
func ExampleManagementClient_BeginPutBastionShareableLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginPutBastionShareableLink(ctx, "rg1", "bastionhosttenant", armnetwork.BastionShareableLinkListRequest{
		VMs: []*armnetwork.BastionShareableLink{
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
			},
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	for res.More() {
		page, err := res.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.BastionShareableLinkListResult = armnetwork.BastionShareableLinkListResult{
		// 	Value: []*armnetwork.BastionShareableLink{
		// 		{
		// 			Bsl: to.Ptr("http://bst-bastionhostid.bastion.com/api/shareable-url/tokenvm1"),
		// 			CreatedAt: to.Ptr("2019-10-18T12:00:00.0000Z"),
		// 			VM: &armnetwork.VM{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 			},
		// 		},
		// 		{
		// 			Bsl: to.Ptr("http://bst-bastionhostid.bastion.com/api/shareable-url/tokenvm2"),
		// 			CreatedAt: to.Ptr("2019-10-17T12:00:00.0000Z"),
		// 			VM: &armnetwork.VM{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/BastionShareableLinkDelete.json
func ExampleManagementClient_BeginDeleteBastionShareableLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginDeleteBastionShareableLink(ctx, "rg1", "bastionhosttenant", armnetwork.BastionShareableLinkListRequest{
		VMs: []*armnetwork.BastionShareableLink{
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
			},
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/BastionShareableLinkGet.json
func ExampleManagementClient_NewGetBastionShareableLinkPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewGetBastionShareableLinkPager("rg1", "bastionhosttenant", armnetwork.BastionShareableLinkListRequest{
		VMs: []*armnetwork.BastionShareableLink{
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
			},
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
				},
			}},
	}, nil)
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
		// page.BastionShareableLinkListResult = armnetwork.BastionShareableLinkListResult{
		// 	Value: []*armnetwork.BastionShareableLink{
		// 		{
		// 			Bsl: to.Ptr("http://bst-bastionhostid.bastion.com/api/shareable-url/tokenvm1"),
		// 			CreatedAt: to.Ptr("2019-10-18T12:00:00.0000Z"),
		// 			VM: &armnetwork.VM{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 			},
		// 		},
		// 		{
		// 			Bsl: to.Ptr("http://bst-bastionhostid.bastion.com/api/shareable-url/tokenvm2"),
		// 			CreatedAt: to.Ptr("2019-10-17T12:00:00.0000Z"),
		// 			VM: &armnetwork.VM{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/BastionSessionsList.json
func ExampleManagementClient_BeginGetActiveSessions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginGetActiveSessions(ctx, "rg1", "bastionhosttenant", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	for res.More() {
		page, err := res.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.BastionActiveSessionListResult = armnetwork.BastionActiveSessionListResult{
		// 	Value: []*armnetwork.BastionActiveSession{
		// 		{
		// 			ResourceType: to.Ptr("VM"),
		// 			SessionDurationInMins: to.Ptr[float32](0),
		// 			SessionID: to.Ptr("sessionId"),
		// 			StartTime: "2019-1-1T12:00:00.0000Z",
		// 			TargetHostName: to.Ptr("vm01"),
		// 			TargetIPAddress: to.Ptr("1.1.1.1"),
		// 			TargetResourceGroup: to.Ptr("rg1"),
		// 			TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm01"),
		// 			TargetSubscriptionID: to.Ptr("subid"),
		// 			UserName: to.Ptr("user"),
		// 			Protocol: to.Ptr(armnetwork.BastionConnectProtocolSSH),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/BastionSessionDelete.json
func ExampleManagementClient_NewDisconnectActiveSessionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewDisconnectActiveSessionsPager("rg1", "bastionhosttenant", armnetwork.SessionIDs{}, nil)
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
		// page.BastionSessionDeleteResult = armnetwork.BastionSessionDeleteResult{
		// 	Value: []*armnetwork.BastionSessionState{
		// 		{
		// 			Message: to.Ptr("session session1 invalidated!"),
		// 			SessionID: to.Ptr("session1"),
		// 			State: to.Ptr("Disconnected"),
		// 		},
		// 		{
		// 			Message: to.Ptr("session session2 could not be disconnected!"),
		// 			SessionID: to.Ptr("session2"),
		// 			State: to.Ptr("Failed"),
		// 		},
		// 		{
		// 			Message: to.Ptr("session session3 not found!"),
		// 			SessionID: to.Ptr("session3"),
		// 			State: to.Ptr("NotFound"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/CheckDnsNameAvailability.json
func ExampleManagementClient_CheckDNSNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().CheckDNSNameAvailability(ctx, "westus", "testdns", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DNSNameAvailabilityResult = armnetwork.DNSNameAvailabilityResult{
	// 	Available: to.Ptr(false),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/expressRouteProviderPort.json
func ExampleManagementClient_ExpressRouteProviderPort() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ExpressRouteProviderPort(ctx, "abc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteProviderPort = armnetwork.ExpressRouteProviderPort{
	// 	Type: to.Ptr("Microsoft.Network/expressRouteProviderPort"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/ExpressRoutePortsLocations/SiliconValley/bvtazureixpportpair1"),
	// 	Location: to.Ptr("uswest"),
	// 	Etag: to.Ptr("W/\"c0e6477e-8150-4d4f-9bf6-bb10e6acb63a\""),
	// 	Properties: &armnetwork.ExpressRouteProviderPortProperties{
	// 		OverprovisionFactor: to.Ptr[int32](4),
	// 		PeeringLocation: to.Ptr("SiliconValley"),
	// 		PortBandwidthInMbps: to.Ptr[int32](4000),
	// 		PortPairDescriptor: to.Ptr("bvtazureixpportpair1"),
	// 		PrimaryAzurePort: to.Ptr("bvtazureixp01"),
	// 		RemainingBandwidthInMbps: to.Ptr[int32](1500),
	// 		SecondaryAzurePort: to.Ptr("bvtazureixp01"),
	// 		UsedBandwidthInMbps: to.Ptr[int32](2500),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkManagerActiveConnectivityConfigurationsList.json
func ExampleManagementClient_ListActiveConnectivityConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ListActiveConnectivityConfigurations(ctx, "myResourceGroup", "testNetworkManager", armnetwork.ActiveConfigurationParameter{
		Regions: []*string{
			to.Ptr("westus")},
		SkipToken: to.Ptr("fakeSkipTokenCode"),
	}, &armnetwork.ManagementClientListActiveConnectivityConfigurationsOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActiveConnectivityConfigurationsListResult = armnetwork.ActiveConnectivityConfigurationsListResult{
	// 	SkipToken: to.Ptr("FakeSkipTokenCode"),
	// 	Value: []*armnetwork.ActiveConnectivityConfiguration{
	// 		{
	// 			ConfigurationGroups: []*armnetwork.ConfigurationGroup{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 					Properties: &armnetwork.GroupProperties{
	// 						Description: to.Ptr("A group for all test Virtual Networks"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/connectivityConfigurations/myTestConnectivityConfig"),
	// 			Properties: &armnetwork.ConnectivityConfigurationProperties{
	// 				Description: to.Ptr("Sample Configuration"),
	// 				AppliesToGroups: []*armnetwork.ConnectivityGroupItem{
	// 					{
	// 						GroupConnectivity: to.Ptr(armnetwork.GroupConnectivityNone),
	// 						IsGlobal: to.Ptr(armnetwork.IsGlobalFalse),
	// 						NetworkGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 						UseHubGateway: to.Ptr(armnetwork.UseHubGatewayTrue),
	// 				}},
	// 				ConnectivityTopology: to.Ptr(armnetwork.ConnectivityTopologyHubAndSpoke),
	// 				DeleteExistingPeering: to.Ptr(armnetwork.DeleteExistingPeeringTrue),
	// 				Hubs: []*armnetwork.Hub{
	// 					{
	// 						ResourceID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myTestConnectivityConfig"),
	// 						ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 				}},
	// 				IsGlobal: to.Ptr(armnetwork.IsGlobalTrue),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			},
	// 			CommitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-10T12:33:22.257Z"); return t}()),
	// 			Region: to.Ptr("westus"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkManagerActiveSecurityAdminRulesList.json
func ExampleManagementClient_ListActiveSecurityAdminRules() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ListActiveSecurityAdminRules(ctx, "myResourceGroup", "testNetworkManager", armnetwork.ActiveConfigurationParameter{
		Regions: []*string{
			to.Ptr("westus")},
		SkipToken: to.Ptr("fakeSkipTokenCode"),
	}, &armnetwork.ManagementClientListActiveSecurityAdminRulesOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActiveSecurityAdminRulesListResult = armnetwork.ActiveSecurityAdminRulesListResult{
	// 	SkipToken: to.Ptr("FakeSkipTokenCode"),
	// 	Value: []armnetwork.ActiveBaseSecurityAdminRuleClassification{
	// 		&armnetwork.ActiveDefaultSecurityAdminRule{
	// 			CommitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-10T12:33:22.257Z"); return t}()),
	// 			ConfigurationDescription: to.Ptr("SampleDescription"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleAdminRule"),
	// 			Kind: to.Ptr(armnetwork.EffectiveAdminRuleKindDefault),
	// 			Region: to.Ptr("westus"),
	// 			RuleCollectionAppliesToGroups: []*armnetwork.ManagerSecurityGroupItem{
	// 				{
	// 					NetworkGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 			}},
	// 			RuleCollectionDescription: to.Ptr("SampleRuleCollectionDescription"),
	// 			RuleGroups: []*armnetwork.ConfigurationGroup{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 					Properties: &armnetwork.GroupProperties{
	// 						Description: to.Ptr("A group for all test Virtual Networks"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			Properties: &armnetwork.DefaultAdminPropertiesFormat{
	// 				Description: to.Ptr("Sample Admin Rule"),
	// 				Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 				DestinationPortRanges: []*string{
	// 					to.Ptr("22")},
	// 					Destinations: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("*"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 					}},
	// 					Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 					Flag: to.Ptr("AllowVnetInbound"),
	// 					Priority: to.Ptr[int32](1),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourcePortRanges: []*string{
	// 						to.Ptr("0-65535")},
	// 						Sources: []*armnetwork.AddressPrefixItem{
	// 							{
	// 								AddressPrefix: to.Ptr("*"),
	// 								AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 						}},
	// 						Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 					},
	// 			}},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkManagerEffectiveConnectivityConfigurationsList.json
func ExampleManagementClient_ListNetworkManagerEffectiveConnectivityConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ListNetworkManagerEffectiveConnectivityConfigurations(ctx, "myResourceGroup", "testVirtualNetwork", armnetwork.QueryRequestOptions{
		SkipToken: to.Ptr("FakeSkipTokenCode"),
	}, &armnetwork.ManagementClientListNetworkManagerEffectiveConnectivityConfigurationsOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagerEffectiveConnectivityConfigurationListResult = armnetwork.ManagerEffectiveConnectivityConfigurationListResult{
	// 	SkipToken: to.Ptr("FakeSkipTokenCode"),
	// 	Value: []*armnetwork.EffectiveConnectivityConfiguration{
	// 		{
	// 			ConfigurationGroups: []*armnetwork.ConfigurationGroup{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 					Properties: &armnetwork.GroupProperties{
	// 						Description: to.Ptr("A group for all test Virtual Networks"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/connectivityConfigurations/myTestConnectivityConfig"),
	// 			Properties: &armnetwork.ConnectivityConfigurationProperties{
	// 				Description: to.Ptr("Sample Configuration"),
	// 				AppliesToGroups: []*armnetwork.ConnectivityGroupItem{
	// 					{
	// 						GroupConnectivity: to.Ptr(armnetwork.GroupConnectivityNone),
	// 						IsGlobal: to.Ptr(armnetwork.IsGlobalFalse),
	// 						NetworkGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/groups/group1"),
	// 						UseHubGateway: to.Ptr(armnetwork.UseHubGatewayTrue),
	// 				}},
	// 				ConnectivityTopology: to.Ptr(armnetwork.ConnectivityTopologyHubAndSpoke),
	// 				DeleteExistingPeering: to.Ptr(armnetwork.DeleteExistingPeeringTrue),
	// 				Hubs: []*armnetwork.Hub{
	// 					{
	// 						ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myHubVnet"),
	// 						ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 				}},
	// 				IsGlobal: to.Ptr(armnetwork.IsGlobalTrue),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkManagerEffectiveSecurityAdminRulesList.json
func ExampleManagementClient_ListNetworkManagerEffectiveSecurityAdminRules() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ListNetworkManagerEffectiveSecurityAdminRules(ctx, "myResourceGroup", "testVirtualNetwork", armnetwork.QueryRequestOptions{
		SkipToken: to.Ptr("FakeSkipTokenCode"),
	}, &armnetwork.ManagementClientListNetworkManagerEffectiveSecurityAdminRulesOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagerEffectiveSecurityAdminRulesListResult = armnetwork.ManagerEffectiveSecurityAdminRulesListResult{
	// 	SkipToken: to.Ptr("FakeSkipTokenCode"),
	// 	Value: []armnetwork.EffectiveBaseSecurityAdminRuleClassification{
	// 		&armnetwork.EffectiveDefaultSecurityAdminRule{
	// 			ConfigurationDescription: to.Ptr("SampleDescription"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleAdminRule"),
	// 			Kind: to.Ptr(armnetwork.EffectiveAdminRuleKindDefault),
	// 			RuleCollectionAppliesToGroups: []*armnetwork.ManagerSecurityGroupItem{
	// 				{
	// 					NetworkGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 			}},
	// 			RuleCollectionDescription: to.Ptr("SampleRuleCollectionDescription"),
	// 			RuleGroups: []*armnetwork.ConfigurationGroup{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 					Properties: &armnetwork.GroupProperties{
	// 						Description: to.Ptr("A group for all test Virtual Networks"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			Properties: &armnetwork.DefaultAdminPropertiesFormat{
	// 				Description: to.Ptr("Sample Admin Rule"),
	// 				Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 				DestinationPortRanges: []*string{
	// 					to.Ptr("22")},
	// 					Destinations: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("*"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 					}},
	// 					Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 					Flag: to.Ptr("AllowVnetInbound"),
	// 					Priority: to.Ptr[int32](1),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourcePortRanges: []*string{
	// 						to.Ptr("0-65535")},
	// 						Sources: []*armnetwork.AddressPrefixItem{
	// 							{
	// 								AddressPrefix: to.Ptr("*"),
	// 								AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 						}},
	// 						Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 					},
	// 			}},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/VirtualWanSupportedSecurityProviders.json
func ExampleManagementClient_SupportedSecurityProviders() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().SupportedSecurityProviders(ctx, "rg1", "wan1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualWanSecurityProviders = armnetwork.VirtualWanSecurityProviders{
	// 	SupportedProviders: []*armnetwork.VirtualWanSecurityProvider{
	// 		{
	// 			Name: to.Ptr("AzureFirewall"),
	// 			Type: to.Ptr(armnetwork.VirtualWanSecurityProviderTypeNative),
	// 			URL: to.Ptr(""),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/GenerateVirtualWanVpnServerConfigurationVpnProfile.json
func ExampleManagementClient_BeginGeneratevirtualwanvpnserverconfigurationvpnprofile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginGeneratevirtualwanvpnserverconfigurationvpnprofile(ctx, "rg1", "wan1", armnetwork.VirtualWanVPNProfileParameters{
		AuthenticationMethod:             to.Ptr(armnetwork.AuthenticationMethodEAPTLS),
		VPNServerConfigurationResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnconfig1"),
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
	// res.VPNProfileResponse = armnetwork.VPNProfileResponse{
	// 	ProfileURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// }
}
