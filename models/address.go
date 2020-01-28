package models

const (
	AddressTypeHome   = "home"
	AddressTypeOffice = "office"
)

type Address struct {
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postal-code,omitempty"  bson:"postal-code"`
	Region     string `json:"region,omitempty"`
	Street1    string `json:"street-1"`
	Street2    string `json:"street-2,omitempty"`
}
