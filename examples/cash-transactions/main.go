package main

import (
	"github.com/CLOUDMODE/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	if cashTransactions, err := primetrust.GetCashTransactions(); err != nil {
		log.Println("Error getting cash transactions:", err.Error())
	} else {
		log.Printf("Cash transactions: %d", len(cashTransactions.Data))
	}

	log.Println("Done")
}
