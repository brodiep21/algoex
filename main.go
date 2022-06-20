package main

import "fmt"

func main() {

	a := []int{3, 4}
	fmt.Println(adder(a))
	fmt.Println(sum([]int{7, 10, 14, 45, 76, 89, 12}))
	fmt.Println(pair(a))
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
	arr := [][]int{}
	for i := 0; i < len(a)*len(a); i++ {
		arr = append(arr, []int{})
	}
	count := 0
	for i, v := range a {
		arr[i] = append(arr[i], v)
		for _, v2 := range a {
			if count >= len(a)*len(a) {
				break
			}
			arr[count] = append(arr[count], v2)
			count++
		}
	}
	return arr
}
