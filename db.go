package main

import (
	"fmt"
	"lib-client-server/database"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	StorageModule struct {
		host, name, cname string
		DB                *gorm.DB
		*mgo.Database
	}
)

func (s StorageModule) CreateConnect() *StorageModule {
	s.Database = database.Connect(s.host, s.name, s.cname)
	return &s
}

var dsn = "host=localhost user=postgres password=password dbname=gorm port=5432 sslmode=disable"

func (s StorageModule) CreateConnectPostgres() *StorageModule {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect err 2 = ", err)
	}
	s.DB = db
	return &s
}
