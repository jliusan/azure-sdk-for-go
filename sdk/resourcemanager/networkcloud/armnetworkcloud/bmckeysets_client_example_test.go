//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_ListByResourceGroup.json
func ExampleBmcKeySetsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBmcKeySetsClient().NewListByResourceGroupPager("resourceGroupName", "clusterName", nil)
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
		// page.BmcKeySetList = armnetworkcloud.BmcKeySetList{
		// 	Value: []*armnetworkcloud.BmcKeySet{
		// 		{
		// 			Name: to.Ptr("bmcKeySetName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/clusters/bmcKeySets"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bmcKeySets/bmcKeySetName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.BmcKeySetProperties{
		// 				AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
		// 				DetailedStatus: to.Ptr(armnetworkcloud.BmcKeySetDetailedStatusSomeInvalid),
		// 				DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
		// 				Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
		// 				LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
		// 				PrivilegeLevel: to.Ptr(armnetworkcloud.BmcKeySetPrivilegeLevelAdministrator),
		// 				ProvisioningState: to.Ptr(armnetworkcloud.BmcKeySetProvisioningStateSucceeded),
		// 				UserList: []*armnetworkcloud.KeySetUser{
		// 					{
		// 						Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
		// 						AzureUserName: to.Ptr("userABC"),
		// 						SSHPublicKey: &armnetworkcloud.SSHPublicKey{
		// 							KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 						},
		// 					},
		// 					{
		// 						Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
		// 						AzureUserName: to.Ptr("userXYZ"),
		// 						SSHPublicKey: &armnetworkcloud.SSHPublicKey{
		// 							KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 						},
		// 				}},
		// 				UserListStatus: []*armnetworkcloud.KeySetUserStatus{
		// 					{
		// 						AzureUserName: to.Ptr("userABC"),
		// 						Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
		// 						StatusMessage: to.Ptr("User has been provisioned"),
		// 					},
		// 					{
		// 						AzureUserName: to.Ptr("userXYZ"),
		// 						Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
		// 						StatusMessage: to.Ptr("User is not a valid Azure user"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_Get.json
func ExampleBmcKeySetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBmcKeySetsClient().Get(ctx, "resourceGroupName", "clusterName", "bmcKeySetName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BmcKeySet = armnetworkcloud.BmcKeySet{
	// 	Name: to.Ptr("bmcKeySetName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusters/bmcKeySets"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bmcKeySets/bmcKeySetName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.BmcKeySetProperties{
	// 		AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.BmcKeySetDetailedStatusSomeInvalid),
	// 		DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
	// 		Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
	// 		LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
	// 		PrivilegeLevel: to.Ptr(armnetworkcloud.BmcKeySetPrivilegeLevelAdministrator),
	// 		ProvisioningState: to.Ptr(armnetworkcloud.BmcKeySetProvisioningStateSucceeded),
	// 		UserList: []*armnetworkcloud.KeySetUser{
	// 			{
	// 				Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 				AzureUserName: to.Ptr("userABC"),
	// 				SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 					KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 				},
	// 			},
	// 			{
	// 				Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 				AzureUserName: to.Ptr("userXYZ"),
	// 				SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 					KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 				},
	// 		}},
	// 		UserListStatus: []*armnetworkcloud.KeySetUserStatus{
	// 			{
	// 				AzureUserName: to.Ptr("userABC"),
	// 				Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
	// 				StatusMessage: to.Ptr("User has been provisioned"),
	// 			},
	// 			{
	// 				AzureUserName: to.Ptr("userXYZ"),
	// 				Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
	// 				StatusMessage: to.Ptr("User is not a valid Azure user"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_Create.json
func ExampleBmcKeySetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBmcKeySetsClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "clusterName", "bmcKeySetName", armnetworkcloud.BmcKeySet{
		Location: to.Ptr("location"),
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
			Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armnetworkcloud.BmcKeySetProperties{
			AzureGroupID:   to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
			Expiration:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t }()),
			PrivilegeLevel: to.Ptr(armnetworkcloud.BmcKeySetPrivilegeLevelAdministrator),
			UserList: []*armnetworkcloud.KeySetUser{
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userABC"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				},
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userXYZ"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				}},
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
	// res.BmcKeySet = armnetworkcloud.BmcKeySet{
	// 	Name: to.Ptr("bmcKeySetName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusters/bmcKeySets"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bmcKeySets/bmcKeySetName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.BmcKeySetProperties{
	// 		AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.BmcKeySetDetailedStatusSomeInvalid),
	// 		DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
	// 		Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
	// 		LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
	// 		PrivilegeLevel: to.Ptr(armnetworkcloud.BmcKeySetPrivilegeLevelAdministrator),
	// 		ProvisioningState: to.Ptr(armnetworkcloud.BmcKeySetProvisioningStateSucceeded),
	// 		UserList: []*armnetworkcloud.KeySetUser{
	// 			{
	// 				Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 				AzureUserName: to.Ptr("userABC"),
	// 				SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 					KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 				},
	// 			},
	// 			{
	// 				Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 				AzureUserName: to.Ptr("userXYZ"),
	// 				SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 					KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 				},
	// 		}},
	// 		UserListStatus: []*armnetworkcloud.KeySetUserStatus{
	// 			{
	// 				AzureUserName: to.Ptr("userABC"),
	// 				Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
	// 				StatusMessage: to.Ptr("User has been provisioned"),
	// 			},
	// 			{
	// 				AzureUserName: to.Ptr("userXYZ"),
	// 				Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
	// 				StatusMessage: to.Ptr("User is not a valid Azure user"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_Delete.json
func ExampleBmcKeySetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBmcKeySetsClient().BeginDelete(ctx, "resourceGroupName", "clusterName", "bmcKeySetName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_Patch.json
func ExampleBmcKeySetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBmcKeySetsClient().BeginUpdate(ctx, "resourceGroupName", "clusterName", "bmcKeySetName", armnetworkcloud.BmcKeySetPatchParameters{
		Properties: &armnetworkcloud.BmcKeySetPatchProperties{
			Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t }()),
			UserList: []*armnetworkcloud.KeySetUser{
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userABC"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				},
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userXYZ"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				}},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
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
	// res.BmcKeySet = armnetworkcloud.BmcKeySet{
	// 	Name: to.Ptr("bmcKeySetName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusters/bmcKeySets"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bmcKeySets/bmcKeySetName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.BmcKeySetProperties{
	// 		AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.BmcKeySetDetailedStatusSomeInvalid),
	// 		DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
	// 		Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
	// 		LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
	// 		PrivilegeLevel: to.Ptr(armnetworkcloud.BmcKeySetPrivilegeLevelAdministrator),
	// 		ProvisioningState: to.Ptr(armnetworkcloud.BmcKeySetProvisioningStateSucceeded),
	// 		UserList: []*armnetworkcloud.KeySetUser{
	// 			{
	// 				Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 				AzureUserName: to.Ptr("userABC"),
	// 				SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 					KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 				},
	// 			},
	// 			{
	// 				Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 				AzureUserName: to.Ptr("userXYZ"),
	// 				SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 					KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 				},
	// 		}},
	// 		UserListStatus: []*armnetworkcloud.KeySetUserStatus{
	// 			{
	// 				AzureUserName: to.Ptr("userABC"),
	// 				Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
	// 				StatusMessage: to.Ptr("User has been provisioned"),
	// 			},
	// 			{
	// 				AzureUserName: to.Ptr("userXYZ"),
	// 				Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
	// 				StatusMessage: to.Ptr("User is not a valid Azure user"),
	// 		}},
	// 	},
	// }
}
