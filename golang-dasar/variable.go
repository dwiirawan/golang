package main

func main() {
	// variabel i akan tetap hidup walaupun looping for sudah selesai
	i := 0
	for i < 5 {
		println(i)
		i++
	}
	println("==========")

	// variabel i hanya hidup dalam blok for
	for i := 0; i < 5; i++ {
		println(i)
	}
	println("==========")

	myMap := map[string]string{"Satu": "Ahad", "Dua": "Senin", "Tiga": "Selasa"}
	// variabel value dan ok tetap hidup walaupun blok if / if else sudah berakhir
	value, ok := myMap["Satu"]
	if ok {
		println(value)
	}
	println("==========")
	// variabel value dan ok hanya hidup dalam blok if / if else
	if value, ok := myMap["Dua"]; ok {
		println(value)
	}
}
