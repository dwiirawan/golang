package main

import "fmt"

// struct
// - sebuah tipe data abstract
// - berisi dari kumpulan dari berbagai type
// - struct bisa digunakan dalam konsep class

type User struct {
	ID   uint32
	Name string
}

func main() {
	var user User
	user.ID = 1
	user.Name = "Jacky"
	fmt.Printf("%v\n", user)
	println(user.Name)

	user2 := User{ID: 2, Name: "JetLee"}
	fmt.Printf("%v\n", user2)
	println(user2.Name)
}
