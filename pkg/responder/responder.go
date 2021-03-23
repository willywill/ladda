package responder

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type ResponseBuilder struct {
	SetStatus            func(status int)
	SetStatusWithMessage func(status int, message string)
	SetStatusWithError   func(status int, errorMessage string)
}

func setStatus(res http.ResponseWriter) func(int) {
	return func(status int) {
		res.WriteHeader(status)
	}
}

func setStatusWithMessage(res http.ResponseWriter) func(int, string) {
	return func(status int, message string) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(status)
		json.NewEncoder(res).Encode(ResponseBody{
			Message: message,
		})
	}
}

func setStatusWithError(res http.ResponseWriter) func(int, string) {
	return func(status int, errorMessage string) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(status)
		json.NewEncoder(res).Encode(ResponseBody{
			Error: errorMessage,
		})
	}
}

func Create(res http.ResponseWriter) ResponseBuilder {
	return ResponseBuilder{
		SetStatus:            setStatus(res),
		SetStatusWithMessage: setStatusWithMessage(res),
		SetStatusWithError:   setStatusWithError(res),
	}
}
