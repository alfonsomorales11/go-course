package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Usuario struct {
	Name   string
	Age    int
	Activo bool
	Admin  bool
	Cursos []Curso
}

type Curso struct {
	Name string
}

func Saludar(nombre string) string {
	valor := fmt.Sprintf("Hola %s desde la función", nombre)
	return valor
}

func RenderTemplate(rw http.ResponseWriter, usuario any, documentName string) {
	err := templates.ExecuteTemplate(rw, documentName, usuario)

	if err != nil {
		Error(rw, 500)
	}
}

// global variable for template
var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.New("T").ParseGlob("templates/error/*.html"))

func Index(rw http.ResponseWriter, r *http.Request) {

	c1 := Curso{"Java"}
	c2 := Curso{"Go"}
	c3 := Curso{"Python"}

	cursos := []Curso{c1, c2, c3}

	usuario := Usuario{"alfonso", 26, true, true, cursos}

	RenderTemplate(rw, usuario, "index.html")
}

func Registro(rw http.ResponseWriter, r *http.Request) {

	RenderTemplate(rw, nil, "registro.html")
}

func Error(rw http.ResponseWriter, errorNumber int) {
	rw.WriteHeader(errorNumber)
	errorTemplate.Execute(rw, nil)

	RenderTemplate(rw, nil, "error.html")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

	fmt.Println("Servidor puerto 3000")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))

}
