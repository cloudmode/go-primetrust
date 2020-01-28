package models

const (
	ContactsType = "contacts"

	ContactTypeNaturalPerson = "natural_person"
	ContactTypeCompany       = "company"
	ContactTypeCCorp         = "c-corp"
	ContactTypeLLC           = "llc"
	ContactTypeSCorp         = "s-corp"
	ContactTypeTrust         = "trust"

	AccountRoleOwner            = "owner"
	AccountRoleTaxFormRecipient = "tax form recipient"
)

type ContactAttributes struct {
	ID                 string               `json:"id,omitempty"`
	Type               string               `json:"type,omitempty"`
	AccountID          string               `json:"account-id,omitempty"`
	AccountRoles       []string             `json:"account-roles,omitempty"`
	ContactType        string               `json:"contact-type,omitempty"`
	AMLCleared         bool                 `json:"aml-cleared,omitempty"`
	CIPCleared         bool                 `json:"cip-cleared,omitempty"`
	DateOfBirth        string               `json:"date-of-birth,omitempty"`
	Email              string               `json:"email,omitempty"`
	Name               string               `json:"name,omitempty"`
	Sex                string               `json:"sex,omitempty"`
	Label              string               `json:"label,omitempty"`
	RegionOfFormation  string               `json:"region-of-formation,omitempty"`
	TaxIDNumber        string               `json:"tax-id-number,omitempty"`
	TaxCountry         string               `json:"tax-country,omitempty"`
	TaxState           string               `json:"tax-state,omitempty"`
	PrimaryAddress     Address              `json:"primary-address,omitempty"`
	PrimaryPhoneNumber PhoneNumber          `json:"primary-phone-number,omitempty"`
	RelatedContacts    []RelatedContactData `json:"related-contacts,omitempty"`
}

type ContactData struct {
	ID            string            `json:"id,omitempty"`
	Type          string            `json:"type"`
	Attributes    ContactAttributes `json:"attributes"`
	Links         Links             `json:"links"`
	Relationships Relationships     `json:"relationships"`
}

type RelatedContactData struct {
	Type               string      `json:"type"`
	ContactType        string      `json:"contact-type,omitempty"`
	DateOfBirth        string      `json:"date-of-birth"`
	Email              string      `json:"email"`
	Name               string      `json:"name"`
	Sex                string      `json:"sex"`
	Label              string      `json:"label"`
	TaxIDNumber        string      `json:"tax-id-number"`
	TaxCountry         string      `json:"tax-country"`
	PrimaryAddress     Address     `json:"primary-address"`
	PrimaryPhoneNumber PhoneNumber `json:"primary-phone-number"`
}

type Contact struct {
	Data ContactData `json:"data"`
}

func NewNaturalPersonContact(accountId string) *Contact {
	contact := Contact{
		Data: ContactData{
			Type: ContactsType,
			Attributes: ContactAttributes{
				AccountID:    accountId,
				AccountRoles: []string{AccountRoleOwner},
				Type:         ContactTypeNaturalPerson,
			},
		},
	}
	return &contact
}

func NewCompanyContact(accountId string) *Contact {
	contact := Contact{
		Data: ContactData{
			Type: ContactsType,
			Attributes: ContactAttributes{
				AccountID:    accountId,
				AccountRoles: []string{AccountRoleOwner},
				Type:         ContactTypeCompany,
			},
		},
	}
	return &contact
}

type ContactsResponse struct {
	CollectionResponse
	Data []ContactData `json:"data"`
}
