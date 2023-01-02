package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("customers.txt")
	defer func() {
		fmt.Println("Ejecución finalizada**")
	}()
	if err != nil {
		panic("ERROR: El archivo indicado no fue encontrado o está dañado.")
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
