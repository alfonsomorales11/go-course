package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El método es - ", r.Method)
	fmt.Fprintln(rw, "Hola mundo xD")
}

func PaginaNoFunciona(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "la pagina no funciona", 404)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Your name is %s and your age is %v", name, age)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PaginaNoFunciona)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	/*
		http.HandleFunc("/", Hola)
		http.HandleFunc("/page", PaginaNoFunciona)
		http.HandleFunc("/error", Error)
		http.HandleFunc("/saludar", Saludar)
	*/

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

	fmt.Println("Servidor puerto 3000")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))

}
