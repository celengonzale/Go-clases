package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumnos) Detalle() {
	fmt.Println(a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

type StockProduct interface {
	price() float64
}

func main() {
	alumnaMati := Alumnos{
		"Matilde",
		"Arg",
		20202020,
		"01/02/2000",
	}
	alumnaMati.Detalle()

}
