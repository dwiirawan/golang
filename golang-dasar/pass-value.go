package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	operator()
}

func passValue() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}

func passReference() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// use pointer
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}

func operator() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// use pointer
	address2 := &address1

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
