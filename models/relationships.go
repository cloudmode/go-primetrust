package models

type RSpec struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Relationship struct {
	Links Links   `json:"links"`
	Data  []RSpec `json:"data"`
}

type Relationships struct {
	Accounts                       Relationship `json:"accounts,omitempty"`
	Account                        Relationship `json:"account,omitempty"`
	Contacts                       Relationship `json:"contacts,omitempty"`
	Contact                        Relationship `json:"contact,omitempty"`
	AccountType                    Relationship `json:"account-type,omitempty"`
	Addresses                      Relationship `json:"addresses,omitempty"`
	AMLChecks                      Relationship `json:"aml-checks,omitempty"`
	CIPChecks                      Relationship `json:"cip-checks,omitempty"`
	Contributions                  Relationship `json:"contributions,omitempty"`
	Currency                       Relationship `json:"currency,omitempty"`
	Disbursements                  Relationship `json:"disbursements,omitempty"`
	FromContactRelationships       Relationship `json:"from-contact-relationships,omitempty"`
	PaymentMethods                 Relationship `json:"payment-methods,omitempty"`
	PhoneNumbers                   Relationship `json:"phone-numbers,omitempty"`
	UploadedDocuments              Relationship `json:"uploaded-documents,omitempty"`
	RelatedFromContacts            Relationship `json:"related-from-contacts,omitempty"`
	RelatedToContacts              Relationship `json:"related-to-contacts,omitempty"`
	ToContactRelationships         Relationship `json:"to-contact-relationships,omitempty"`
	PrimaryAddress                 Relationship `json:"primary-address,omitempty"`
	PrimaryContact                 Relationship `json:"primary-contact,omitempty"`
	PrimaryPhoneNumber             Relationship `json:"primary-phone-number,omitempty"`
	AccountAssetTotals             Relationship `json:"account-asset-totals,omitempty"`
	AccountCashTotals              Relationship `json:"account-cash-totals,omitempty"`
	AccountQuestionnaire           Relationship `json:"account-questionnaire,omitempty"`
	AccountPolicy                  Relationship `json:"account-policy,omitempty"`
	AccountAggregatePolicy         Relationship `json:"account-aggregate-policy,omitempty"`
	AccountTransferAuthorizations  Relationship `json:"account-transfer-authorizations,omitempty"`
	AssetTransactions              Relationship `json:"asset-transactions,omitempty"`
	AssetTransfers                 Relationship `json:"asset-transfers,omitempty"`
	AssetTransferMethods           Relationship `json:"asset-transfer-methods,omitempty"`
	CashTransactions               Relationship `json:"cash-transactions,omitempty"`
	ContactFundsTransferReferences Relationship `json:"contact-funds-transfer-references,omitempty"`
	LatestAgreement                Relationship `json:"latest-agreement,omitempty"`
	Owners                         Relationship `json:"owners,omitempty"`
	Beneficiaries                  Relationship `json:"beneficiaries,omitempty"`
	Grantors                       Relationship `json:"grantors,omitempty"`
	OwnersAndGrantors              Relationship `json:"owners-and-grantors,omitempty"`
	FundsTransfers                 Relationship `json:"funds-transfers,omitempty"`
	Organization                   Relationship `json:"organization,omitempty"`
	WebhookConfig                  Relationship `json:"webhook-config,omitempty"`
}
