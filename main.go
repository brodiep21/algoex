package main

import "fmt"

func main() {

	a := []int{3, 4}
	fmt.Println(adder(a))
	fmt.Println(sum([]int{7, 10, 14, 45, 76, 89, 12}))
	fmt.Println(pair(a))
	b := []int{3, 7, 12, 4, 6, 8}
	fmt.Println(bubsort(b))
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
    var bestTeam string
    teams[bestTeam] = 0 
    
    for k, v := range competitions {
        result := results[k]

        if result == 0 {
            name, ok := teams[v[result+1]]; ok {
                teams[name] += 3
                } 
            teams[v[result+1]] = 3
            } else {
             name, ok := teams[v[result-1]]; ok {
                teams[name] += 3
            }
        }
    }
	return ""
}