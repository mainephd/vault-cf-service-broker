package vaulthandlers

import "github.com/pivotal-cf/brokerapi"

func ProvisionNewService(spaceid string, instanceid string ) (brokerapi.ProvisionedServiceSpec,error){
	return brokerapi.ProvisionedServiceSpec{}, nil
}
