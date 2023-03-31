package main

import "fmt"

type Animal interface {
	mover() string
}

type Perro struct{}

type Pez struct{}

type Pajaro struct{}

func (perro *Perro) mover() string {
	return "Soy perro y camino"
}

func (pez *Pez) mover() string {
	return "Soy pez y nado"
}

func (pajaro *Pajaro) mover() string {
	return "Soy pajaro y vuelo"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {
	perro := Perro{}
	moverAnimal(&perro)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)

	pez := Pez{}
	moverAnimal(&pez)

}
