package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error"`
}

func UploadFile(w http.ResponseWriter, req *http.Request) {
	log.Println("Uploading File Handler Called")

	if req.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ResponseError{
			Error: "This HTTP method is not supported.",
		})
		log.Println("Unsupported HTTP method")
	}
}
