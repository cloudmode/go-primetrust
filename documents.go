package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"path"

	"github.com/cloudmode/go-primetrust/models"
	"github.com/fatih/color"
	"github.com/globalsign/mgo"
)

func UploadDocument(file *mgo.GridFile, fileHeader multipart.FileHeader, contactId string, label string, description string) (*models.DocumentResponse, error) {
	apiUrl := fmt.Sprintf("%s/uploaded-documents", _apiPrefix)

	filename := fileHeader.Filename
	fileExtension := path.Ext(filename)
	contentType := fileHeader.Header.Get("Content-Type")

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"contact-id":  contactId,
		"description": description,
		"label":       label,
		"extension":   fileExtension,
		"mime_type":   contentType,
	}

	color.Green("UploadDocument data:%v", PrettyPrint(data))
	for key, val := range data {
		_ = writer.WriteField(key, val.(string))
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", _jwt)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyResp, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusCreated {
		log.Println(string(bodyResp))
		return nil, err
	}

	response := models.DocumentResponse{}
	if err := json.Unmarshal(bodyResp, &response); err != nil {
		return nil, errors.New("unmarshal error")
	}

	return &response, nil
}
