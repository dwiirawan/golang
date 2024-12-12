package main

// method
// - kita bisa mendefinisikan suatu method pada sebuah type
// - method adalah fungsi yang mempunyai argumen khusus receiver berupa type

type MyStr string

func (m MyStr) Salam() {
	m = "Selamat Pagi"
	println(m)
}

// receiver bisa berupa pointer
func (m *MyStr) Change() {
	*m = MyStr("Selamat Sore")
}

func (m *MyStr) Print() {
	println(*m)
}

func main() {
	var str MyStr
	str.Salam() // Selamat Pagi

	str = MyStr("Selamat Siang")
	str.Print() // Selamat Siang

	str.Change()
	str.Print() // Selamat Sore
}
