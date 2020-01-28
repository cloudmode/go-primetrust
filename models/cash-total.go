package models

type CashTotal struct {
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
	Data  []CashData `json:"data"`
}

type CashData struct {
	ID            string         `json:"id,omitempty bson:"id,omitempty"`
	Type          string         `json:"type"`
	Attributes    CashAttributes `json:"attributes"`
	Links         Links          `json:"links"`
	Relationships Relationships  `json:"relationships"`
}

type CashAttributes struct {
	ContingentHold  float64 `json:"contingent-hold"`
	Disbursable     float64 `json:"disbursable"`
	PendingTransfer float64 `json:"pending-transfer"`
	Settled         float64 `json:"settled"`
	CurrencyType    string  `json:"currency-type,omitempty"`
	UpdatedAt       string  `json:"updated-at,omitempty"`
}
