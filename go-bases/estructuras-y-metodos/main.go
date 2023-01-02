package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{
		ID:          1,
		Name:        "Galletas",
		Price:       100.0,
		Description: "Galletas de chocolate",
		Category:    "Dulce",
	},
	{
		ID:          2,
		Name:        "Arroz",
		Price:       50.0,
		Description: "Arroz integral para personas ",
		Category:    "Salado",
	},
}

func (p *Product) Save(prod Product) Product {
	Products = append(Products, prod)
	return prod

}
func (p *Product) GetAll() {
	fmt.Println(Products)
}
func getById(id int) Product {
	var findProduct Product
	for _, product := range Products {
		if id == product.ID {
			findProduct = product
		}
	}
	return findProduct
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}
type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(e.Name, e.Position, e.DateOfBirth, e.ID)

}
func main() {
	structProduct := Product{}
	structProduct.GetAll()
	newProd := Product{
		3,
		"Manzana",
		10,
		"*************",
		"Fitness",
	}
	saved := structProduct.Save(newProd)
	prod := getById(3)
	fmt.Println(prod, "ID Success")
	fmt.Println(saved, "Producto Añadido")
	structProduct.GetAll()
	pers := Person{
		1,
		"Cele",
		"22/01/1998",
	}
	empl := Employee{
		1,
		"Developer",
		pers,
	}
	fmt.Println(empl)
	empl.PrintEmployee()

}
