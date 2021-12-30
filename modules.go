package main

type (
	ClientServer interface {
	}
)

var (
	pages   *PagesModule
	storage *StorageModule
	auth    *AuthModule
	user    *UserModule
)

func CreateModules() {
	auth = createAuthModule()
	user = createUserModule()
	pages = createPagesModule()
	storage = createStorageModule("localhost", "libDB", "libraryDatabase")
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
