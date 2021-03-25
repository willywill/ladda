package middleware

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAuthorize(t *testing.T) {
	t.Run("should NOT call 'next.ServeHTTP' if the request does not have the same auth token as the server", func(t *testing.T) {
		mockWasCalled := false

		mockHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			mockWasCalled = true
		})

		secret := "tok3n"
		err := os.Setenv("SECRET", secret)

		if err != nil {
			t.Fatal("Failed to set secret for test.")
		}

		handlerToTest := Authorize(mockHandler)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "http://testing", nil)

		req.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte("differentTokenString")))

		handlerToTest.ServeHTTP(res, req)

		if mockWasCalled {
			t.Fatal("Expected 'next.ServeHTTP' method to NOT be called, but it was called.")
		}
	})

	t.Run("should call 'next.ServeHTTP' if the request has the same auth token as the server", func(t *testing.T) {
		mockWasCalled := false

		mockHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			mockWasCalled = true
		})

		secret := "tok3n"
		err := os.Setenv("SECRET", secret)
		encodedToken := base64.StdEncoding.EncodeToString([]byte(secret))

		if err != nil {
			t.Fatal("Failed to set secret for test.")
		}

		handlerToTest := Authorize(mockHandler)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "http://testing", nil)

		req.Header.Set("Authorization", encodedToken)

		handlerToTest.ServeHTTP(res, req)

		if !mockWasCalled {
			t.Fatal("Expected 'next.ServeHTTP' method to be called, but it was not called.")
		}
	})
}
