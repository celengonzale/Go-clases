package main

import "fmt"

func main() {
	fullname := "Celena Gonzalez"
	address := "Laborde 9383"
	fmt.Printf("%s %s", fullname, address)
	fmt.Println()
	var (
		temp float32 = 30
		hum  int     = 43
		pres float32 = 1006.4
	)
	fmt.Println(temp)
	fmt.Println(hum)
	fmt.Println(pres)

	var a = true
	fmt.Println(a)
	//Ejercicio 4
	var edad int = 24
	boolean := false
	var sueldo string = "100"
	fmt.Println(edad)
	fmt.Println(boolean)
	fmt.Println(sueldo)

}
