package models

import (
	"encoding/json"
	"time"
)

const (
	WebhooksType = "webhook-configs"
)

type WebhooksResponse struct {
	CollectionResponse
	Data []WebhookData `json:"data"`
}

type WebhookAttribute struct {
	AccountID    string    `json:"account_id`
	URL          string    `json:"url,omitempty"`
	Enabled      bool      `json:"enabled,omitempty"`
	ContactEmail string    `json:"contact-email,omitempty"`
	SharedSecret string    `json:"shared-secret,omitempty"`
	CreatedAt    time.Time `json:"created-at,omitempty"`
	Failures     int       `json:"failures,omitempty"`
	UpdatedAt    time.Time `json:"updated-at,omitempty"`
}

type WebhookData struct {
	ID            string           `json:"id,omitempty"`
	Type          string           `json:"type"`
	Attributes    WebhookAttribute `json:"attributes"`
	Links         Links            `json:"links,omitempty"`
	Relationships Relationships    `json:"relationships,omitempty"`
}

type Webhook struct {
	Data WebhookData `json:"data"`
}

func NewWebhook(accountId string, url string, email string, secret string, enabled bool) *Webhook {
	webhook := Webhook{
		Data: WebhookData{
			Type: WebhooksType,
			Attributes: WebhookAttribute{
				Enabled:      enabled,
				URL:          url,
				AccountID:    accountId,
				ContactEmail: email,
				SharedSecret: secret,
			},
		},
	}
	return &webhook
}

type WebhookPayloadData struct {
	Changes    []string          `json:"changes,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type WebhookPayload struct {
	ID               string             `json:"id,omitempty"`
	AccountID        string             `json:"account_id,omitempty"`
	Action           string             `json:"action,omitempty"`
	CreatedAt        *time.Time         `json:"created_at,omitempty"`
	Data             WebhookPayloadData `json:"data,omitempty"`
	Failures         int                `json:"failures,omitempty"`
	LastRequestAt    *time.Time         `json:"last_request_at,omitempty"`
	LastResponseCode int                `json:"last_response_code,omitempty"`
	ResourceID       string             `json:"resource_id,omitempty"`
	ResourceType     string             `json:"resource_type,omitempty"`
	Success          bool               `json:"success,omitempty"`
}

func (w *WebhookPayload) ensureEncoded() ([]byte, error) {
	return json.Marshal(w)
}

func (w *WebhookPayload) Length() int {
	b, _ := w.ensureEncoded()
	return len(b)
}

func (w *WebhookPayload) Encode() ([]byte, error) {
	return w.ensureEncoded()
}

func (m *WebhookPayload) Decode(bytes []byte) error {
	err := json.Unmarshal(bytes, m)
	return err
}
