package models

type AccountTypeAttributes struct {
	Description          string   `json:"description"`
	Label                string   `json:"label"`
	ManageInvestments    bool     `json:"manage-investments" bson:"manage-investments"`
	OwnerRole            string   `json:"owner-role" bson:"owner-role"`
	StatementRoles       []string `json:"statement-roles" bson:"statement-roles"`
	TaxFormRecipientRole string   `json:"tax-form-recipient-role" bson:"tax-form-recipient"`
	Trust                bool     `json:"trust"`
}

type AccountTypeData struct {
	ID            string                `json:"id,omitempty"`
	Type          string                `json:"type"`
	Attributes    AccountTypeAttributes `json:"attributes"`
	Links         Links                 `json:"links"`
	Relationships Relationships         `json:"relationships"`
}

type AccountType struct {
	Data AccountTypeData `json:"data"`
}

type AccountTypesResponse struct {
	CollectionResponse
	Data     []AccountTypeData `json:"data"`
	Included interface{}       `json:"included"`
}
