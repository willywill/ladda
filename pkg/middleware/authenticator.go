package middleware

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/willywill/ladda/pkg/responder"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Get the secret from the environment variables
		secret, _ := os.LookupEnv("SECRET")

		// Check the authorization header
		encodedAuthToken := req.Header.Get("Authorization")

		// Decode the authorization token, expected in base64
		authToken, err := base64.StdEncoding.DecodeString(encodedAuthToken)

		// Assume if the token is malformed, that the response is unauthorized anyway.
		// Check the secret against the token in the request, if we don't have a match we
		// send back an error in the response with a 401 status code
		if secret != string(authToken) || err != nil {
			responder.
				Create(res).
				SetStatusWithError(http.StatusUnauthorized, "You are not authorized to use this resource.")

			return
		}

		// Continue to the next middleware or handler
		next.ServeHTTP(res, req)
	})
}
