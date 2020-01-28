package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type AccountAttributes struct {
	CreatedAt          time.Time `json:"created-at" bson:"created-at"`
	Name               string    `json:"name"`
	Number             string    `json:"number"`
	Status             string    `json:"status"`
	DisbursmentsFrozen bool      `json:"disbursments-frozen"`
	OrganizationLabel  string    `json:"organization-label"`
	Statments          bool      `json:"statments"`
}

type AccountData struct {
	ID            string            `json:"id,omitempty bson:"id,omitempty"`
	Type          string            `json:"type"`
	Attributes    AccountAttributes `json:"attributes"`
	Links         Links             `json:"links"`
	Relationships Relationships     `json:"relationships"`
}

type Account struct {
	ID       bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Data     AccountData   `json:"data"`
	Included []ContactData `json:"included"`
}

type AccountsResponse struct {
	CollectionResponse
	Data []AccountData `json:"data"`
}

type AccountForm struct {
	ID   bson.ObjectId     `json:"-" bson:"_id,omitempty"`
	Data AccountDefinition `json:"data"`
}

type AccountDefinition struct {
	Type       string        `json:"type"`
	Attributes AccountInputs `json:"attributes"`
}

type AccountInputs struct {
	AccountType         string            `json:"account-type"`
	Name                string            `json:"name"`
	AuthorizedSignature string            `json:"authorized-signature"`
	Owner               ContactAttributes `json:"owner"`
	WebhookConfig       WebhookAttribute  `json:"webhook-config"`
	Licensed            bool              `json:"licensed`
	Submitted           bool              `json:"submitted`
}
