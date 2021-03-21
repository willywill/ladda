package handler

import (
	"log"
	"net/http"
)

func UploadFile(w http.ResponseWriter, req *http.Request) {
	log.Println("Uploading File Handler Called")

	if req.Method != "POST" {
		// TODO: Send a response error
		log.Println("Unsupported HTTP method")
	}
}
