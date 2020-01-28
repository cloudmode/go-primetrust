package main

import (
	"github.com/cloudmode/go-primetrust"
	"log"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	contactId := os.Getenv("PRIMETRUST_CONTACT_ID")
	label := "Passport"
	description := "{\"validFrom\": \"2010-01-01\", \"validTill\": \"2020-01-01\"}"

	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "/examples/documents/upload/empty.jpg")
	_, fileName := filepath.Split(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileHeader := multipart.FileHeader{
		Filename: fileName,
		Header: textproto.MIMEHeader{
			"Content-Type": []string{"image/jpeg"},
		},
	}

	response, err := primetrust.UploadDocument(file, fileHeader, contactId, label, description)
	if err != nil {
		log.Println("Error uploading new document:", err)
	}

	log.Println("Document uploaded OK:", response)
}
