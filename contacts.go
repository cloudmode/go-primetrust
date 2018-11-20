package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BANKEX/go-primetrust/models"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

func GetContacts() (*models.ContactsResponse, error) {
	apiUrl := fmt.Sprintf("%s/contacts", _apiPrefix)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _authHeader)

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
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func GetContact(contactId uuid.UUID) (*models.Contact, error) {
	apiUrl := fmt.Sprintf("%s/contacts/%s", _apiPrefix, contactId)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _authHeader)

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
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}

func CreateNewContact(contact *models.Contact) (bool, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(contact)

	apiUrl := fmt.Sprintf("%s/contacts", _apiPrefix)
	req, err := http.NewRequest("POST", apiUrl, jsonData)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", _authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.ContactsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return false, errors.New("Unmarshal error")
	}

	return true, nil
}
