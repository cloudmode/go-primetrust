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
	contact := models.NewCompanyContact(accountId)

	contact.Data.Attributes.RegionOfFormation = "TX"
	contact.Data.Attributes.Email = "acme@example.com"
	contact.Data.Attributes.Name = "ACME LLC"
	contact.Data.Attributes.Label = "Label"
	contact.Data.Attributes.TaxIDNumber = "888777666"
	contact.Data.Attributes.TaxCountry = "US"
	contact.Data.Attributes.TaxState = "TX"
	contact.Data.Attributes.PrimaryAddress = models.Address{
		Type:       models.AddressTypeOffice,
		City:       "Houston",
		Country:    "US",
		PostalCode: "60000",
		Region:     "TX",
		Street1:    "1 Main str.",
	}
	contact.Data.Attributes.PrimaryPhoneNumber = models.PhoneNumber{
		Number: "+1 800 555 11 22",
	}
	contact.Data.Attributes.RelatedContacts = []models.RelatedContactData{
		models.RelatedContactData{
			Type:        models.ContactTypeNaturalPerson,
			DateOfBirth: "1981-10-31",
			Email:       "name@example.com",
			Sex:         "male",
			Name:        "Sample Person",
			TaxCountry:  "US",
			Label:       "CEO",
			TaxIDNumber: "123123123",
			PrimaryAddress: models.Address{
				Type:       models.AddressTypeHome,
				City:       "Beverly Hills",
				Country:    "US",
				PostalCode: "90210",
				Region:     "CA",
				Street1:    "1 Sunset Blvd.",
			},
			PrimaryPhoneNumber: models.PhoneNumber{
				Number: "+15555555577",
			},
		},
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
