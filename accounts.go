package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudmode/go-primetrust/models"
	"github.com/fatih/color"
)

func GetAccounts() (*models.AccountsResponse, error) {
	apiUrl := fmt.Sprintf("%s/accounts?include=contacts", _apiPrefix)
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

	response := models.AccountsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func GetAccount(accountId string) (*models.Account, error) {
	apiUrl := fmt.Sprintf("%s/accounts/%s?include=contacts", _apiPrefix, accountId)
	color.Green("GetAccount:%v", apiUrl)
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

	response := models.Account{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func GetCashTotals(accountId string) (*models.CashTotal, error) {
	apiUrl := fmt.Sprintf("%s/account-cash-totals?account.id=%s", _apiPrefix, accountId)
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
	//color.Red("GetCashTotals:body:%v", string(body))

	response := models.CashTotal{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func CreateNewAccount(account *models.AccountForm) (*models.Account, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(account)

	apiUrl := fmt.Sprintf("%s/accounts?include=contacts", _apiPrefix)

	color.Red("CreateNewAccount:apiUrl:%v", apiUrl)

	req, err := http.NewRequest("POST", apiUrl, jsonData)
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
		return nil, errors.New(fmt.Sprintf("%s: %s", res.Status, string(body)))
	}

	response := models.Account{}

	if err := json.Unmarshal(body, &response); err != nil {
		color.Black("%v", string(body))
		return nil, err
	}

	color.Red("CreateNewAccount:response:%+v", response)

	return &response, nil
}

func GetAgreement(agreementId string) (*models.Agreement, error) {
	apiUrl := fmt.Sprintf("%s/agreements/%s", _apiPrefix, agreementId)
	color.Green("GetAgreement apiURL:%v", apiUrl)
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

	response := models.Agreement{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}
	color.Red("GetAgreement:%v, %+v", apiUrl, response)
	return &response, nil
}

func GetAgreementPreview(inputs *models.AgreementForm) (*models.Agreement, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(inputs)

	apiURL := fmt.Sprintf("%s/agreement-previews", _apiPrefix)

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

	response := models.Agreement{}

	if err := json.Unmarshal(body, &response); err != nil {
		color.Black("%v", string(body))
		return nil, err
	}

	// color.Red("GetAgreementPreview:response:%+v", response)

	return &response, nil
}

func SandboxAccountOpen(accountId string) (*models.AccountData, error) {
	apiURL := fmt.Sprintf("%s/accounts/%s/sandbox/open", _apiPrefix, accountId)
	req, err := http.NewRequest("POST", apiURL, nil)
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		color.Red("SandboxAccountOpen: error:%v %v", apiURL, err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		color.Red("SandboxAccountOpen: error:%v %v", apiURL, errors.New(res.Status))

		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.AccountData{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func SandboxAccountFund(accountId string, amount int) (*models.CashTransaction, error) {
	cash := models.CashTransaction{}
	cash.Data.Type = "accounts"
	cash.Data.Attributes.Amount = float64(amount)
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(cash)

	apiUrl := fmt.Sprintf("%s/accounts/%s/sandbox/fund", _apiPrefix, accountId)
	req, err := http.NewRequest("POST", apiUrl, jsonData)
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
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}
