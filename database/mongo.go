package database

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"

	// "library/mgo.v2"
	"time"
)

type (
	dataBase struct {
		name    string
		session *mgo.Session
		err     error
		dbhost  string
		dname   string
		*mgo.Database
	}
)

//Connect - создает коннект в базу MONGODB
func Connect(dbhost, dbName, connectName string) *mgo.Database {
	base := dataBase{
		name:   connectName,
		dbhost: dbhost,
		dname:  dbName,
	}

	username := getDBUsername()
	password := getDBPassword()

	if username == "" {
		base.session, base.err = mgo.Dial(base.dbhost)
	} else {
		base.session, base.err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{base.dbhost},
			Timeout:  60 * time.Second,
			Username: username,
			Password: password,
		})
	}
	if base.err != nil {
		fmt.Println("database.go LocalConnect() ERROR CONNECTION:", base.err)
	}

	base.session.SetMode(mgo.Monotonic, true)

	go func() {
		for range time.Tick(time.Second * 1) {
			base.ConnectionCheck()
		}
	}()

	base.ConnectionCheck()

	fmt.Println("database.go() Connect(): успешный коннект на", base.dbhost, ";база:", base.dname, "; connectName:", connectName)
	return base.session.DB(base.dname)
}

// ConnectionCheck implements db reconnection
func (m *dataBase) ConnectionCheck() {
	if err := m.session.Ping(); err != nil {
		fmt.Println("Lost connection to db! name:", m.name)
		m.session.Refresh()
		if err := m.session.Ping(); err == nil {
			fmt.Println("Reconnect to db successful, name:", m.name)
		}
	}
}

func getDBUsername() string {
	dbUsername := os.Getenv("DB_USERNAME_MG")
	if len(dbUsername) == 0 {
		dbUsername = "" //возможно дефаул может быть другим
	}
	return dbUsername
}

func getDBPassword() string {
	dbPassword := os.Getenv("DB_PASSWORD_MG")
	if len(dbPassword) == 0 {
		dbPassword = "" //возможно дефаул может быть другим
	}
	return dbPassword
}
