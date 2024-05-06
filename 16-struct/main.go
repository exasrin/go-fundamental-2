package main

import "fmt"

func main() {
	var s1 student
	s1.name = "Alfaathir Rasyid Sulaiaman"
	s1.grade = 88

	var s2 = student{"ethan", 2}
	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
}

type student struct {
	name  string
	grade int
}
