package models

import (
	"paquetes/db"
)

type Users []User

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
}

func MigrateUser() {
	db.Database.AutoMigrate(User{})
}
