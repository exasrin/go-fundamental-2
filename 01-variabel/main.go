package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string = "Rasyid Sulaiman"
	firstName = "Alfaathir"

	fmt.Println("Hello My Name Is " + firstName + " " + lastName)

	var age = 23
	address := "Jakarta"
	fmt.Printf("Age: %d, address: %s\n", age, address)

	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	var first, second, third string = "satu", "dua", "tiga"
	fmt.Println(first + " " + second + " " + third)

	one, two, three := 1, "Dua", 3.0
	fmt.Printf("%d, %s, %.2f \n", one, two, three)

	_ = "Alfaathir Rasyid Sulaiaman"
	_ = 2000
	namaPertama, _ := "Asrin", "Tidak Ada"
	fmt.Println(namaPertama)

	newName := new(string)
	fmt.Println(newName)  // by value
	fmt.Println(*newName) // by reference
}
