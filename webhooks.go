package primetrust

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudmode/go-primetrust/models"
	"github.com/fatih/color"
)

func CreateNewWebhook(webhook *models.Webhook) (*models.Webhook, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(webhook)

	apiUrl := fmt.Sprintf("%s/webhook-configs", _apiPrefix)
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

	response := models.Webhook{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func UpdateWebhook(webhook *models.Webhook) (*models.Webhook, error) {
	jsonData := new(bytes.Buffer)
	json.NewEncoder(jsonData).Encode(webhook)

	apiUrl := fmt.Sprintf("%s/webhook-configs/%s", _apiPrefix, webhook.Data.ID)
	req, err := http.NewRequest("PATCH", apiUrl, jsonData)
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
		return nil, errors.New(fmt.Sprintf("%s: %s", res.Status, string(body)))
	}

	response := models.Webhook{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func GetLastWebhook(accountId string) (*models.Webhook, error) {
	apiUrl := fmt.Sprintf("%s/webhook-configs?account.id=%s&one", _apiPrefix, accountId)
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

	response := models.Webhook{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func GetWebhook(webhookId string) (*models.WebhookPayload, error) {
	apiUrl := fmt.Sprintf("%s/webhooks/%s", _apiPrefix, webhookId)
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

	response := models.WebhookPayload{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}

func GetWebhookConfigs(accountID string) (*models.Webhook, error) {
	apiUrl := fmt.Sprintf("%s/webhook-configs?account.id=%s&one", _apiPrefix, accountID)
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		color.Red("GetwebhookConfigs error GETting:%v", apiUrl)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.Webhook{}
	if err := json.Unmarshal(body, &response); err != nil {
		color.Red("error unmarshalling:\n%v", string(body))
		color.Red("GetwebhookConfigs error GETting:%v", apiUrl)
		return nil, errors.New("unmarshal webhook payload error")
	}

	return &response, nil
}

func GetWebhookPayload(r *http.Request, secret string) (*models.WebhookPayload, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		color.Red("processWebhook:error:%v", err)
		return nil, err
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	webhookHMAC := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if r.Header.Get("X-Prime-Trust-Webhook-Hmac") != webhookHMAC {
		color.White("Hmac header didn't match")
		return nil, errors.New("X-Prime-Trust-Webhook-Hmac header is absent or not valid")
	}

	var webhookPayload models.WebhookPayload
	if err := json.Unmarshal(body, &webhookPayload); err != nil {
		color.White("couldn't decode webhook payload")
		return nil, errors.New("error decoding webhook payload")
	}
	color.Red("GetWebhookPayload:%v", string(body))
	return &webhookPayload, nil
}

func FromPrimeTrust(r *http.Request, body []byte, secret string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	webhookHMAC := base64.StdEncoding.EncodeToString(h.Sum(nil))
	//color.Blue("FromPrimeTrust getting headers....")
	if r.Header.Get("X-Prime-Trust-Webhook-Hmac") != webhookHMAC {
		return false
	}

	return true
}
