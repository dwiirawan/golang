package main

func main() {
	i := 10
	p := &i     // menunjuk ke i
	println(p)  // alamat memory p
	println(*p) // baca i lewat pointer

	*p = 20    // set i lewat pointer
	println(i) // lihat nilai terbaru dari i
}
