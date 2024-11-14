package armsql_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/v3/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// C:\Users\v-liujudy\go\pkg\mod\github.com\!azure\azure-sdk-for-go\sdk\resourcemanager
const (
	SubscriptionID    = "faa080af-c1d8-40ad-9cce-e1a450ca5b57"
	ResourceGroupName = "judytest04-record"
	ResourceLocation  = "eastus2"
	SqlServerName     = "judy-sql-record"
	DatabaseName      = "test-record"
	PathToPackage     = "sdk/resourcemanager/sql/armsql/testdata"
)

type SqlAccessTestSuite struct {
	suite.Suite
	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	sqlServersName    string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *SqlAccessTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), PathToPackage)
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.location = recording.GetEnvVariable("LOCATION", ResourceLocation)
	testsuite.resourceGroupName = recording.GetEnvVariable("RESOURCE_GROUP_NAME", ResourceGroupName)
	testsuite.subscriptionId = recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", SubscriptionID)

	clientFactory, err := armresources.NewClientFactory(testsuite.subscriptionId, testsuite.cred, nil)
	client := clientFactory.NewResourceGroupsClient()
	ctx := context.Background()

	_, err = client.CreateOrUpdate(ctx, ResourceGroupName, armresources.ResourceGroup{
		Location: to.Ptr(ResourceLocation),
	}, nil)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = ResourceGroupName
	fmt.Println("create new resource group ", ResourceGroupName, " of ", SubscriptionID, "successfully")
}

func (testsuite *SqlAccessTestSuite) TestCreateServer() {
	fmt.Println("start to create sql server~")
	// create new msql resource under group
	sqlClientFactory, err := armsql.NewClientFactory(SubscriptionID, testsuite.cred, nil)
	testsuite.Require().NoError(err)
	serverclient := sqlClientFactory.NewServersClient()

	ctx := context.Background()
	server, err := createServer(testsuite.ctx, serverclient)
	testsuite.Require().NoError(err)
	fmt.Println("create serverId", *server.ID)

	server, err = getServer(ctx, serverclient)
	testsuite.Require().NoError(err)
	fmt.Println("get server:", *server.ID)

	// create database
	databasesClient := sqlClientFactory.NewDatabasesClient()
	testsuite.Require().NotNil(databasesClient)
	database, err := createDatabase(ctx, databasesClient)
	testsuite.Require().NoError(err)
	testsuite.Require().NotNil(database)

	fmt.Println("database:", *database.ID)
}

func TestSqlAccessTestSuite(t *testing.T) {
	suite.Run(t, new(SqlAccessTestSuite))
}

func TestCreateOrUpdateGroupOfSubscriptionId(t *testing.T) {
	// subsriptionId := os.Getenv("AZURE_SUBSCRIPTION_ID")
	subsriptionId := "faa080af-c1d8-40ad-9cce-e1a450ca5b57"
	assert.NotEmpty(t, subsriptionId)
	// get default credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	assert.Nil(t, err)
	assert.NotNil(t, cred)
	fmt.Println("get default credential")

	// new client factory
	clientFactory, err := armresources.NewClientFactory(subsriptionId, cred, nil)
	assert.Nil(t, err)
	assert.NotNil(t, clientFactory)
	client := clientFactory.NewResourceGroupsClient()
	assert.NotNil(t, client)
	ctx := context.Background()

	createOrUpdateGroupResponse, err := client.CreateOrUpdate(ctx, ResourceGroupName, armresources.ResourceGroup{
		Location: to.Ptr(ResourceLocation),
	}, nil)
	assert.Nil(t, err)
	assert.NotNil(t, createOrUpdateGroupResponse)
	fmt.Println("create resource group client")

	// check whether create new group successfully
	checkGroupExistResponse, err := client.CheckExistence(ctx, ResourceGroupName, nil)
	assert.Nil(t, err)
	assert.True(t, checkGroupExistResponse.Success)
	fmt.Println("create new resource group ", ResourceGroupName, " of ", subsriptionId, "successfully")
}

func NewTestDeleteGroupOfSubscriptionId(t *testing.T, ctx context.Context, client *armresources.ResourceGroupsClient) {
	// delete resource group created newly
	resourceGroupsClientDeleteResponse, err := client.BeginDelete(ctx, ResourceGroupName, nil)
	assert.Nil(t, err)
	time.Sleep(time.Second * 2)
	response, err := resourceGroupsClientDeleteResponse.Poll(ctx)
	assert.Nil(t, err)
	assert.True(t, response.StatusCode >= 200 && response.StatusCode < 300)
	fmt.Println("delete resource group successfully")
}

