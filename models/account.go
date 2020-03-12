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
	ID            string            `json:"id" bson:"id,omitempty"`
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
	ID   bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Data AccountDefinition `json:"data"`
}

type AccountDefinition struct {
	Type       string        `json:"type"`
	ID         string        `json:"id"`
	Attributes AccountInputs `json:"attributes"`
}

type AccountInputs struct {
	AccountType         string               `json:"account-type"`
	Name                string               `json:"name"`
	AuthorizedSignature string               `json:"authorized-signature"`
	Owner               ContactAttributes    `json:"owner"`
	Licensed            bool                 `json:"licensed"`
	Submitted           bool                 `json:"submitted"`
	WebhookConfig       WebhookAttribute     `json:"webhook-config"`
	Questionnaire       AccountQuestionnaire `json:"account-questionnaire,omitempty"`
}

type AccountQuestionnaire struct {
	Nature       string `json:"nature-of-business-of-the-company,omitempty"`
	Purpose      string `json:"purpose-of-account,omitempty"`
	Assets       string `json:"anticipated-types-of-assets,omitempty"`
	Source       string `json:"source-of-assets-and-income,omitempty"`
	Use          string `json:"intended-use-of-account,omitempty"`
	Volume       string `json:"anticipated-monthly-cash-volume,omitempty"`
	Incoming     string `json:"anticipated-monthly-transactions-incoming,omitempty"`
	Outgoing     string `json:"anticipated-monthly-transactions-outgoing,omitempty"`
	Patterns     string `json:"anticipated-trading-patterns,omitempty"`
	Associations string `json:"associations-with-other-accounts,omitempty"`
}

type AgreementForm struct {
	ID   bson.ObjectId       `json:"-" bson:"_id,omitempty"`
	Data AgreementDefinition `json:"data"`
}

type AgreementDefinition struct {
	Type       string          `json:"type"`
	Attributes AgreementInputs `json:"attributes"`
}

type AgreementInputs struct {
	AccountType         string            `json:"account-type"`
	Name                string            `json:"name"`
	AuthorizedSignature string            `json:"authorized-signature"`
	Owner               ContactAttributes `json:"owner"`
}
