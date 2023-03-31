package main

import (
	"fmt"
	"paquete/mensajes"
	"paquete/models"

	figuras "github.com/alfonsomorales11/figurasgo"
)

func main() {

	mensajes.Saludar()
	mensajes.Hola()

	cuadrado1 := figuras.Cuadrado{10, 20}
	circulo1 := figuras.Circulo{10}

	figuras.CalculaAreaFigura(&cuadrado1)
	figuras.CalculaPerimetroFigura(&circulo1)

	persona := models.Persona{}
	persona.Constructor("pachi", 22)
	fmt.Println(persona)

	persona.SetNombre("Alfonso Morales Gutierrez")
	persona.SetEdad(26)

	fmt.Println(persona.GetEdad())

	fmt.Println(persona.GetNombre())

	fmt.Println(figuras.HolaMundo())

}

// go mod init paquete
