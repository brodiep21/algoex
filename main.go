package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{3, 4}
	fmt.Println(adder(a))
	fmt.Println(sum([]int{7, 10, 14, 45, 76, 89, 12}))
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
	for i := 0; i < len(a)*len(a); i++ {
		arr = append(arr, []int{})
	}
	count := len(a) - 1
	for i, v := range a {
		arr[i] = append(arr[i], v)
		arr[i] = append(arr[i], v)
		count++
		for _, v2 := range a {
			arr[count] = append(arr[count], v2)
			arr[count] = append(arr[count], v2)
			if count >= len(a)*len(a) {
				break
			}
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
