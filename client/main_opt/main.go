package main_opt // todo rename to main after transferring

import (
	"lib-client-server/client/customer"
	"lib-client-server/client/models"
)

type (
)

//storage = createStorageModule("localhost", "libDB", "libraryDatabase")
func CreateCustomerBookStorageModule(host, name, cname string) models.StorageInterface {
	return customer.CreateConnect(host, name, cname)
}
