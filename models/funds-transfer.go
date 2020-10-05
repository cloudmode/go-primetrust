package models

import "time"

type FundsTransfer struct {
	Data     FundsTransferData2 `json:"data"`
	Included []ContactData      `json:"included"`
}

type FundsTransferData2 struct {
	ID            string                   `json:"id,omitempty"`
	Type          string                   `json:"type"`
	Attributes    FundsTransferAttributes2 `json:"attributes"`
	Links         Links                    `json:"links"`
	Relationships Relationships            `json:"relationships"`
}

type FundsTransferAttributes2 struct {
	Amount                 float64   `json:"amount"`
	AmountExpected         float64   `json:"amount-expected"`
	CancelledAt            string    `json:"cancelled-at"`
	ClearsOn               string    `json:"clears-on"`
	CreatedAt              time.Time `json:"created-at"`
	ContingenciesClearedAt time.Time `json:"contingencies-cleared-at"`
	ContingenciesClearedOn string    `json:"contingencies-cleared-on"`
	CurrencyType           string    `json:"currency-type,omitempty"`
	EqualityHash           string    `json:"equality-hash"`
	FundsSourceName        string    `json:"funds-source-name"`
	FundsSourceType        string    `json:"funds-source-type"`
	PrivateMemo            string    `json:"private-memo"`
	Reference              string    `json:"reference"`
	ReversedAmount         float64   `json:"reversed-amount"`
	ReversalDetails        string    `json:"reversal-details"`
	ReversedAt             time.Time `json:"reversed-at"`
	SettledAt              time.Time `json:"settled-at"`
	SettlementDetails      string    `json:"settlement-details"`
	SignetDepositAddress   string    `json:"signet-deposit-address"`
	SpecialInstructions    string    `json:"special-instructions"`
	SpecialType            string    `json:"special-type"`
	Status                 string    `json:"status"`
	UpdatedAt              time.Time `json:"updated-at"`
	WireInstructions       string    `json:"wire-instructions"`
}
