package main_test

import (
	"net/http"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"

	"github.com/pivotal-cf/brokerapi"
)

var _ = Describe("Service Broker API", func() {
	var credentials = brokerapi.BrokerCredentials{
		Username: "username",
		Password: "password",
	}

	var brokerAPI http.Handler

	Describe("response headers", func() {
		makeRequest := func() *httptest.ResponseRecorder {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/v2/catalog", nil)
			request.SetBasicAuth(credentials.Username, credentials.Password)
			brokerAPI.ServeHTTP(recorder, request)
			return recorder
		}

		It("has a Content-Type header", func() {
			response := makeRequest()
			header := response.Header().Get("Content-Type")
			Ω(header).Should(Equal("application/json"))
		})
	})
})