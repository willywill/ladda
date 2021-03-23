package middleware

import (
	"net/http"
	"strings"

	"github.com/willywill/ladda/pkg/responder"
)

func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Only allow a POST method type for this handler
		if req.Method != "POST" {
			responder.
				Create(res).
				SetStatusWithError(http.StatusMethodNotAllowed, "HTTP method is not supported.")

			return
		}

		// Must include header with content type specifying multipart/form-data
		contentType := strings.ToLower(req.Header.Get("Content-Type"))

		if !strings.Contains(contentType, "multipart/form-data") {
			responder.
				Create(res).
				SetStatusWithError(http.StatusUnsupportedMediaType, "Content-Type is not supported.")

			return
		}

		// Continue to the next middleware or handler
		next.ServeHTTP(res, req)
	})
}
