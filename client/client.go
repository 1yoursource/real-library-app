package client

type (
	AuthInterface interface {
		Login()
		Logout()
		Register()
	}

	PageInterface interface {
		Main()
		Auth()
		Search()
		Shell()
		About()
	}

	DBInterface interface {
		Insert()
		Update()
		Find()
	}

	UserInterface interface {
		Create()
		Block()
		Info()
		Update() // недоступно для админа
	}
)
