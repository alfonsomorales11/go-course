package main

import "fmt"

// Relacion de uno a uno
type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

func main() {

	alfonso := User{"alfonso", "alfonso@gmail.com", true}
	//alfonso1 := User{"pachi", "pachi@gmail.com", true}

	student1 := Student{alfonso, "asd123"}

	fmt.Println(student1)

}
