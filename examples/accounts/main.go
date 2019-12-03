package main

import (
	"github.com/CLOUDMODE/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	if accounts, err := primetrust.GetAccounts(); err != nil {
		log.Println("Error getting accounts:", err.Error())
	} else {
		log.Printf("Accounts: %d", len(accounts.Data))
		if account, err := primetrust.GetAccount(accounts.Data[0].ID); err != nil {
			log.Println("Error getting account:", err.Error())
		} else {
			log.Printf("Account: %+v", account)
		}
	}

	log.Println("Done")
}
