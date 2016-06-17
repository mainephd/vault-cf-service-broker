package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVaultCfServiceBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VaultCfServiceBroker Suite")
}
