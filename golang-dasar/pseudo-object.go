package main

import "fmt"

type becak struct {
	roda  int
	warna string
}

func (b *becak) caraJalan() string {
	return "dikayuh"
}

func main() {
	becak1 := becak{roda: 3, warna: "biru"}
	fmt.Printf("%v,%T\n", becak1, becak1)
	println("cara jalan becak1:", becak1.caraJalan())
	println("--------------------------")

	becak2 := &becak1
	fmt.Printf("%v,%T\n", becak2, becak2)
	println("cara jalan becak2:", becak2.caraJalan())
	println("--------------------------")

	becak3 := new(becak)
	becak3.roda = 3
	becak3.warna = "merah"
	fmt.Printf("%v, %T\n", becak3, becak3)
	println("cara jalan becak3:", becak3.caraJalan())
}
