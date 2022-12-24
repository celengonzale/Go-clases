package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImpuestoDeSalarioPorEncimaDelMinimo(t *testing.T) {
	//Arrange
	salario := float64(52000)
	//Act
	result := impuestosDeSalario(salario)
	//Assert
	expectedValue := float64(8840)
	assert.Equal(t, expectedValue, result)
}

func TestImpuestoDeSalarioPorDebajoDelMinimo(t *testing.T) {
	//Arrange
	salario := float64(33000)
	//Act
	result := impuestosDeSalario(salario)
	//Assert
	expectedValue := float64(0)
	fmt.Println(result)
	assert.Equal(t, expectedValue, result)
}
func TestImpuestoDeSalarioPorEncimaDelMaximo(t *testing.T) {
	//Arrange
	salario := float64(400000)
	//Act
	result := impuestosDeSalario(salario)
	//Assert
	expectedValue := float64(108000)
	fmt.Println(result)
	assert.Equal(t, expectedValue, result)
}
func TestPromedioDeNotas(t *testing.T) {
	//Arrange
	notas := []int{6, 4, 4}
	//Act
	promedio := promedioNotas(notas...)
	//Assert
	fmt.Println(promedio, "Testing promedio")
	expectedValue := float64(4)
	assert.Equal(t, expectedValue, promedio)
}
func TestPromedioDeNotasNegativas(t *testing.T) {
	//Arrange
	notas := []int{3, 7, -1}
	//Act
	promedio := promedioNotas(notas...)
	//Assert
	fmt.Println(promedio, "Con negativos")
	expectedValue := 0.0
	assert.Equal(t, expectedValue, promedio)

}
func TestCalcularSalarioCategoriaA(t *testing.T) {
	//Arrange
	horas := float64(176)
	categoria := "A"
	//Act
	salarioFinal := calcularSalario(horas, categoria)
	//Assert
	fmt.Println(salarioFinal, "Salario Final")
	expectedValue := float64(792000)
	assert.Equal(t, expectedValue, salarioFinal)
}
func TestCalcularSalarioCategoriaB(t *testing.T) {
	//Arrange
	horas := float64(176)
	categoria := "B"
	//Act
	salarioFinal := calcularSalario(horas, categoria)
	//Assert
	fmt.Println(salarioFinal, "Salario Final")
	expectedValue := float64(316800)
	assert.Equal(t, expectedValue, salarioFinal)
}
func TestCalcularSalarioCategoriaC(t *testing.T) {
	//Arrange
	horas := float64(176)
	categoria := "C"
	//Act
	salarioFinal := calcularSalario(horas, categoria)
	//Assert
	fmt.Println(salarioFinal, "Salario Final")
	expectedValue := float64(176000)
	assert.Equal(t, expectedValue, salarioFinal)
}
