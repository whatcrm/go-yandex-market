package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/whatcrm/go-yandex-market"
)

func main() {
	apiKey := os.Getenv("YANDEX_MARKET_API_KEY")
	if apiKey == "" {
		log.Fatal("set YANDEX_MARKET_API_KEY")
	}

	client, err := yandexmarket.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	info, err := client.GetAuthTokenInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("token: %s, scopes: %v\n", info.Result.ApiKey.Name, info.Result.ApiKey.AuthScopes)

	list, err := client.GetCampaigns(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, camp := range list.Campaigns {
		domain := ""
		if camp.Domain != nil {
			domain = *camp.Domain
		}
		id := int64(0)
		if camp.Id != nil {
			id = *camp.Id
		}
		fmt.Printf("campaign id=%d domain=%s\n", id, domain)
	}
}
