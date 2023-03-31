package main

import "fmt"

func main() {
	var consumo, descuento float64
	var datosDescuento string

	fmt.Println("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			// Descuento
			datosDescuento = "10%"
			descuento = consumo * 0.10

		} else if consumo > 100 && consumo <= 200 {
			datosDescuento = "20%"
			descuento = consumo * 0.20

		} else if consumo > 200 {
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}

	} else {
		fmt.Println("Error al ingresar consumo")
	}

	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	fmt.Printf("El descuento es de %s , monto de descuento es %f , igv es %f y el monto a pagar es %f", datosDescuento, descuento, igv, totalPagar)

}
