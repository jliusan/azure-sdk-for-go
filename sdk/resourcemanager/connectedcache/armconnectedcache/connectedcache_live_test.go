//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
package armconnectedcache_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/v3/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/stretchr/testify/suite"
)

const (
	// ResourceLocation = "eastus2"
	ResourceLocation = "westus"
)

type ConnectedCacheTestSuite struct {
	suite.Suite
	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	armEndpoint       string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *ConnectedCacheTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), pathToPackage)
	fmt.Println("start recording~~~~~~~")
	testsuite.ctx = context.Background()
	testsuite.armEndpoint = "https://management.azure.com"
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.location = recording.GetEnvVariable("LOCATION", ResourceLocation)
	testsuite.resourceGroupName = recording.GetEnvVariable("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	testsuite.resourceGroupName = "sadafsdfsdb-" + strconv.Itoa(rand.Intn(100))
	fmt.Println("testsuite.resourceGroupName:", testsuite.resourceGroupName)
	testsuite.Prepare()
}

func (testsuite *ConnectedCacheTestSuite) TearDownSuite() {
	testsuite.Cleanup()
	// _, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	// testsuite.Require().NoError(err)
	time.Sleep(time.Second * 3)
	testutil.StopRecording(testsuite.T())
}

func TestConnectedCacheTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectedCacheTestSuite))
}

func (testsuite *ConnectedCacheTestSuite) Cleanup() {
	clientFactory, err := armresources.NewClientFactory(testsuite.subscriptionId, testsuite.cred, nil)
	testsuite.Require().NoError(err)
	resourceGroupsClientDeleteResponse, err := clientFactory.NewResourceGroupsClient().BeginDelete(testsuite.ctx, testsuite.resourceGroupName, nil)
	testsuite.Require().NoError(err)
	_, err = resourceGroupsClientDeleteResponse.Poll(testsuite.ctx)
	testsuite.Require().NoError(err)
}

func (testsuite *ConnectedCacheTestSuite) TestCreateIspCustomersClient() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	fmt.Println("start TestCreateIspCustomersClient====")
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory(recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000"), cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIspCustomersClient().BeginCreateOrUpdate(ctx, testsuite.resourceGroupName, "MccRPTest2", armconnectedcache.IspCustomerResource{
		Location: to.Ptr(testsuite.location),
		Properties: &armconnectedcache.CustomerProperty{
			// Customer: &armconnectedcache.CustomerEntity{
			// 	FullyQualifiedResourceID: to.Ptr("uqsbtgae"),
			// 	CustomerName:             to.Ptr("mkpzynfqihnjfdbaqbqwyhd"),
			// 	ContactEmail:             to.Ptr("xquos"),
			// 	ContactPhone:             to.Ptr("vue"),
			// 	ContactName:              to.Ptr("wxyqjoyoscmvimgwhpitxky"),
			// 	IsEntitled:               to.Ptr(true),
			// 	ReleaseVersion:           to.Ptr[int32](20),
			// 	ClientTenantID:           to.Ptr("fproidkpgvpdnac"),
			// 	IsEnterpriseManaged:      to.Ptr(true),
			// 	ShouldMigrate:            to.Ptr(true),
			// 	ResendSignupCode:         to.Ptr(true),
			// 	VerifySignupCode:         to.Ptr(true),
			// 	VerifySignupPhrase:       to.Ptr("tprjvttkgmrqlsyicnidhm"),
			// },
			// AdditionalCustomerProperties: &armconnectedcache.AdditionalCustomerProperties{
			// 	CustomerEmail:                 to.Ptr("zdjgibsidydyzm"),
			// 	CustomerTransitAsn:            to.Ptr("habgklnxqzmozqpazoyejwiphezpi"),
			// 	CustomerAsn:                   to.Ptr("hgrelgnrtdkleisnepfolu"),
			// 	CustomerEntitlementSKUID:      to.Ptr("b"),
			// 	CustomerEntitlementSKUGUID:    to.Ptr("rvzmdpxyflgqetvpwupnfaxsweiiz"),
			// 	CustomerEntitlementSKUName:    to.Ptr("waaqfijr"),
			// 	CustomerEntitlementExpiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.773Z"); return t }()),
			// 	// OptionalProperty1:             to.Ptr("qhmwxza"),
			// OptionalProperty2:             to.Ptr("l"),
			// OptionalProperty3:             to.Ptr("mblwwvbie"),
			// OptionalProperty4:             to.Ptr("vzuek"),
			// OptionalProperty5:             to.Ptr("fzjodscdfcdr"),
			// },
			Error: &armconnectedcache.ErrorDetail{},
		},
		Tags: map[string]*string{
			"key1878": to.Ptr("warz"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	fmt.Println("finish TestCreateIspCustomersClient====")
}

func (testsuite *ConnectedCacheTestSuite) TestGetIspCustomersClient() {
	// make sure that the group has been created
	ResourceName := recording.GetEnvVariable("RESOURCE_NAME", "MccRPTest2")
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory(testsuite.subscriptionId, cred, nil)

	_, err1 := clientFactory.NewIspCustomersClient().Get(ctx, testsuite.resourceGroupName, ResourceName, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err1)
	}
	fmt.Println("TestGetIspCustomersClient over")
}

func (testsuite *ConnectedCacheTestSuite) Prepare() {
	// get default credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	testsuite.Require().NoError(err)
	// new client factory

	fmt.Println("subscriptionId", testsuite.subscriptionId, "groupName", testsuite.resourceGroupName, "location", testsuite.location)
	clientFactory, err := armresources.NewClientFactory(testsuite.subscriptionId, cred, nil)
	testsuite.Require().NoError(err)
	client := clientFactory.NewResourceGroupsClient()
	ctx := context.Background()

	testsuite.Require().NoError(err)
	// check whether create new group successfully
	res, err := client.CheckExistence(ctx, testsuite.resourceGroupName, nil)
	testsuite.Require().NoError(err)
	if !res.Success {
		_, err = client.CreateOrUpdate(ctx, testsuite.resourceGroupName, armresources.ResourceGroup{
			Location: to.Ptr(testsuite.location),
		}, nil)
		testsuite.Require().NoError(err)
	}

	fmt.Println("create new resource group ", testsuite.resourceGroupName, " of ", testsuite.subscriptionId, "successfully")
}
