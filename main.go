package main

import "fmt"

func main() {

	a := []int{3, 4, 5, 6, 7, 8, 9}
	fmt.Println(adder(a))
	fmt.Println(sum(a))
}

func adder(a []int) int {
	return 1 + a[0]
}

func sum(a []int) int {
	var total int
	for _, v := range a {
		total += v
	}
	return total
}
