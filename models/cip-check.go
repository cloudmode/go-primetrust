package models

import "time"

type CipCheck struct {
	Links    Links          `json:"links"`
	Meta     Meta           `json:"meta"`
	Data     []CipCheckData `json:"data"`
	Included interface{}    `json:"included"`
}

type CipCheckData struct {
	ID            string             `json:"id,omitempty"`
	Type          string             `json:"type"`
	Attributes    CipCheckAttributes `json:"attributes"`
	Links         Links              `json:"links"`
	Relationships Relationships      `json:"relationships"`
}

type CipCheckAttributes struct {
	CreatedAt        time.Time `json:"created-at" bson:"created-at"`
	UpdatedAt        time.Time `json:"updated-at" bson:"updated-at"`
	ExceptionDetails string    `json:"exception-details"`
	Status           string    `json:"status"`
	Exceptions       []string  `json:"exceptions"`
}

type ContactException struct {
	AccountID  string   `json:"account-id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Number     string   `json:"number,omitempty"`
	ContactID  string   `json:"contact-id,omitempty"`
	Email      string   `json:"email,omitempty"`
	Details    string   `json:"details"`
	Exceptions []string `json:"exceptions"`
}
