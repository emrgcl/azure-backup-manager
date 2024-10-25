package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/charmbracelet/log"
)

func main() {
	// Create a context for the authentication request
	ctx := context.Background()

	// Use DefaultAzureCredential to authenticate with Azure
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain default azure credential: %v", err)
	}

	// Use the credential to create a client to interact with Azure
	client, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create Azure subscriptions client: %v", err)
	}

	log.Info("Successfully created Azure subscriptions client")

	// Get the subscription details
	pager := client.NewListPager(nil)
	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, sub := range resp.SubscriptionListResult.Value {
			fmt.Printf("Subscription ID: %s, Subscription Name: %s\n", *sub.SubscriptionID, *sub.DisplayName)
		}
	}
}
