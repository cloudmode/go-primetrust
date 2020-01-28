package main

import (
	"github.com/cloudmode/go-primetrust"
	"github.com/cloudmode/go-primetrust/models"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	accountId := os.Getenv("PRIMETRUST_ACCOUNT_ID")
	webhookURL := os.Getenv("PRIMETRUST_WEBHOOK_URL")
	webhookEmail := os.Getenv("PRIMETRUST_WEBHOOK_EMAIL")
	webhookSecret := os.Getenv("PRIMETRUST_WEBHOOK_SECRET")

	webhook := models.NewWebhook(accountId, webhookURL, webhookEmail, webhookSecret, true)

	if _, err := primetrust.CreateNewWebhook(webhook); err != nil {
		log.Println("Error creating new webhook:", err)
	} else {
		log.Println("Webhook created OK")
	}

	log.Println("Done")
}
