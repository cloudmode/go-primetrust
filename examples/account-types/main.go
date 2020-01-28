package main

import (
	"github.com/cloudmode/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	if accountTypes, err := primetrust.GetAccountTypes(); err != nil {
		log.Println("Error getting account types:", err.Error())
	} else {
		log.Printf("Account Types: %d", len(accountTypes.Data))
		if accountType, err := primetrust.GetAccountType(accountTypes.Data[0].ID); err != nil {
			log.Printf("Error getting \"%s\" account type:", accountTypes.Data[0].ID, err.Error())
		} else {
			log.Printf("Account Type \"%s\": %+v", accountTypes.Data[0].ID, accountType)
		}
	}

	log.Println("Done")
}
