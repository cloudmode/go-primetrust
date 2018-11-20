package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type AccountAttributes struct {
	CreatedAt time.Time `json:"created-at"`
	Name      string    `json:"name"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
}

type AccountData struct {
	Type          string            `json:"type"`
	ID            uuid.UUID         `json:"id"`
	Attributes    AccountAttributes `json:"attributes"`
	Links         Links             `json:"links"`
	Relationships Relationships     `json:"relationships"`
}

type Account struct {
	Data AccountData `json:"data"`
}

type AccountsResponse struct {
	CollectionResponse
	Data []AccountData `json:"data"`
}
