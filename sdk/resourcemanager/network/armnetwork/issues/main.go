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
	opts := &azidentity.DefaultAzureCredentialOptions{
		TenantID: "72f988bf-86f1-41af-91ab-2d7cd011db47",
	}

	creds, err := azidentity.NewDefaultAzureCredential(opts)
	if err != nil {
		printErr(err)
	}
	var rawResponse *http.Response
	ctx := context.TODO() // your context
	ctxWithResp := runtime.WithCaptureResponse(ctx, &rawResponse)
	printNetworkInterfaces(ctxWithResp, creds, "faa080af-c1d8-40ad-9cce-e1a450ca5b57", "judytest01")
}

func printNetworkInterfaces(ctx context.Context, creds azcore.TokenCredential, subscriptionID string, resourceGroup string) {
	factory, err := armnetwork.NewClientFactory(subscriptionID, creds, &arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Logging: policy.LogOptions{
				// include HTTP body for log
				IncludeBody: true,
			},
		},
	})

	// azlog.SetListener(func(event azlog.Event, s string) {
	// 	fmt.Println(s)
	// })

	// // include only azidentity credential logs
	// azlog.SetEvents(azidentity.EventAuthentication)

	if err != nil {
		printErr(err)
	}

	client := factory.NewInterfacesClient()

	pager := client.NewListPager(resourceGroup, &armnetwork.InterfacesClientListOptions{})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			printErr(err)
		}

		for _, item := range page.Value {
			fmt.Printf("ID: %s\n", *item.ID)
			if item.Properties != nil {
				if item.Properties.IPConfigurations != nil {
					for _, ip := range item.Properties.IPConfigurations {
						if ip.Properties != nil {
							fmt.Printf("\tPrivate IP: %s\n", *ip.Properties.PrivateIPAddress)
							fmt.Printf("\tAllocation method: %s\n", *ip.Properties.PrivateIPAllocationMethod)
							fmt.Printf("\tSubnet: %v\n", ip.Properties.Subnet.Name) // Always `nil'
						}
					}
				}
			}
		}
	}
}
