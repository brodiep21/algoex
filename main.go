package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 2, 4, 3, 5}
	fmt.Println(pair(a))
	b := []int{3, 7, 12, 4, 6, 8}
	fmt.Println(bubsort(b))

	competitions := [][]string{
		{"HTML", "Java"},
		{"Java", "Python"},
		{"Python", "HTML"},
		{"C#", "Python"},
		{"Java", "C#"},
		{"C#", "HTML"},
		{"SQL", "C#"},
		{"HTML", "SQL"},
		{"SQL", "Python"},
		{"SQL", "Java"},
	}
	results := []int{0, 1, 1, 1, 0, 1, 0, 1, 1, 0}
	fmt.Println(TournamentWinner(competitions, results))
	fmt.Println(MinimumWaitingTime(a))
	fmt.Println(GetNthFib(5))
	fmt.Println(InsertionSort(b))
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

//solve this problem
func pair(a []int) [][]int {
	arr := [][]int{}
	size := len(a) * len(a)
	counter := 1
	for i := 0; i < size; i++ {
		arr = append(arr, []int{})
	}
	for i := 0; i < len(a); i++ {
		arr[counter-1] = append(arr[counter-1], a[i])
		for j := 0; j < len(a); j++ {
			arr[j] = append(arr[j], a[i])
		}
		counter++
		if counter < size {
			arr[counter] = append(arr[counter-1], a[i])
		}
	}
	return arr
}

func bubsort(ar []int) []int {
	//switch elements, one-time pass, full swap
	for j := 0; j < len(ar); j++ {
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] > ar[i+1] {
				change := ar[i]
				ar[i] = ar[i+1]
				ar[i+1] = change
			}
		}
	}
	return ar
}

func revbubsort(ar []int) []int {
	//switch elements, one-time pass, full swap
	for j := 0; j < len(ar); j++ {
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] < ar[i+1] {
				change := ar[i]
				ar[i] = ar[i+1]
				ar[i+1] = change
			}
		}
	}
	return ar
}

func TournamentWinner(competitions [][]string, results []int) string {
	teams := make(map[string]int)

	for k, v := range competitions {
		result := results[k]
		winner := ""

		if result == 0 {
			winner = v[result+1]
		} else {
			winner = v[result-1]
		}

		if score, ok := teams[winner]; ok {
			teams[winner] = score + 3
		} else {
			teams[winner] = 3
		}
	}

	var bestTeam string
	var bestTeamScore int
	for team, score := range teams {
		if score > bestTeamScore {
			bestTeamScore = score
			bestTeam = team
		}
	}
	return bestTeam
}

func NonConstructibleChange(coins []int) int {
	if len(coins) < 1 {
		return 1
	}
	sort.Ints(coins)
	smallestChange := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] > smallestChange+1 {
			break
		}
		smallestChange += coins[i]
	}
	return smallestChange + 1
}

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	time := 0
	totalWaitTime := 0

	for i := 0; i < len(queries)-1; i++ {
		time += queries[i]
		totalWaitTime += time
	}
	return totalWaitTime
}

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	classPhotoWorks := false
	if redShirtHeights[0] > blueShirtHeights[0] {
		classPhotoWorks = checkBackRow(redShirtHeights, blueShirtHeights)
		return classPhotoWorks
	}
	classPhotoWorks = checkBackRow(blueShirtHeights, redShirtHeights)

	return classPhotoWorks
}

func checkBackRow(backRow []int, frontRow []int) bool {
	photoWorks := false
	result := 0

	for k, v := range backRow {
		result = frontRow[k]

		if v > result {
			photoWorks = true
			continue
		}
		photoWorks = false
		break
	}
	return photoWorks
}

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {

	if fastest {
		return fastestSort(redShirtSpeeds, blueShirtSpeeds)
	}
	return slowestSort(redShirtSpeeds, blueShirtSpeeds)

}

func fastestSort(arr1 []int, arr2 []int) int {
	bubsort(arr1)
	revbubsort(arr2)

	totalSpeed := 0
	for k, v := range arr1 {
		comparison := arr2[k]
		if v > comparison {
			totalSpeed += v
		} else {
			totalSpeed += comparison
		}
	}
	return totalSpeed
}

func slowestSort(arr1 []int, arr2 []int) int {
	bubsort(arr1)
	bubsort(arr2)

	totalSpeed := 0
	for k, v := range arr1 {
		comparison := arr2[k]
		if v > comparison {
			totalSpeed += v
		} else {
			totalSpeed += comparison
		}
	}
	return totalSpeed
}

func GetNthFib(n int) int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	fib[2] = 1

	if n <= 2 {
		return fib[n-1]
	}
	for i := 3; i < n; i++ {
		fib[i] = (fib[i-1] + fib[i-2])
	}
	return fib[n-1]
}

// type SpecialArray []interface{}

// // Tip: Each element of `array` can either be cast
// // to `SpecialArray` or `int`.
// func ProductSum(array []interface{}) int {

// 	return -1
// }

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		temp := array[i]
		position := i - 1
		for position >= 0 {
			if array[position] > temp {
				array[position+1] = array[position]
				position -= 1
			} else {
				break
			}

		}
		array[position+1] = temp
	}
	return array
}
