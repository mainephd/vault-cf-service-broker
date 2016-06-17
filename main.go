package main

import (
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"
	"net/http"
	"os"
)

type vaultServiceBroker struct {}

func (*vaultServiceBroker) Services() []brokerapi.Service {
	// Return a []brokerapi.Service here, describing your service(s) and plan(s)
}

func (*vaultServiceBroker) Provision(
instanceID string,
details brokerapi.ProvisionDetails,
asyncAllowed bool,
) (brokerapi.ProvisionedServiceSpec, error) {
	// Provision a new instance here. If async is allowed, the broker can still
	// chose to provision the instance synchronously.
}

func (*vaultServiceBroker) LastOperation(instanceID string) (brokerapi.LastOperation, error) {
	// If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
	// for the status of the provisioning operation.
	// This also applies to deprovisioning (work in progress).
}

func (*vaultServiceBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	// Deprovision a new instance here. If async is allowed, the broker can still
	// chose to deprovision the instance synchronously, hence the first return value.
}

func (*vaultServiceBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	// Bind to instances here
	// Return a binding which contains a credentials object that can be marshalled to JSON,
	// and (optionally) a syslog drain URL.
}

func (*vaultServiceBroker) Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	// Unbind from instances here
}

func (*vaultServiceBroker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	// Update instance here
}

func main() {
	serviceBroker := &vaultServiceBroker{}
	logger := lager.NewLogger("vault-service-broker")
	credentials := brokerapi.BrokerCredentials{
		Username: "username",
		Password: "password",
	}

	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}