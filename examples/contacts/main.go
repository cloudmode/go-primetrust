package main

import (
	"github.com/cloudmode/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	if contacts, err := primetrust.GetContacts(); err != nil {
		log.Println("Error getting contacts:", err.Error())
	} else {
		log.Printf("Contacts: %d", len(contacts.Data))
		if contact, err := primetrust.GetContact(contacts.Data[0].ID); err != nil {
			log.Println("Error getting contact:", err.Error())
		} else {
			log.Printf("Contact: %+v", contact)
		}
	}

	log.Println("Done")
}