func TestCreateOrUpdateArmMysqlResourceOnGroup(t *testing.T) {
	// make sure that the group has been created
	TestCreateOrUpdateGroupOfSubscriptionId(t)

	cred := GetAzureCredentail()
	assert.NotNil(t, cred)
	// create new msql resource under group
	sqlClientFactory, err := armsql.NewClientFactory(SubscriptionID, cred, nil)
	assert.Nil(t, err)
	assert.NotNil(t, sqlClientFactory)

	serverclient := sqlClientFactory.NewServersClient()
	assert.NotNil(t, serverclient)

	// resourceGroup, err := CreateAzureResourceGroup()
	// assert.Nil(t, err)
	ctx := context.Background()
	server, err := createServer(ctx, serverclient)
	assert.Nil(t, err)
	assert.NotNil(t, server)
	fmt.Println("create serverId", *server.ID)

	server, err = getServer(ctx, serverclient)
	assert.Nil(t, err)
	assert.NotNil(t, server)
	fmt.Println("get server:", *server.ID)

	// create database
	databasesClient := sqlClientFactory.NewDatabasesClient()
	assert.NotNil(t, databasesClient)
	database, err := createDatabase(ctx, databasesClient)
	assert.Nil(t, err)
	assert.NotNil(t, database)
	fmt.Println("database:", *database.ID)

	// cleanup(context.Background(), resourceGroup)

}

func TestCleanUpArmResourceGroup(t *testing.T) {
	cred := GetAzureCredentail()
	clientFactory, err := armresources.NewClientFactory(SubscriptionID, cred, nil)
	assert.Nil(t, err)
	assert.NotNil(t, clientFactory)

	// clean up resource group
	ctx := context.Background()
	client := clientFactory.NewResourceGroupsClient()
	err = cleanup(ctx, client)
	assert.Nil(t, err)

	// check group is cleaned up successfully
	checkGroupExistResponse, err := client.CheckExistence(ctx, ResourceGroupName, nil)
	assert.Nil(t, err)
	assert.False(t, checkGroupExistResponse.Success)
}

func GetAzureCredentail() (cred *azidentity.DefaultAzureCredential) {
	// get default credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("get default credentail fail, error:", err.Error())
		return
	}
	fmt.Println("get default credential")
	return cred
}

func CreateAzureResourceGroup(groupName string) (group *armresources.ResourceGroup, err error) {
	cred := GetAzureCredentail()
	clientFactory, err := armresources.NewClientFactory(SubscriptionID, cred, nil)
	if err != nil {
		fmt.Println("create resource client factory fail, error:", err.Error())
		return
	}
	client := clientFactory.NewResourceGroupsClient()
	createOrUpdateGroupResponse, err := client.CreateOrUpdate(context.TODO(), groupName, armresources.ResourceGroup{
		Location: to.Ptr(ResourceLocation),
	}, nil)
	if err != nil {
		fmt.Println("create resource group fail, error:", err.Error())
		return
	}
	group = &createOrUpdateGroupResponse.ResourceGroup
	return
}

func createServer(ctx context.Context, serversClient *armsql.ServersClient) (*armsql.Server, error) {

	pollerResp, err := serversClient.BeginCreateOrUpdate(
		ctx,
		ResourceGroupName,
		SqlServerName,
		armsql.Server{
			Location: to.Ptr(ResourceLocation),
			Properties: &armsql.ServerProperties{
				AdministratorLogin:         to.Ptr("dummylogin"),
				AdministratorLoginPassword: to.Ptr("QWE123!@#"),
			},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.Server, nil
}

func getServer(ctx context.Context, serversClient *armsql.ServersClient) (*armsql.Server, error) {

	resp, err := serversClient.Get(ctx, ResourceGroupName, SqlServerName, nil)
	if err != nil {
		return nil, err
	}
	return &resp.Server, nil
}

func cleanup(ctx context.Context, resourceGroupClient *armresources.ResourceGroupsClient) error {

	pollerResp, err := resourceGroupClient.BeginDelete(ctx, ResourceGroupName, nil)
	if err != nil {
		return err
	}
	_, err = pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func createDatabase(ctx context.Context, databasesClient *armsql.DatabasesClient) (*armsql.Database, error) {

	pollerResp, err := databasesClient.BeginCreateOrUpdate(
		ctx,
		ResourceGroupName,
		SqlServerName,
		DatabaseName,
		armsql.Database{
			Location: to.Ptr(ResourceLocation),
			Properties: &armsql.DatabaseProperties{
				ReadScale: to.Ptr(armsql.DatabaseReadScaleDisabled),
			},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.Database, nil
}
