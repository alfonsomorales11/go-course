package models

import (
	"fmt"
	"paquetes/db"
)

type Users struct {
	UsersList []User
}

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
}

func NewUser(username string, password string, email string) *User {
	user := &User{}
	user.UserName = username
	user.Password = password
	user.Email = email

	return user
}

func (user *User) insert() {
	sql := "insert users  SET username = ? , password = ?, email = ?"
	result, _ := db.Exec(sql, user.UserName, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func Insertar(username string, password string, email string) *User {
	user := NewUser(username, password, email)
	user.insert()

	return user
}

func ShowList() {

	rows, _ := db.Query("select * from users")

	users := Users{}

	for rows.Next() {
		user := User{}

		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)

		users.UsersList = append(users.UsersList, user)
	}

	fmt.Println("=========")
	fmt.Println(users)
	fmt.Println("=========")

}

func GetUser(id int) *User {

	user := User{}

	sql := fmt.Sprintf("select * from users where id=%d;", id)

	rows, error := db.Query(sql)

	if error != nil {
		panic(error)
	}

	for rows.Next() {

		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	}

	return &user
}

func UpdateUser(id int, newUserName string, newPassword string, newEmail string) {

	userToUpdate := GetUser(id)

	if newUserName == "" {
		newUserName = userToUpdate.UserName
	}

	if newPassword == "" {
		newPassword = userToUpdate.Password
	}

	if newEmail == "" {
		newEmail = userToUpdate.Email
	}

	userToUpdate.UserName = newUserName
	userToUpdate.Password = newPassword
	userToUpdate.Email = newEmail

	sql := fmt.Sprintf("update users set username='%s', password='%s', email='%s' where id=%d;", newUserName, newPassword, newEmail, id)

	db.Exec(sql)
}

func DeleteUser(id int) {
	sql := fmt.Sprintf("delete from users where id=%d", id)

	db.Exec(sql)
}

const UserSchema string = `CREATE TABLE users (
    Id int auto_increment primary key,
    username varchar(255) not null,
    password varchar(255) not null,
    email varchar(255)
);`
