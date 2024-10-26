package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("Student 1 :", s1.name)
	fmt.Println("Student 2 :", s2.name)
	fmt.Println("Student 3 :", s3.name)
}
