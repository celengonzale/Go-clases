package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "Error: El salario ingresado no alcanza el mínimo imponible"
}

func Salary(salary int) string {
	if salary < 150000 {
		err := MyError{}
		return err.Error()
	} else {
		return "Debe pagar impuesto"
	}
}

type MyErrorDos struct{}

var errSalaryDos = errors.New("Error: el salario es menor a 10.000")

func (e *MyErrorDos) Error() error {
	return errSalaryDos
}
func SalaryDos(salary int) error {
	if salary <= 10000 {
		err := MyErrorDos{}
		fmt.Println(err.Error())
		return err.Error()
	} else {
		return nil
	}
}

type MyErrorTres struct{}

// func (e *MyErrorTres) Error() error {
//
// }
func main() {
	// Ejercicio 1
	var salary int
	salary = 10000
	resultSalary := Salary(salary)
	fmt.Println(resultSalary)

	//Ejercicio 2
	resultSalaryDos := SalaryDos(salary)
	coincidence := errors.Is(resultSalaryDos, errSalaryDos)
	fmt.Println(coincidence)

	//Ejercicio 3

	//Ejercicio 4

}
