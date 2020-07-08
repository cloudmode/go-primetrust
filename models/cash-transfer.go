package models

import (
	"encoding/json"
	"time"
)

type TypedJson struct {
	Type string
	Data json.RawMessage
}

type CashTransferAttributes struct {
	Amount          float64       `json:"amount"`
	CreatedAt       time.Time     `json:"created-at"`
	UpdatedAt       time.Time     `json:"updated-at"`
	CurrencyType    string        `json:"currency-type,omitempty"`
	Reference       string        `json:"reference,omitempty"`
	Status          string        `json:"status,omitempty"`
	ReversalDetails string        `json:"reversal-details,omitempty"`
	Included        []interface{} `json:"included"`
}

type CashTransferData struct {
	ID            string                 `json:"id,omitempty"`
	Type          string                 `json:"type"`
	Attributes    CashTransferAttributes `json:"attributes"`
	Links         Links                  `json:"links"`
	Relationships Relationships          `json:"relationships"`
}

type CashTransfer struct {
	Data                CashTransferData `json:"data"`
	Included            []TypedJson      `json:"included"`
	FromCashData        CashData
	ToCashData          CashData
	FromCashTransaction CashTransactionData
	ToCashTransaction   CashTransactionData
}

type CashTransferDataNoExtra struct {
	ID         string                 `json:"id,omitempty"`
	Type       string                 `json:"type"`
	Attributes CashTransferAttributes `json:"attributes"`
	//Links         Links                  `json:"links"`
	//Relationships Relationships          `json:"relationships"`
}
