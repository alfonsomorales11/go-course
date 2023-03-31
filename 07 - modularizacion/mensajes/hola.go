package mensajes

import "fmt"

func Saludar() {
	fmt.Println("Hola desde el paquete mensajes")
}

const mensaje = "Hola desde paquete xd"

func Hola() {
	fmt.Println(mensaje)
}
