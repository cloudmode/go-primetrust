package main

import (
	"github.com/cloudmode/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	accountId := os.Getenv("PRIMETRUST_ACCOUNT_ID")

	if webhook, err := primetrust.GetLastWebhook(accountId); err != nil {
		log.Println("Error getting webhooks:", err.Error())
	} else {
		log.Printf("Webhook: %+v", webhook)
	}

	log.Println("Done")
}
