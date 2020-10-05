package models

type ContactFundsTransferReferences struct {
	Data ContactFundsTransferReferenceData `json:"data"`
}

type ContactFundsTransferReferenceData struct {
	ID            string                         `json:"id,omitempty"`
	Type          string                         `json:"type"`
	Attributes    ContactFundsTransferAttributes `json:"attributes"`
	Links         Links                          `json:"links"`
	Relationships Relationships                  `json:"relationships"`
}

type ContactFundsTransferAttributes struct {
	Reference string `json:"reference"`
	AccountID string `json:"account-id,omitempty"`
	ContactID string `json:"contact-id,omitempty"`
}
