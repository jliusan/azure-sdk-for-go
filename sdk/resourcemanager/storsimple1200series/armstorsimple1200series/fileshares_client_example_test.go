//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListByFileServer.json
func ExampleFileSharesClient_NewListByFileServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListByFileServerPager("HSDK-ARCSX4MVKZ", "HSDK-ARCSX4MVKZ", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.FileShareList = armstorsimple1200series.FileShareList{
		// 	Value: []*armstorsimple1200series.FileShare{
		// 		{
		// 			Name: to.Ptr("TieredFileShareForSDKTest"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/fileServers/HSDK-ARCSX4MVKZ/shares/TieredFileShareForSDKTest"),
		// 			Properties: &armstorsimple1200series.FileShareProperties{
		// 				Description: to.Ptr("Demo FileShare for SDK Test Tiered"),
		// 				AdminUser: to.Ptr("fareast\\idcdlslb"),
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](174063616),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOnline),
		// 				UsedCapacityInBytes: to.Ptr[int64](174063616),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesGet.json
func ExampleFileSharesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Get(ctx, "HSDK-4XY4FI2IVG", "HSDK-4XY4FI2IVG", "Auto-TestFileShare1", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorsimple1200series.FileShare{
	// 	Name: to.Ptr("Auto-TestFileShare1"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG/shares/Auto-TestFileShare1"),
	// 	Properties: &armstorsimple1200series.FileShareProperties{
	// 		Description: to.Ptr("Demo FileShare for SDK Test Tiered"),
	// 		AdminUser: to.Ptr("fareast\\idcdlslb"),
	// 		DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
	// 		LocalUsedCapacityInBytes: to.Ptr[int64](0),
	// 		MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
	// 		ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
	// 		ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOnline),
	// 		UsedCapacityInBytes: to.Ptr[int64](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesCreateOrUpdate.json
func ExampleFileSharesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSharesClient().BeginCreateOrUpdate(ctx, "HSDK-4XY4FI2IVG", "HSDK-4XY4FI2IVG", "Auto-TestFileShare1", "ResourceGroupForSDKTest", "hAzureSDKOperations", armstorsimple1200series.FileShare{
		Name: to.Ptr("Auto-TestFileShare1"),
		Properties: &armstorsimple1200series.FileShareProperties{
			Description:                to.Ptr("Demo FileShare for SDK Test Tiered"),
			AdminUser:                  to.Ptr("fareast\\idcdlslb"),
			DataPolicy:                 to.Ptr(armstorsimple1200series.DataPolicyTiered),
			MonitoringStatus:           to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
			ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
			ShareStatus:                to.Ptr(armstorsimple1200series.ShareStatusOnline),
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
	// res.FileShare = armstorsimple1200series.FileShare{
	// 	Name: to.Ptr("Auto-TestFileShare1"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG/shares/Auto-TestFileShare1"),
	// 	Properties: &armstorsimple1200series.FileShareProperties{
	// 		Description: to.Ptr("Demo FileShare for SDK Test Tiered"),
	// 		AdminUser: to.Ptr("fareast\\idcdlslb"),
	// 		DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
	// 		LocalUsedCapacityInBytes: to.Ptr[int64](0),
	// 		MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
	// 		ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
	// 		ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOnline),
	// 		UsedCapacityInBytes: to.Ptr[int64](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesDelete.json
func ExampleFileSharesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSharesClient().BeginDelete(ctx, "HSDK-DMNJB2PET0", "HSDK-DMNJB2PET0", "Auto-TestFileShare2", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListMetrics.json
func ExampleFileSharesClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListMetricsPager("HSDK-DMNJB2PET0", "HSDK-DMNJB2PET0", "Auto-TestFileShare2", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.FileSharesClientListMetricsOptions{Filter: to.Ptr("startTime%20ge%20'2018-08-10T18:30:00Z'%20and%20endTime%20le%20'2018-08-11T18:30:00Z'")})
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
		// page.MetricList = armstorsimple1200series.MetricList{
		// 	Value: []*armstorsimple1200series.Metrics{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Storage Used"),
		// 				Value: to.Ptr("HostUsedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares/metrics"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Share"),
		// 					Value: to.Ptr("Auto-TestFileShare2"),
		// 			}},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-11T18:30:00Z"); return t}()),
		// 			PrimaryAggregation: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374/devices/HSDK-DMNJB2PET0/fileservers/HSDK-DMNJB2PET0/shares/Auto-TestFileShare2"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:30:00Z"); return t}()),
		// 			TimeGrain: to.Ptr("1.00:00:00"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 			Values: []*armstorsimple1200series.MetricData{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListMetricDefinition.json
func ExampleFileSharesClient_NewListMetricDefinitionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListMetricDefinitionPager("HSDK-DMNJB2PET0", "HSDK-DMNJB2PET0", "Auto-TestFileShare2", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.MetricDefinitionList = armstorsimple1200series.MetricDefinitionList{
		// 	Value: []*armstorsimple1200series.MetricDefinition{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Storage Used"),
		// 				Value: to.Ptr("HostUsedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares/metricsDefinitions"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Share"),
		// 					Value: to.Ptr("Auto-TestFileShare2"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple1200series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1Y"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374/devices/HSDK-DMNJB2PET0/fileservers/HSDK-DMNJB2PET0/shares/Auto-TestFileShare2"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListByDevice.json
func ExampleFileSharesClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListByDevicePager("HSDK-4XY4FI2IVG", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.FileShareList = armstorsimple1200series.FileShareList{
		// 	Value: []*armstorsimple1200series.FileShare{
		// 		{
		// 			Name: to.Ptr("Auto-TestFileShare2"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG/shares/Auto-TestFileShare2"),
		// 			Properties: &armstorsimple1200series.FileShareProperties{
		// 				Description: to.Ptr("Demo FileShare for SDK Test Local"),
		// 				AdminUser: to.Ptr("fareast\\idcdlslb"),
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyLocal),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOnline),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Auto-TestFileShare1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG/shares/Auto-TestFileShare1"),
		// 			Properties: &armstorsimple1200series.FileShareProperties{
		// 				Description: to.Ptr("Updated: Demo FileShare for SDK Test Tiered"),
		// 				AdminUser: to.Ptr("fareast\\idcdlslb"),
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOffline),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 	}},
		// }
	}
}