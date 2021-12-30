package main

import (
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
	clientStorage models.StorageInterface

)

func CreateModules() {
	auth = createAuthModule()
	user = createUserModule()
	pages = createPagesModule()
	storage = createStorageModule("localhost", "libDB", "libraryDatabase")
	clientStorage = createBookStorageModule("localhost", "libDB", "libraryDatabase")
}

func createBookStorageModule(host, name, cname string) models.StorageInterface {
	return main_opt.CreateBookStorageModule(host, name, cname)
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
