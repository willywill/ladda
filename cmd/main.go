package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/willywill/ladda/pkg/handler"
	"github.com/willywill/ladda/pkg/middleware"
)

// TODO: Pass this as an environment variable
const PORT int = 3001

func main() {
	multiplexer := http.NewServeMux()

	uploadHandler := http.HandlerFunc(handler.UploadFile)

	multiplexer.Handle("/", middleware.Authenticate(uploadHandler))

	fmt.Printf("> Server started on port %v\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), multiplexer)

	if err != nil {
		log.Fatal(err)
	}
}
