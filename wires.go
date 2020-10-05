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

// FundsTransferMethod ...
func FundsTransferMethod(method *models.FundsTransferForm) (*models.FundsTransferMethod, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(method)

	apiURL := fmt.Sprintf("%s/funds-transfer-methods?include=bank", _apiPrefix)

	color.Red("FundsTransferMethod:apiUrl:%v", apiURL)

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

	response := models.FundsTransferMethod{}

	if err := json.Unmarshal(body, &response); err != nil {
		color.Black("%v", string(body))
		return nil, err
	}

	color.Red("FundsTransferMethod:response:%+v", response)

	return &response, nil
}

func GetDisbursement(disbursementID string) (*models.Disbursements, error) {
	apiURL := fmt.Sprintf("%s/disbursements/%s?include=funds-transfer", _apiPrefix, disbursementID)
	req, err := http.NewRequest("GET", apiURL, nil)
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

	response := models.Disbursements{}
	if err := json.Unmarshal(body, &response); err != nil {
		color.Red("GetDisbursement:error:%v", err)
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func Disbursement(disburse *models.DisbursementsForm) (*models.Disbursements, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(disburse)

	apiURL := fmt.Sprintf("%s/disbursements", _apiPrefix)

	color.Red("Disbursement:apiUrl:%v", apiURL)

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

	response := models.Disbursements{}

	if err := json.Unmarshal(body, &response); err != nil {
		color.Black("%v", string(body))
		return nil, err
	}

	color.Red("Disbursement:response:%+v", response)

	return &response, nil
}

func FundsTransferReference(accountID, contactID string) (*models.ContactFundsTransferReferences, error) {
	ft := models.ContactFundsTransferReferences{}
	ft.Data.Attributes.AccountID = accountID
	ft.Data.Attributes.ContactID = contactID
	ft.Data.Type = "contact-funds-transfer-references"
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(ft)

	apiURL := fmt.Sprintf("%s/contact-funds-transfer-references", _apiPrefix)

	color.Red("FundsTransferReference:apiUrl:%v", apiURL)

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
	response := models.ContactFundsTransferReferences{}
	if err := json.Unmarshal(body, &response); err != nil {
		color.Black("%v", string(body))
		return nil, err
	}

	color.Red("FundsTransferReference:response:%+v", response)

	return &response, nil

	return nil, nil
}
