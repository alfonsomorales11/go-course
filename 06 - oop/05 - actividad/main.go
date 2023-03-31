package main

import (
	"fmt"
	"math"
)

// definir interface y los metodos que va a implementar
type figuraGeometrica interface {
	calculaArea() float32
	calculaPerimetro() float32
}

// definir estructuras
type Cuadrado struct {
	ancho  float32
	altura float32
}

type Circulo struct {
	radio float32
}

// Crear metodos para cada estructura
func (cuadrado *Cuadrado) calculaArea() float32 {
	return cuadrado.ancho * cuadrado.altura
}

func (cuadrado *Cuadrado) calculaPerimetro() float32 {
	return cuadrado.altura*2 + cuadrado.ancho*2
}

func (circulo *Circulo) calculaArea() float32 {
	return math.Pi * circulo.radio * circulo.radio
}

func (circulo *Circulo) calculaPerimetro() float32 {
	return math.Pi * circulo.radio * 2
}

// Definir funcion para la interface
func calculaAreaFigura(figura figuraGeometrica) {
	resultado := figura.calculaArea()
	fmt.Println(resultado)
}

func calculaPerimetroFigura(figura figuraGeometrica) {
	resultado := figura.calculaPerimetro()
	fmt.Println(resultado)
}

func main() {
	cuadrado1 := Cuadrado{10, 20}
	calculaAreaFigura(&cuadrado1)
	calculaPerimetroFigura(&cuadrado1)

	circulo1 := Circulo{10}
	calculaAreaFigura(&circulo1)
	calculaPerimetroFigura(&circulo1)
}
