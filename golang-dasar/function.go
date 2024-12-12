package main

type operasi func(a int, b int) int

// Format Fungsi
// func NAMA (argument type) type_return
func tambah(x, y int) (result int) {
	return x + y
}

func main() {
	// Lambda
	println(func() string {
		return "lambda"
	}())

	// Closure
	var GetClosure = func() string {
		return "closure"
	}
	println(GetClosure())

	// Callback dengan lambda :: double square
	println(square(func(i int) int {
		return i * i
	}, 2))

	// Callback dengan closure, dengan tipe data implisit
	var Jumlah = func(a int, b int) int {
		return a + b
	}

	// Callback dengan closure, dengan tipe data explisit
	var Kurang operasi = func(a int, b int) int {
		return a - b
	}

	println("Operasi Jumlah : ", Hitung(Jumlah, 5, 3))
	println("Operasi Kurang : ", Hitung(Kurang, 5, 3))
}

func square(f func(int) int, x int) int {
	return f(x * x)
}

func Hitung(o operasi, x int, y int) int {
	return o(x, y)
}
