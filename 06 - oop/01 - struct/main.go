package main

import "fmt"

// struct persona
type Persona struct {
	nombre string
	edad   int
}

type Empleado struct {
	// Herencia
	Persona

	sueldo int
}

// metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s \nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(newNombre string) {
	p.nombre = newNombre
}

func (p *Persona) editarEdad(newEdad int) {
	p.edad = newEdad
}

func main() {
	p1 := Persona{"Alfonso", 26}

	fmt.Println(p1)

	// p1.nombre = "Pachimix"
	p1.editarNombre("Alfonso Morales Gutierrez")
	p1.editarEdad(100)

	p1.imprimir()

	p2 := Persona{
		nombre: "Nina",
		edad:   50,
	}

	// fmt.Println(p2)
	p2.imprimir()

	empleado1 := Empleado{}

	empleado1.nombre = "Faith"
	empleado1.edad = 10
	empleado1.sueldo = 10000
	empleado1.imprimir()

	fmt.Println(empleado1)
}
