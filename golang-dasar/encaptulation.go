package main

import "hello/latihan"

func main() {
	becak := latihan.NewBecak()
	becak.SetWarna("Biru")
	println(becak.CaraJalan())
	println("Jumlah roda:", becak.GetRoda())
	println("warna:", becak.GetWarna())
}
