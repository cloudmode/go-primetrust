package main

import (
	"github.com/CLOUDMODE/go-primetrust"
	"github.com/CLOUDMODE/go-primetrust/models"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	accountId := os.Getenv("PRIMETRUST_ACCOUNT_ID")
	contact := models.NewNaturalPersonContact(accountId)

	contact.Data.Attributes.DateOfBirth = "1981-10-31"
	contact.Data.Attributes.Email = "name@example.com"
	contact.Data.Attributes.Sex = "male"
	contact.Data.Attributes.Name = "Sample Person"
	contact.Data.Attributes.TaxCountry = "US"
	contact.Data.Attributes.Label = "CEO"
	contact.Data.Attributes.TaxIDNumber = "123123123"
	contact.Data.Attributes.PrimaryAddress = models.Address{
		Type:       models.AddressTypeHome,
		City:       "Beverly Hills",
		Country:    "US",
		PostalCode: "90210",
		Region:     "CA",
		Street1:    "1 Sunset Blvd.",
	}
	contact.Data.Attributes.PrimaryPhoneNumber = models.PhoneNumber{
		Number: "+15555555577",
	}

	if newContact, err := primetrust.CreateNewContact(contact); err != nil {
		log.Println("Error creating new contact:", err)
		log.Printf("%+v", contact)
	} else {
		log.Println("Contact created OK")
		log.Printf("%+v", newContact)
	}

	log.Println("Done")
}
