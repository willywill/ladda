package middleware

import (
	"log"
	"net/http"
)

// TODO: Setup authentication by reading the auth headers and verifying the secret
// against the secret used in the environment variables
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")
		next.ServeHTTP(w, r)
	})
}
