package main

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	println(total)
}

func main() {
	sum(1)
	sum(1, 2)
	sum(2, 3, 4)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
