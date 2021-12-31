package main

import (
	"lib-client-server/client/admin"
	"lib-client-server/client/customer"
	"lib-client-server/client/main_opt"
	"lib-client-server/client/models"
)

type (
	ClientServer interface {
	}
)

var (
	pages   *PagesModule
	storage *StorageModule // authorization storage
	auth    *AuthModule
	user    *UserModule
	// book stotage
	clientStorage models.StorageInterface

	adminPages            *admin.PagesModule
	userPages             *customer.PagesModule
	adminBooks, userBooks models.BookInterface
)

func CreateModules() {
	auth = createAuthModule()
	user = createUserModule()
	pages = createPagesModule()
	storage = createStorageModule("localhost", "libDB", "clientAuthDBc")
	clientStorage = createBookStorageModule("localhost", "libDB", "clientBookDBc")

	//
	adminPages = createAdminPagesModule("admin")
	adminBooks = createAdminBookModule("localhost", "libDB")
	userPages = createUserPagesModule("user")
	userBooks = createUserBookModule("localhost", "libDB")
}

func createBookStorageModule(host, name, cname string) models.StorageInterface {
	return main_opt.CreateCustomerBookStorageModule(host, name, cname)
}

func createPagesModule() *PagesModule {
	return &PagesModule{}
}

func createUserModule() *UserModule {
	return UserModule{}.Create()
}

func createAuthModule() *AuthModule {
	return AuthModule{}.Create()
}

func createStorageModule(host, name, cname string) *StorageModule {
	return StorageModule{host: host, name: name, cname: cname}.CreateConnect()
}

//

func createAdminPagesModule(prefixPage string) *admin.PagesModule {
	return admin.CreateAdminPagesModule(prefixPage)
}
func createUserPagesModule(prefixPage string) *customer.PagesModule {
	return customer.CreateUserPagesModule(prefixPage)
}
func createAdminBookModule(host, name string) models.BookInterface {
	return admin.CreateAdminBookModule(host, name)
}
func createUserBookModule(host, name string) models.BookInterface {
	return customer.CreateUserBookModule(host, name)
}
