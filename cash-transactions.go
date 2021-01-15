package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cloudmode/go-primetrust/models"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func GetCashTransaction(transactionID string) (*models.CashTransaction, error) {
	apiUrl := fmt.Sprintf("%s/cash-transactions/%s", _apiPrefix, transactionID)
	color.Green("GetCashTransactions:%v", apiUrl)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.CashTransaction{}
	if err := json.Unmarshal(body, &response); err != nil {
		color.Red("body:%v", string(body))
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

// GetCashTransactionsPage returns a single page of transactions using the USD filter
func GetCashTransactionsPage(accountID string, number, size int64) (*models.CashTransactionsResponse, error) {
	filter := "filter[currency-type eq]=USD"
	apiURL := fmt.Sprintf("%s/cash-transactions?%s&page[number]=%d&page[size]=%d&include=account-cash-transfer-from,account-cash-transfer-to&sort=-created-at&account.id=%s",
		_apiPrefix, url.PathEscape(filter), number, size, accountID)
	color.Red("GetCashTransactionsPage:%v", apiURL)

	req, err := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		color.Red("error getting accountID:%s:%s", filter)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		color.Red("wetf:%v", res.Status)
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)
	//color.Red("body:%s", body)

	transactions := models.CashTransactionsResponse{}
	if err := json.Unmarshal(body, &transactions); err != nil {
		return nil, errors.New("unmarshal error")
	}
	//color.Red("GetCashTransactionsPage:%v", PrettyPrint(transactions.Data[0:size-1]))
	return &transactions, nil
}

// GetCashTransactions returns all cash transactions between from and to
func GetCashTransactions(accountID string, from, to time.Time) (*models.CashTransactionsResponse, error) {
	filter := fmt.Sprintf("filter[created-at gte]=%s&filter[created-at lte]=%s",
		from.Format(time.RFC3339), to.Format(time.RFC3339))
	color.Blue("filter:%s", filter)
	apiURL := fmt.Sprintf("%s/cash-transactions?%s&page[number]=1&page[size]=1000&include=account-cash-transfer-from,account-cash-transfer-to&sort=-created-at&account.id=%s",
		_apiPrefix, url.PathEscape(filter), accountID)

	//apiURL = url.QueryEscape(apiURL)

	req, err := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		color.Red("error getting accountID:%s:%s", filter)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		color.Red("wetf:%v", res.Status)
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)
	//color.Red("body:%s", body)

	transactions := models.CashTransactionsResponse{}
	if err := json.Unmarshal(body, &transactions); err != nil {
		return nil, errors.New("unmarshal error")
	}
	//color.Red("GetTransactions:%v", PrettyPrint(transactions.Data[0:]))
	return &transactions, nil
}

// NewFundsTransfer ...
func NewFundsTransfer(from, to, reference string, amount float64) *models.FundTransfer {
	attrs := models.FundTransferAttributes{FromAccountID: from, ToAccountID: to, Amount: amount, Reference: reference, CurrencyType: "USD"}
	data := models.FundTransferData{Type: "account-cash-transfers", Attributes: attrs}
	ft := models.FundTransfer{Data: data}
	return &ft
}

// GetFundsTransfer from funds-transfer-id
func GetFundsTransfer(fundsTransferID string) (*models.FundsTransfer, error) {
	apiUrl := fmt.Sprintf("%s/funds-transfers/%s", _apiPrefix, fundsTransferID)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.FundsTransfer{}
	if err := json.Unmarshal(body, &response); err != nil {
		color.Red("GetFundsTransfer unmarshal error:%v", err)
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

// FundsTransfer ...
// trying a JSON parsing trick from https://stackoverflow.com/questions/35583735/unmarshaling-into-an-interface-and-then-performing-type-assertion
// because Included is an array of CashTransfer and CashTransaction
func FundsTransfer(from, to, amount, reference string) (*models.CashTransfer, error) {
	tAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}
	ft := NewFundsTransfer(from, to, reference, tAmount)
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(ft)

	apiURL := fmt.Sprintf("%s/account-cash-transfers?include=to-account-cash-totals,from-account-cash-totals,to-cash-transaction,from-cash-transaction", _apiPrefix)
	//color.Green("FundsTransfer:%v", apiUrl)

	req, err := http.NewRequest("POST", apiURL, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("%s: %s", res.Status, string(body))
	}

	response := models.CashTransfer{}

	if err := json.Unmarshal(body, &response); err != nil {
		color.Black("FundsTransfer unmarshal error %v", string(body))
		return nil, err
	}

	included := gjson.ParseBytes(body).Get("included")
	for _, raw := range included.Array() {
		t := raw.Get("type")
		r := raw.Get("relationships.account.links.related")
		if strings.Contains(r.String(), from) {
			if t.String() == "account-cash-totals" {
				if err := json.Unmarshal([]byte(raw.Raw), &response.FromCashData); err != nil {
					color.Red("shit:%v", err)
				}
				//color.Green("CashTotal:%v", PrettyPrint(response.FromCashData))
			} else if t.String() == "cash-transactions" {
				if err := json.Unmarshal([]byte(raw.Raw), &response.FromCashTransaction); err != nil {
					color.Red("shit:%v", err)
				}
			}
		} else if strings.Contains(r.String(), to) {
			if t.String() == "account-cash-totals" {
				if err := json.Unmarshal([]byte(raw.Raw), &response.ToCashData); err != nil {
					color.Red("shit:%v", err)
				}
			} else if t.String() == "cash-transactions" {
				if err := json.Unmarshal([]byte(raw.Raw), &response.ToCashTransaction); err != nil {
					color.Red("shit:%v", err)
				}
			}
		}

	}
	//color.Red("FundsTransfer:response:%+v", PrettyPrint(response))

	return &response, nil
}

// PrettyPrint ...
func PrettyPrint(thing interface{}) string {
	s, _ := json.MarshalIndent(thing, "", "\t")
	return string(s)
}
