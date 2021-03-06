package primetrust

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudmode/go-primetrust/models"
	"github.com/fatih/color"
)

func GetCipCheck(contactId string) (*models.CipCheck, error) {
	apiUrl := fmt.Sprintf("%s/cip-checks?contact.id=%s", _apiPrefix, contactId)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _jwt)

	color.Red("GetCipCheck apiUrl:%v", apiUrl)

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

	response := models.CipCheck{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func SandboxApproveCipCheck(contactId string) (*models.CipCheckData, error) {
	// first get the cipcheck for contactId
	cc, err := GetCipCheck(contactId)
	if err != nil {
		return nil, err
	}

	// approve the cip_check_id
	apiUrl := fmt.Sprintf("%s/cip-checks/%s/sandbox/approve", _apiPrefix, cc.Data[0].ID)
	req, err := http.NewRequest("POST", apiUrl, nil)
	req.Header.Add("Authorization", _jwt)

	color.Red("SandboxApproveCipCheck apiUrl:%v", apiUrl)

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

	response := models.CipCheckData{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}
