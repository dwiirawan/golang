package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	passValue()
	passReference()
	operator()
}

func passValue() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1) //{Subang Jawa Barat Indonesia}
	fmt.Println(address2) //{Bandung Jawa Barat Indonesia}
}

func passReference() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// use pointer
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1) //{Bandung Jawa Barat Indonesia}
	fmt.Println(address2) //&{Bandung Jawa Barat Indonesia}
}

func operator() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// use pointer
	address2 := &address1

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) //{Jakarta DKI Jakarta Indonesia}
	fmt.Println(address2) //&{Jakarta DKI Jakarta Indonesia}
}
