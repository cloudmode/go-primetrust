package models

type Agreement struct {
	Data AgreementData `json:"data,omitempty"`
}
type AgreementData struct {
	ID         string              `json:"id,omitempty"`
	Type       string              `json:"type,omitempty"`
	Attributes AgreementAttributes `json:"attributes"`
}

type AgreementAttributes struct {
	Content string `json:"content,omitempty"`
	FileURL string `json:"file-url,omitempty"`
}
