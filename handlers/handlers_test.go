package vaulthandlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/mainephd/vault-cf-service-broker/handlers"
)

var _ = Describe("Handler", func(){
	Describe("Provision Tests", func(){
		It("Should call vault to create a new application ID and user Id", func(){
			response, err := vaulthandlers.ProvisionNewService("serviceid", "spaceid")
			Expect(err).ToNot(HaveOccurred())
			Expect(response).ToNot(Equal(nil))
			Expect(response.DashboardURL).To(Equal(""))
			Expect(response.IsAsync).To(Equal(false))
		})
	})
})