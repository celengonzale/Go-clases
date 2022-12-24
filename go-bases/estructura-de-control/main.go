package main

import (
	"fmt"
)

func main() {
	frutas := []string{"kiwi", "frutilla", "manzana"}
	for i, elem := range frutas {
		fmt.Println(i, elem)
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "es impar")
	}
	sum := 0
	for {
		sum++
		if sum >= 25 {
			break
		}
	}
	fmt.Println(sum)

	//Ejercicio de practica
	meses := []string{"Ene", "Cele", "Abr", "Ago", "Oct"}
	ObtenerEstaciones(meses)
	palabra := "Cele"
	RealAcademiaEspañola(palabra)
	Prestamo(30, true, 1, 350.000)
}

func ObtenerEstaciones(meses []string) {
	for _, mes := range meses {
		switch {
		case mes == "Ene" || mes == "Feb" || mes == "Mar":
			fmt.Println("Verano")
		case mes == "Abr" || mes == "May" || mes == "Jun":
			fmt.Println("Otoño")
		case mes == "Jul" || mes == "Ago" || mes == "Sep":
			fmt.Println("Invierno")
		case mes == "Oct" || mes == "Nov" || mes == "Dic":
			fmt.Println("Primavera")
		default:
			fmt.Println("Error")
		}
	}

}
func RealAcademiaEspañola(palabra string) {
	cantidadDeLetras := len(palabra)
	fmt.Println(cantidadDeLetras)
	for _, elem := range palabra {
		value := string(elem)
		fmt.Println(value)
	}
}
func Prestamo(edad int, empleado bool, antiguedad int, sueldo float32) {
	if edad >= 22 && empleado && antiguedad >= 1 {
		fmt.Println("Prestamo Success")
	} else {
		fmt.Println("Lo sentimos, no cumple con los requisitos")
	}

	switch {
	case edad < 22:
		fmt.Println("Usted no cumple con la edad")
	case empleado == false:
		fmt.Println("Usted no cumple con empleado")
	case antiguedad <= 0:
		fmt.Println("No cumple con la antiguedad")
	case sueldo < 100.000:
		fmt.Println("Prestamo con intereses ")
	case sueldo > 100.000:
		fmt.Println("Prestamo sin intereses")

	}
}
func AQueMesCorresponde(mes int) {
	
}
