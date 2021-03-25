package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("should not call 'next.ServeHTTP' if the request method is not 'POST'", func(t *testing.T) {
		mockWasCalled := false

		mockHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			mockWasCalled = true
		})

		handlerToTest := Validate(mockHandler)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "http://testing", nil)

		req.Header.Set("Content-Type", "multipart/form-data")

		handlerToTest.ServeHTTP(res, req)

		if mockWasCalled {
			t.Fatal("Expected 'next.ServeHTTP' method to NOT be called, but it was called.")
		}
	})

	t.Run("should not call 'next.ServeHTTP' if the request headers does not include 'multipart/form-data'", func(t *testing.T) {
		mockWasCalled := false

		mockHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			mockWasCalled = true
		})

		handlerToTest := Validate(mockHandler)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "http://testing", nil)

		req.Header.Set("Content-Type", "application/json")

		handlerToTest.ServeHTTP(res, req)

		if mockWasCalled {
			t.Fatal("Expected 'next.ServeHTTP' method to NOT be called, but it was called.")
		}
	})

	t.Run("should call 'next.ServeHTTP' when the request method is 'POST' and the 'Content-Type' contains 'multipart/form-data'", func(t *testing.T) {
		mockWasCalled := false

		mockHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			mockWasCalled = true
		})

		handlerToTest := Validate(mockHandler)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "http://testing", nil)

		req.Header.Set("Content-Type", "multipart/form-data")

		handlerToTest.ServeHTTP(res, req)

		if !mockWasCalled {
			t.Fatal("Expected 'next.ServeHTTP' method to be called, but it was not called.")
		}
	})
}
