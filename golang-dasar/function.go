package main

import "fmt"

func kurang(x, y int) (result int) {
	result = x - y
	return result
}

func main() {
	fmt.Println(kurang(10, 5))
}
