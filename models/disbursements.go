package models

import (
	"strconv"
	"time"

	"github.com/fatih/color"
)

// Output
type Disbursements struct {
	Data DisbursementsData `json:"data"`
}

type DisbursementsData struct {
	ID            string                  `json:"id,omitempty"`
	Type          string                  `json:"type"`
	Attributes    DisbursementsAttributes `json:"attributes"`
	Links         Links                   `json:"links"`
	Relationships Relationships           `json:"relationships"`
}

type DisbursementsAttributes struct {
	Amount            float64   `json:"amount"`
	CreatedAt         time.Time `json:"created-at"  bson:"created-at"`
	Description       string    `json:"description"`
	CustomerReference string    `json:"customer-reference" bson:"customer-reference"`
	CurrencyType      string    `json:"currency-type"`
	PaymentDetails    string    `json:"payment-details"`
	ReferenceNumber   string    `json:"reference-number"`
	SpecialType       string    `json:"special-type"`
	Status            string    `json:"status"`
	TransactionNumber string    `json:"transaction-number"`
}

// Input
type DisbursementsForm struct {
	Data DisbursementsDefinition `json:"data"`
}

type DisbursementsDefinition struct {
	Type       string                      `json:"type"`
	Attributes DisbursementsFormAttributes `json:"attributes"`
}

type DisbursementsFormAttributes struct {
	AccountID             string  `json:"account-id"`
	Amount                float64 `json:"amount"`
	Reference             string  `json:"reference,omitempty"`
	CurrencyType          string  `json:"currencty-type"`
	Description           string  `json:"description"`
	FundsTransferMethodID string  `json:"funds-transfer-method-id"`
}

func NewDisbursementsForm(method *FundsTransferMethod, pargs *map[string]string) *DisbursementsForm {
	args := *pargs

	f := DisbursementsForm{}

	var amount float64
	var err error
	if amount, err = strconv.ParseFloat(args["amount"], 64); err != nil {
		if amount, err = strconv.ParseFloat(args["Amount"], 64); err != nil {
			color.Red("Invalid amount in pargs:%v", args["Amount"])
			return nil
		}
	}

	f.Data.Type = "disbursements"
	f.Data.Attributes = DisbursementsFormAttributes{
		AccountID:             args["accountID"],
		Amount:                amount,
		Reference:             args["reference"],
		CurrencyType:          "USD",
		Description:           args["Note"],
		FundsTransferMethodID: method.Data.ID,
	}

	return &f
}

// Output
type FundsTransferMethod struct {
	Data FundsTransferData `json:"data"`
}

type FundsTransferData struct {
	ID            string                  `json:"id,omitempty"`
	Type          string                  `json:"type"`
	Attributes    FundsTransferAttributes `json:"attributes"`
	Links         Links                   `json:"links"`
	Relationships Relationships           `json:"relationships"`
	Included      interface{}             `json:"included"`
}

type FundsTransferAttributes struct {
	BankAccountName string `json:"bank-account-name"`
	BankAccountType string `json:"bank-account-type"`
	BankName        string `json:"bank-name"`
	ContactEmail    string `json:"contact-email"`
	ContactName     string `json:"contact-name"`
	RoutingNumber   string `json:"routing-number"`
}

// Input
type FundsTransferForm struct {
	Data FundsTransferDefinition `json:"data"`
}

type FundsTransferDefinition struct {
	Type       string                      `json:"type"`
	Attributes FundsTransferFormAttributes `json:"attributes"`
}

type FundsTransferFormAttributes struct {
	BankAccountName            string  `json:"bank-account-name"`
	BankAccountNumber          string  `json:"bank-account-number"`
	BankName                   string  `json:"bank-name"`
	ContactEmail               string  `json:"contact-email"`
	ContactName                string  `json:"contact-name"`
	RoutingNumber              string  `json:"routing-number"`
	FundsTransferType          string  `json:"funds-transfer-type"`
	FurtherCreditAccountName   string  `json:"further-credit-account-name"`
	FurtherCreditAccountNumber string  `json:"further-credit-account-number"`
	BeneficiaryAddress         Address `json:"beneficiary-address"`
}

func NewFundsTransferForm(pargs *map[string]string) *FundsTransferForm {
	f := FundsTransferForm{}
	f.Data.Type = "funds-transfer-method"
	args := *pargs

	color.Red("NewFundsTransferForm: \n\tpargs:%v \n\t args:%v", pargs, args)

	addr := Address{
		Street1:    args["Street"],
		City:       args["City"],
		Region:     args["State"],
		Country:    "US",
		PostalCode: args["Zipcode"],
	}

	f.Data.Attributes = FundsTransferFormAttributes{
		BeneficiaryAddress:         addr,
		ContactEmail:               args["email"],
		ContactName:                args["contactName"],
		BankName:                   args["DepositoryBankName"],
		BankAccountName:            args["AccountName"],
		BankAccountNumber:          args["AccountNumber"],
		RoutingNumber:              args["RoutingNumber"],
		FundsTransferType:          "wire",
		FurtherCreditAccountName:   "intendedUse",
		FurtherCreditAccountNumber: args["FurtherCredit"],
	}
	return &f
}
