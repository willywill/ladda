package handler

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/willywill/ladda/pkg/responder"
	"github.com/willywill/ladda/pkg/shipper"
)

func debugFile(file *multipart.FileHeader) {
	log.Printf("Filename: %+v\n", file.Filename)
	log.Printf("File Size: %+v\n", file.Size)
	log.Printf("MIME Header: %+v\n", file.Header)
}

func UploadFile(res http.ResponseWriter, req *http.Request) {
	// Get the max file size allowed to be uploaded
	maxFileSize := 10 << 20
	envMaxFileSize, ok := os.LookupEnv("MAX_FILE_SIZE")

	if !ok {
		log.Println("No max file size specified, using default - 10MB")
	} else {
		maxFileSize, _ = strconv.Atoi(envMaxFileSize)
	}

	// Parse the multipart request
	parseErr := req.ParseMultipartForm(int64(maxFileSize))

	if parseErr != nil {
		log.Println("Failed to parse multipart form data.")
		responder.
			Create(res).
			SetStatusWithError(http.StatusBadRequest, "Unable to parse the resource.")

		return
	}

	// Retrieve the file form-data
	file, handler, fileErr := req.FormFile("file")

	if fileErr != nil {
		log.Println("Error retrieving the file from form-data.")
		responder.
			Create(res).
			SetStatusWithError(http.StatusBadRequest, "Unable to retrieve the resource.")
		return
	}

	// Close the file when this function execution has completed
	defer file.Close()

	// Debug the file contents if the flag was set
	debugString, _ := os.LookupEnv("DEBUG")
	debug, _ := strconv.ParseBool(debugString)

	if debug {
		debugFile(handler)
	}

	// Read file contents to byte array
	fileBytes, readErr := ioutil.ReadAll(file)

	if readErr != nil {
		log.Println("Error reading parsed file contents.")
		responder.
			Create(res).
			SetStatusWithError(http.StatusInternalServerError, "Failed to read contents of file.")
		return
	}

	ext := filepath.Ext(handler.Filename)

	// Pass along file contents to shipper
	shipper.WriteFile(fileBytes, ext)

	responder.
		Create(res).
		SetStatusWithMessage(http.StatusOK, "Success!")
}
