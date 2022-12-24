package main

import (
	"fmt"
)

func main() {
	totalImpuesto := impuestosDeSalario(52000)
	fmt.Println(totalImpuesto, "impuestooo")
	promedio := promedioNotas(6, 5, 7)
	fmt.Println(promedio)
	salario := calcularSalario(2, "A")
	fmt.Println(salario)
	result, _ := operacion("minimum")
	resultdos, strdos := result(8.0, 4.0, 6.0)
	fmt.Println(resultdos, strdos, "EJERCICIO CUATRO")

	funcCantidad, str := animal("Tarántula")
	cantidadDeAlimeto := funcCantidad(1.0)
	fmt.Println(str)
	fmt.Println(cantidadDeAlimeto, "EJERCICIO CINCO")

}
func impuestosDeSalario(salario float64) float64 {
	switch {
	case salario > 50000 && salario <= 150000:
		return impuesto(salario, 17)
	case salario > 150000:
		return impuesto(salario, 27)
	default:
		return 0.0
	}
}
func impuesto(salario float64, porcentaje float64) float64 {
	return salario * porcentaje / 100
}
func promedioNotas(notas ...int) float64 {
	var resultadoNotas int
	for _, nota := range notas {
		if nota > 0 {
			resultadoNotas += nota
		} else {
			return float64(0)
		}
	}
	promedio := resultadoNotas / (len(notas))
	return float64(promedio)
}
func calcularSalario(horas float64, categoria string) float64 {
	switch categoria {
	case "A":
		return sueldoMensual(horas, 50, 3000)
	case "B":
		return sueldoMensual(horas, 20, 1500)
	case "C":
		return sueldoMensual(horas, 0, 1000)

	}
	return 0.0
}
func sueldoMensual(horas float64, porcentajeSueldoMensual float64, valorHora float64) (salario float64) {
	salario = horas * valorHora
	if porcentajeSueldoMensual > 0 {
		total := impuesto(salario, porcentajeSueldoMensual)
		return total + salario
	} else {
		return salario
	}
}

func minFunc(notas ...float64) (float64, error) {
	var min = notas[0]
	for _, nota := range notas {
		if min > nota {
			min = nota
		}
	}
	fmt.Println(min)
	return min, nil
}

func maxFunc(notas ...float64) (float64, error) {
	var max = notas[0]
	for _, nota := range notas {
		if max < nota {
			max = nota
		}
	}
	fmt.Println(max)
	return max, nil
}

func avgFunc(notas ...float64) (float64, error) {
	var result float64
	for _, nota := range notas {
		if nota > 0 {
			result += nota
		}
	}
	result = result / float64(len(notas))
	fmt.Println(result)
	return result, nil

}

type calculoFun func(...float64) (float64, error)

//type nico func(string2 string) int

const (
	minimum = "minimum"
	avg     = "average"
	maximum = "maximum"
)

func operacion(operation string) (calculoFun, string) {
	switch operation {
	case minimum:
		return minFunc, "Success"
	case maximum:
		return maxFunc, "Success"
	case avg:
		return avgFunc, "Success"
	default:
		return nil, "error"

	}
}

type cantidadDeAlimento func(float64) float64

const (
	perro     = "Perro"
	gato      = "Gato"
	hamster   = "Hamster"
	tarantula = "Tarántula"
)

func animal(tipoAnimal string) (cantidadDeAlimento, string) {
	switch tipoAnimal {
	case perro:
		return animalPerro, "PERRO"
	case gato:
		return animalGato, "GATO"
	case hamster:
		return animalHamster, "HAMSTER"
	case tarantula:
		return animalTarantula, "TARÁNTULA"
	default:
		return nil, "No existe"
	}
}
func animalPerro(cantidad float64) float64 {
	result := cantidad * 10.0
	return result
}
func animalGato(cantidad float64) float64 {
	result := cantidad * 5.0
	return result
}

func animalHamster(cantidad float64) float64 {
	result := cantidad * 2.50
	return result
}
func animalTarantula(cantidad float64) float64 {
	result := cantidad * 1.50
	return result
}
