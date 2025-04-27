package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootEndpointHandler returns HTML from index.html file for root endpoint.
// Point 1.
func RootEndpointHandler(res http.ResponseWriter, req *http.Request) {
	indexData, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	res.Write(indexData)
}

// UploadEndpointHandler for endpoint /upload.
func UploadEndpointHandler(res http.ResponseWriter, req *http.Request) {

	// Point 2. Getting file from Form.
	file, _, err := req.FormFile("attach")
	if err != nil {
		http.Error(res, "file upload error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Point 3. Read data from file.
	dataFromFile, err := io.ReadAll(file)
	file.Close()
	if err != nil {
		http.Error(res, "file reading error", http.StatusInternalServerError)
		return
	}

	// Point 4. Sending data from file to ConvertText function from package service.
	convertedText := service.ConvertText(string(dataFromFile))

	// Point 5. Creating local file.
	localFileName := time.Now().UTC().String() + filepath.Ext("text.txt")

	localFile, err := os.Create(localFileName)
	if err != nil {
		http.Error(res, "creating file error", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	// Point 6. Writting result of converting text.
	_, err = io.WriteString(localFile, convertedText)
	if err != nil {
		http.Error(res, "writing to file error", http.StatusInternalServerError)
		return
	}
	res.Write([]byte(convertedText))
}
