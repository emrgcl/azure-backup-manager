package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
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

	subscriptions, err := GetSubscriptions(ctx, client)
	if err != nil {
		log.Fatalf("failed to get subscriptions: %v", err)
	}

	fmt.Println("Stored Subscriptions:")
	for _, sub := range subscriptions {
		fmt.Printf("Subscription ID: %s, Subscription Name: %s\n", *sub.SubscriptionID, *sub.DisplayName)
	}
}

// function to retrieve all storage accounts for the  specified subscription
func GetStorageAccounts(ctx context.Context, client *armsubscriptions.Client, subscriptionID string) ([]*armstorage.Account, error) {
	var storageAccounts []*armstorage.Account

	pager := client.NewListStorageAccountsPager(subscriptionID, nil)
	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to advance page: %v", err)
		}
		// Append each storage account object to the slice
		storageAccounts = append(storageAccounts, resp.StorageAccountListResult.Value...)
	}

	return storageAccounts, nil
}

// GetSubscriptions retrieves all subscriptions for the authenticated user

func GetSubscriptions(ctx context.Context, client *armsubscriptions.Client) ([]*armsubscriptions.Subscription, error) {
	var subscriptions []*armsubscriptions.Subscription

	pager := client.NewListPager(nil)
	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to advance page: %v", err)
		}
		// Append each subscription object to the slice
		subscriptions = append(subscriptions, resp.SubscriptionListResult.Value...)
	}

	return subscriptions, nil
}
