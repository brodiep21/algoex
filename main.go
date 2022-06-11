package main

import "fmt"

func main() {

	a := []int{3, 4, 5, 6, 7, 8, 9}
	fmt.Println(adder(a))
	fmt.Println(sum([]int{7, 10, 14, 45, 76, 89, 12}))
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

func pair(a []int) [][]int {

}
