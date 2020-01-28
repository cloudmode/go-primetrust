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
	webhookId := os.Getenv("PRIMETRUST_WEBHOOK_ID")
	webhookURL := os.Getenv("PRIMETRUST_WEBHOOK_URL")
	webhookEmail := os.Getenv("PRIMETRUST_WEBHOOK_EMAIL")
	webhookSecret := os.Getenv("PRIMETRUST_WEBHOOK_SECRET")

	log.Printf("Getting webhook: %s", webhookId)

	webhook := models.NewWebhook(accountId, webhookURL, webhookEmail, webhookSecret, true)
	webhook.Data.ID = webhookId
	if _, err := primetrust.UpdateWebhook(webhook); err != nil {
		log.Println("Error updating webhook:", err)
	} else {
		log.Println("Webhook updated OK")
	}

	log.Println("Done")
}
