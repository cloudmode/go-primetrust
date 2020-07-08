package models

import (
	"time"
)

type CashTransactionAttributes struct {
	ID                string    `json:"id,omitempty"`
	Amount            float64   `json:"amount"`
	Comments1         string    `json:"comments-1" bson:"comments-1"`
	Comments2         string    `json:"comments-2" bson:"comments-2"`
	Comments3         string    `json:"comments-3" bson:"comments-3"`
	Comments4         string    `json:"comments-4" bson:"comments-4"`
	CreatedAt         time.Time `json:"created-at"  bson:"created-at"`
	CurrencyType      string    `json:"currency-type"`
	CustomerReference string    `json:"customer-reference" bson:"customer-reference"`
	EffectiveAt       time.Time `json:"effective-at"  bson:"effective-at"`
	SettledOn         string    `json:"settled-on"`
}

type CashTransactionData struct {
	ID            string                    `json:"id,omitempty"`
	Type          string                    `json:"type"`
	Attributes    CashTransactionAttributes `json:"attributes"`
	Links         Links                     `json:"links"`
	Relationships Relationships             `json:"relationships"`
	Included      interface{}               `json:"included"`
}

type CashTransactionDataNoExtra struct {
	ID         string                    `json:"id,omitempty"`
	Type       string                    `json:"type"`
	Attributes CashTransactionAttributes `json:"attributes"`
	//Links         Links                     `json:"links"`
	//Relationships Relationships             `json:"relationships"`
	//Included      interface{}               `json:"included"`
}

type CashTransaction struct {
	Data CashTransactionData `json:"data"`
}

type CashTransactionsResponse struct {
	CollectionResponse
	Data     []CashTransactionDataNoExtra `json:"data"`
	Included []CashTransferDataNoExtra    `json:"included"`
}

type FundTransfer struct {
	Data FundTransferData `json:"data"`
}

type FundTransferData struct {
	Type       string                 `json:"type"`
	Attributes FundTransferAttributes `json:"attributes"`
}

type FundTransferAttributes struct {
	FromAccountID string  `json:"from-account-id,omitempty"`
	ToAccountID   string  `json:"to-account-id,omitempty"`
	Amount        float64 `json:"amount"`
	CurrencyType  string  `json:"currency-type"`
	Reference     string  `json:"reference"`
}
