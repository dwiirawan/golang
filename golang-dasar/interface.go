package main

import "fmt"

// interface
// - interface adalah kumpulan yang berisi method yang abstract
// - type lain akan mengimplementasikan method dalam interface
// - tidak ada perintah implement, suatu interface akan dipenuhi secara implisit begitu ada yang mengimplementasikan

type iface interface {
	method()
}

type myStr string

func (m *myStr) method() {
	println(*m)
}

func main() {
	var i iface
	str := myStr("Hello")
	i = &str
	i.method()

	describe(i)
}

func describe(i iface) {
	fmt.Printf("(%v, %T)\n", i, i)
}
