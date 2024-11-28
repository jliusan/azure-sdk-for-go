package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

func printErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	// clientID := "my-app-uuid"
	// tenantID := "72f988bf-86f1-41af-91ab-2d7cd011db47"
	subscriptionID := "faa080af-c1d8-40ad-9cce-e1a450ca5b57"
	resourceGroup := "judytest01"

	// tokenFile := "kube-token.txt"
	// opts := &azidentity.WorkloadIdentityCredentialOptions{
	// 	AdditionallyAllowedTenants: []string{},
	// 	ClientID:                   clientID,
	// 	TenantID:                   tenantID,
	// 	TokenFilePath:              tokenFile,
	// }

	// creds, err := azidentity.NewWorkloadIdentityCredential(opts)
	// if err != nil {
	// 	printErr(err)
	// }
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("get default credentail fail, error:", err.Error())
		return
	}
	fmt.Println("get default credential")

	var rawResponse *http.Response
	ctx := context.TODO() // your context
	ctxWithResp := runtime.WithCaptureResponse(ctx, &rawResponse)

	printLoadBalancers(ctxWithResp, creds, subscriptionID, resourceGroup)
}

func printLoadBalancers(ctx context.Context, creds azcore.TokenCredential, subscriptionID string, resourceGroup string) {
	factory, err := armnetwork.NewClientFactory(subscriptionID, creds, &arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Logging: policy.LogOptions{
				// include HTTP body for log
				IncludeBody: true,
			},
		},
	})
	if err != nil {
		printErr(err)
	}

	client := factory.NewLoadBalancersClient()
	pager := client.NewListPager(resourceGroup, &armnetwork.LoadBalancersClientListOptions{})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			printErr(err)
		}

		for _, item := range page.Value {
			fmt.Printf("ID: %s\n", *item.ID)
			fmt.Printf("\tLocation: %s\n", *item.Location)
			fmt.Printf("\tName: %s\n", *item.Name)
			fmt.Printf("\tEtag: %s\n", *item.Etag)
			fmt.Printf("\tType: %s\n", *item.Type)
			fmt.Printf("\tSKU Name: %s\n", *item.SKU.Name)
			fmt.Printf("\tSKU Tier: %s\n\n", *item.SKU.Tier)

			for _, frontend := range item.Properties.FrontendIPConfigurations {
				if frontend.Properties != nil {
					fmt.Printf("\t\tName: %s\n", *frontend.Name)
					fmt.Printf("\t\tPrivate IP: %v\n", frontend.Properties.PrivateIPAddress)
					fmt.Printf("\t\tPublic IP: %v\n", frontend.Properties.PublicIPAddress.Properties)
				}
			}
		}
	}
}
