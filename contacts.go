package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cloudmode/go-primetrust/models"
	"github.com/fatih/color"
)

func GetContacts() (*models.ContactsResponse, error) {
	apiUrl := fmt.Sprintf("%s/contacts", _apiPrefix)
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

	response := models.ContactsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func GetContactByEmail(email string) (*models.Contacts, error) {
	//email = url.QueryEscape(email)
	filter := url.QueryEscape("filter[email eq]")
	email = url.QueryEscape(email)
	apiURL := fmt.Sprintf("%s/contacts?include=account&%s=%s", _apiPrefix, filter, email)
	//apiURL := fmt.Sprintf("%s/contacts?include=account&filter%5Bemail+eq%5D=stan%40f.co")
	color.Green("GetContactByEmail apiURL:%v", apiURL)
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

	response := models.Contacts{}
	if err := json.Unmarshal(body, &response); err != nil {
		color.Red("GetContactByEmail unmarshall error:%v", string(body))
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func GetContact(contactId string) (*models.Contact, error) {
	apiUrl := fmt.Sprintf("%s/contacts/%s", _apiPrefix, contactId)
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

	response := models.Contact{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func CreateNewContact(contact *models.Contact) (*models.Contact, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(contact)

	apiUrl := fmt.Sprintf("%s/contacts", _apiPrefix)
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

	response := models.Contact{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func toContact(contactID string, rcd *models.RelatedContactData) *models.Contact {
	contact := models.Contact{}
	contact.Data.Type = "contacts"
	contact.Data.Attributes = models.ContactAttributes{
		ContactType: rcd.ContactType,
		DateOfBirth: rcd.DateOfBirth,
		Email:       rcd.Email,
		Name:        rcd.Name,
		Sex:         rcd.Sex,
		TaxIDNumber: rcd.TaxIDNumber,
		TaxCountry:  rcd.TaxCountry,
		TaxState:    rcd.TaxState,
	}
	contact.Data.Attributes.PrimaryPhoneNumber = rcd.PrimaryPhoneNumber
	contact.Data.Attributes.PrimaryAddress = rcd.PrimaryAddress
	return &contact
}

func UpdateContact(contactID string, contact *models.RelatedContactData) (*models.Contact, error) {
	cData := toContact(contactID, contact)
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(cData)

	apiURL := fmt.Sprintf("%s/contacts/%s", _apiPrefix, contactID)
	req, err := http.NewRequest("PATCH", apiURL, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		color.Green("UpdateContact: statusCode:%v", res.StatusCode)
		color.Red("UpdateContact error response for %v: %v", apiURL, string(body))
		return nil, fmt.Errorf("%s: %s", res.Status, string(body))
	}

	response := models.Contact{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}
