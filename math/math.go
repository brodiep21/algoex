package math

func Adder(a []int) int {
	return 1 + a[0]
}

func Sum(a []int) int {
	var total int
	for _, v := range a {
		total += v
	}
	return total
}

func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	return (n - 1) + (n - 2)
}
