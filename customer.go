package main

import (
	"fmt"
	"lib-client-server/client/models"
	"time"
)

type (
	UserModule struct{}

	AuthData struct {
		Login    string `form:"login" bson:"login"`
		Password string `form:"password" bson:"password"`
	}
)

func (u UserModule) Create() *UserModule {
	return &u
}

func (u *UserModule) CreateUser(data Registration, id uint64) error {
	if err := data.CheckEmail(); err != nil {
		return err
	}

	password, err := createHashPassword(data.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Id:           id,
		TicketNumber: fmt.Sprint(id, "-", data.Faculty, "-", time.Now().Unix()),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		SurName:      data.SurName,
		Faculty:      data.Faculty,
		Email:        data.Email,
		Password:     password,
	}

	if err := storage.C("users").Insert(user); err != nil {
		return err
	}

	return nil
}

func (r *Registration) CheckEmail() error {
	count, err := storage.C("users").Find(obj{"email": r.Email}).Count()
	switch {
	case err != nil:
		break
	case count > 0:
		err = fmt.Errorf("already registered")
	}

	return err
}
