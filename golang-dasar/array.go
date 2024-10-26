package main

import "fmt"

type foo struct {
	Id     int
	Nama   string
	Kelas  string
	Alamat string
}

func main() {
	var bar []foo

	bar = append(bar,
		foo{Id: 1, Nama: "Dwi", Kelas: "VII A", Alamat: "Kebumen"},
		foo{Id: 2, Nama: "Irawan", Kelas: "VII B", Alamat: "Kebumen"})

	// Looping for
	for i := 0; i < len(bar); i++ {
		fmt.Println("Perulangan ke-", bar[i].Nama)
	}
}
