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
	color.Red("GetCashTotals:body:%v", string(body))

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
