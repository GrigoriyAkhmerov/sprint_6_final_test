package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootEndpointHandler returns HTML from index.html file for root endpoint.
// Point 1.
func RootEndpointHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "text/html")
	fmt.Printf("Method: %s\n", r.Method)
	http.ServeFile(w, r, "./index.html")
}

// UploadEndpointHandler for endpoint /upload.
func UploadEndpointHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Printf("Method: %s\n", req.Method)
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
	convertedText, err := service.ConvertText(string(dataFromFile))

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
	_, err = res.Write([]byte(convertedText))
	if err != nil {
		http.Error(res, "writing to slice error", http.StatusInternalServerError)
		return
	}
}
