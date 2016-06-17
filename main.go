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
	return []brokerapi.Service{}
}

func (*vaultServiceBroker) Provision(
instanceID string,
details brokerapi.ProvisionDetails,
asyncAllowed bool,
) (brokerapi.ProvisionedServiceSpec, error) {
	// Provision a new instance here. If async is allowed, the broker can still
	// chose to provision the instance synchronously
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (*vaultServiceBroker) LastOperation(instanceID string) (brokerapi.LastOperation, error) {
	// If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
	// for the status of the provisioning operation.
	// This also applies to deprovisioning (work in progress).
	return brokerapi.LastOperation{}, nil
}

func (*vaultServiceBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	// Deprovision a new instance here. If async is allowed, the broker can still
	// chose to deprovision the instance synchronously, hence the first return value.
	return brokerapi.IsAsync(false), nil
}

func (*vaultServiceBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	// Bind to instances here
	// Return a binding which contains a credentials object that can be marshalled to JSON,
	// and (optionally) a syslog drain URL.
	return brokerapi.Binding{}, nil
}

func (*vaultServiceBroker) Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	// Unbind from instances here
	return nil
}

func (*vaultServiceBroker) Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.IsAsync, error) {
	// Update instance here
	return brokerapi.IsAsync(false), nil
}

func main() {
	serviceBroker := &vaultServiceBroker{}
	logger := lager.NewLogger("vault-service-broker")
	credentials := brokerapi.BrokerCredentials{
		Username: os.Getenv("BROKER_USERNAME"),
		Password: os.Getenv("BROKER_PASSWORD"),
	}

	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}