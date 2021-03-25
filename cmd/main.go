package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/willywill/ladda/pkg/handler"
	"github.com/willywill/ladda/pkg/middleware"
)

func main() {
	// Get the secret from the environment variables
	_, ok := os.LookupEnv("SECRET")

	if !ok {
		panic("No secret auth token was found in the environment.")
	}

	port := "3001"
	envPort, ok := os.LookupEnv("PORT")

	if !ok {
		log.Println("No port specified, using default - 3001")
	} else {
		port = envPort
	}

	multiplexer := http.NewServeMux()

	uploadHandler := http.HandlerFunc(handler.UploadFile)

	multiplexer.Handle("/api/v1/upload", middleware.Validate(middleware.Authorize(uploadHandler)))

	fmt.Printf("> Server started on port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), multiplexer)

	if err != nil {
		log.Fatal(err)
	}
}
