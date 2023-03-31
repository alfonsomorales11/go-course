package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:root@tcp(localhost:3306)/goweb"

var db *sql.DB

func Connect() {
	connection, error := sql.Open("mysql", url)

	if error != nil {
		panic(error)
	}

	fmt.Println("Successful connection")
	db = connection

}

func Close() {
	db.Close()
}

func Ping() {
	error := db.Ping()

	if error != nil {
		fmt.Println("Ping error")
	}
}

// Crear una tabla
func CreateTable(schema string, tablename string) {
	result := TableExists(tablename)

	if !result {
		fmt.Println("The table doesn't exist yet")
		Exec(schema)
	} else {
		fmt.Println("The table already exists")
	}
}

// verificar si una tabla existe
func TableExists(tablename string) bool {
	sql := fmt.Sprintf("show tables like '%s'", tablename)

	rows, error := Query(sql)

	if error != nil {
		fmt.Println(error)
	}
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, error := db.Exec(query, args...)

	if error != nil {
		fmt.Println(error)
	}

	return result, error
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, error := db.Query(query, args...)
	if error != nil {
		fmt.Println(error)
	}

	return rows, error
}

func TruncateTable(name string) {
	sql := fmt.Sprintf("truncate %s", name)
	Exec(sql)
}
