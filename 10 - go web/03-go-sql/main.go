package main

import (
	"fmt"
	"paquetes/db"
	"paquetes/models"
)

func main() {
	// Connect to database
	db.Connect()

	// Create table
	db.CreateTable(models.UserSchema, "users")

	// Insert users
	models.Insertar("alfonso", "contraseña", "alfonso96xd@gmail.com")
	models.Insertar("luneta", "contraseñaluneta", "luneta@gmail.com")
	models.Insertar("paca", "paquita", "paca@gmail.com")
	models.Insertar("tia mary", "maryxd", "mary@gmail.com")
	models.Insertar("pachis", "pachimix", "pachis@gmail.com")

	// Update user by id
	models.UpdateUser(4, "tia mary gutierrez", "2299842250", "")

	// Show the list of users
	models.ShowList()

	// Delete user by id
	models.DeleteUser(3)

	// Get user by id
	var id int
	fmt.Print("Give the id to show user data: ")
	fmt.Scan(&id)
	fmt.Println(models.GetUser(id))

	var response string
	fmt.Print("Delete?: ")
	fmt.Scan(&response)

	// Delete all entries users table
	if response == "yes" {
		db.TruncateTable("users")
	}

	// Close connection to database
	db.Close()
}
