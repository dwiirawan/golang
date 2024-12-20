package main

import "fmt"

type Hewan struct {
	Nama  string
	Nyata bool
}

func (c *Hewan) Cetak() {
	fmt.Printf("Nama: '%s', Nyata: %t\n", c.Nama, c.Nyata)
}

type HewanTerbang struct {
	Hewan
	PanjangSayap int
}

func (c HewanTerbang) Cetak() {
	fmt.Printf("Nama: '%s', Nyata: %t, PanjangSayap: %d\n", c.Nama, c.Nyata, c.PanjangSayap)
}

type Unicorn struct {
	Hewan
}

type Naga struct {
	HewanTerbang
}

type Pterodactilus struct {
	HewanTerbang
}

func NewPterodactyl(panjangSayap int) *Pterodactilus {
	p := new(Pterodactilus)
	p.Nama = "Pterodactilus"
	p.Nyata = true
	p.PanjangSayap = panjangSayap

	return p
}

func main() {
	hewan := new(Hewan)
	hewan.Nama = "Sembarang hewan"
	hewan.Nyata = false

	naga := new(Naga)
	naga.Nama = "Naga"
	hewan.Nyata = false

	uni := new(Unicorn)
	uni.Nama = "Unicorn"
	uni.Nyata = false

	p1 := new(Pterodactilus)
	p1.Nama = "Pterodactilus"
	p1.Nyata = true
	p1.PanjangSayap = 5

	p2 := NewPterodactyl(8)

	hewan.Cetak()
	naga.Cetak()
	uni.Cetak()
	p1.Cetak()
	p2.Cetak()

	animals := []*Hewan{
		hewan,
		&naga.Hewan,
		&uni.Hewan,
		&p1.Hewan,
		&p2.Hewan,
	}
	fmt.Println("Cetak() melalui embedded type Hewan")
	for _, c := range animals {
		c.Cetak()
	}
}
