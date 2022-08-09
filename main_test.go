package main

import (
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
	arr := []int{4, 5, 6, 7, 8, 9, 10}

	have := adder(arr)
	want := 5

	if want != have {
		t.Errorf("Have %v, want %v", have, want)
	}

}

func TestSum(t *testing.T) {
	t.Run("test slice with only 1 element", func(t *testing.T) {
		arr := []int{1}

		have := sum(arr)
		want := 1

		if want != have {
			t.Errorf("Have %v, want %v", have, want)
		}
	})
	t.Run("test Large slice with multiple elements", func(t *testing.T) {
		arr := []int{7, 10, 14, 45, 90, 89, 12}
		have := sum(arr)
		want := 267

		if want != have {
			t.Errorf("Have %v, want %v", have, want)
		}
	})
}

func TestPair(t *testing.T) {
	t.Skip()
	arr := []int{1, 2}
	t.Run("test the first pair in the first slice", func(t *testing.T) {
		firstpair := pair(arr)

		got := firstpair[0]
		want := [][]int{{1, 1}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Have %d, want %d", got, want)
		}
	})
	t.Skip()
	t.Run("test all of the slices match up", func(t *testing.T) {
		got := pair(arr)
		want := [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}

		if !reflect.DeepEqual(got, got) {
			t.Errorf("Have %d, want %d", got, want)
		}

	})

}

func TestBubSort(t *testing.T) {
	arr := []int{3, 7, 12, 4, 6, 8}

	got := bubsort(arr)
	want := []int{3, 4, 6, 7, 8, 12}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Have %d, want %d", got, want)
	}
}

func TestTournyWin(t *testing.T) {
	t.Run("test large set of results", func(t *testing.T) {
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
		got := TournamentWinner(competitions, results)
		want := "C#"
		if got != want {
			t.Errorf("Have %s, want %s", got, want)
		}
	})
}

func TestConstructChange(t *testing.T) {
	t.Run("test 1 number", func(t *testing.T) {
		arr := []int{1}

		got := NonConstructibleChange(arr)
		want := 2

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("test 2-number ray", func(t *testing.T) {
		arr := []int{3, 2}

		got := NonConstructibleChange(arr)
		want := 1

		if got != want {
			t.Errorf("Have %d, want %d", got, want)
		}
	})
	t.Run("test large numberset", func(t *testing.T) {
		arr := []int{5, 6, 1, 1, 2, 3, 4, 9}

		got := NonConstructibleChange(arr)
		want := 32

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}

func TestMinWaitingTime(t *testing.T) {
	t.Run("test array with 1 element", func(t *testing.T) {
		arr := []int{4}

		got := MinimumWaitingTime(arr)
		want := 0

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("test a large array", func(t *testing.T) {
		arr := []int{1, 5, 2, 4, 3}

		got := MinimumWaitingTime(arr)
		want := 20

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("test a 2 element array", func(t *testing.T) {
		arr := []int{3, 1}

		got := MinimumWaitingTime(arr)
		want := 1

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}

func TestClassPhotos(t *testing.T) {
	t.Run("test passing case for redshirts", func(t *testing.T) {
		redshirt := []int{6, 9, 2, 4, 5}
		blueshirt := []int{5, 8, 1, 3, 4}

		got := ClassPhotos(redshirt, blueshirt)
		want := true

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("test passing case for blueshirts", func(t *testing.T) {
		redshirt := []int{5, 8, 1, 3, 4}
		blueshirt := []int{6, 9, 2, 4, 5}

		got := ClassPhotos(redshirt, blueshirt)
		want := true

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("test an array that MUST be in ascending order to pass", func(t *testing.T) {
		redshirt := []int{5, 6}
		blueshirt := []int{5, 4}

		got := ClassPhotos(redshirt, blueshirt)
		want := true

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}

func TestTandemBicycle(t *testing.T) {
	t.Run("test Large array for false (slowest)", func(t *testing.T) {
		redshirts := []int{1, 2, 1, 9, 12, 3, 1, 54, 21, 231, 32, 1}
		blueshirts := []int{3, 3, 4, 6, 1, 2, 5, 6, 34, 256, 123, 32}
		fastest := false

		want := 484
		got := TandemBicycle(redshirts, blueshirts, fastest)

		if want != got {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	})
	t.Run("test Large array for true (fastest)", func(t *testing.T) {
		redshirts := []int{1, 2, 1, 9, 12, 3, 1, 54, 21, 231, 32, 1}
		blueshirts := []int{3, 3, 4, 6, 1, 2, 5, 6, 34, 256, 123, 32}
		fastest := true

		want := 816
		got := TandemBicycle(redshirts, blueshirts, fastest)

		if want != got {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	})
	t.Run("test small array", func(t *testing.T) {
		redshirts := []int{6}
		blueshirts := []int{1}
		fastest := true

		want := 6
		got := TandemBicycle(redshirts, blueshirts, fastest)

		if want != got {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	})
	t.Run("test zero element arrays for true", func(t *testing.T) {
		redshirts := []int{}
		blueshirts := []int{}
		fastest := true

		want := 0
		got := TandemBicycle(redshirts, blueshirts, fastest)

		if want != got {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	})
	t.Run("test zero element arrays for false", func(t *testing.T) {
		redshirts := []int{}
		blueshirts := []int{}
		fastest := false

		want := 0
		got := TandemBicycle(redshirts, blueshirts, fastest)

		if want != got {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	})
}

func TestGetNthfib(t *testing.T) {
	t.Run("Test a high number in the sequence", func(t *testing.T) {
		sequenceNum := 18

		want := 1597
		got := GetNthFib(sequenceNum)

		if got != want {
			t.Errorf("Want %d, got %d", want, got)
		}
	})
	t.Run("Test a low number in the sequence", func(t *testing.T) {
		sequenceNum := 6

		want := 5
		got := GetNthFib(sequenceNum)

		if got != want {
			t.Errorf("Want %d, got %d", want, got)
		}
	})
	t.Run("Test a medium number in the sequence", func(t *testing.T) {
		sequenceNum := 9

		want := 21
		got := GetNthFib(sequenceNum)

		if got != want {
			t.Errorf("Want %d, got %d", want, got)
		}
	})
}

func TestInsertionSort(t *testing.T) {
	arr := []int{4, 3, 7, 10, 8, 12, 2, 1}

	want := []int{1, 2, 3, 4, 7, 8, 10, 12}
	got := InsertionSort(arr)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	t.Run("Test a valid palindrome", func(t *testing.T) {
		str := "radar"

		want := true
		got := IsPalindrome(str)

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Test an invalid Palindrome", func(t *testing.T) {
		str := "Arnold"

		want := false
		got := IsPalindrome(str)

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Test a one letter word", func(t *testing.T) {
		str := "a"

		want := true
		got := IsPalindrome(str)

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})

}

func TestCaesarCipher(t *testing.T) {
	t.Run("testing a large change in keys which wrap the alphabet", func(t *testing.T) {
		str := "iwufqnkqkqoolxzzlzihqfm"
		key := 25

		want := "hvtepmjpjpnnkwyykyhgpel"
		got := CaesarCipherEncryptor(str, key)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("testing a small string with low key change", func(t *testing.T) {
		str := "abc"
		key := 2

		want := "cde"
		got := CaesarCipherEncryptor(str, key)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("testing a high key value that wraps multiple times", func(t *testing.T) {
		str := "abc"
		key := 57

		want := "fgh"
		got := CaesarCipherEncryptor(str, key)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
}

func TestRunLengthEncoding(t *testing.T) {
	t.Run("Test lowercase and uppercase characters", func(t *testing.T) {
		str := "aA"

		want := "1a1A"
		got := RunLengthEncoding(str)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test a long string with all Uppercase characters", func(t *testing.T) {

		str := "AAAAAAAAAAAAABBCCCCDD"

		want := "9A4A2B4C2D"
		got := RunLengthEncoding(str)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test a long string with unique characters", func(t *testing.T) {

		str := "************^^^^^^^$$$$$$%%%%%%%!!!!!!AAAAAAAAAAAAAAAAAAAA"

		want := "9*3*7^6$7%6!9A9A2A"
		got := RunLengthEncoding(str)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test a long string with unique characters", func(t *testing.T) {

		str := "........______=========AAAA   AAABBBB   BBB"

		want := "8.6_9=4A3 3A4B3 3B"
		got := RunLengthEncoding(str)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})

}
func TestGenerateDocument(t *testing.T) {
	t.Run("test true sentence with special characters", func(t *testing.T) {
		str := "Bste!hetsi ogEAxpelrt x "
		document := "AlgoExpert is the Best!"

		want := true
		got := GenerateDocument(str, document)

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("test empty string", func(t *testing.T) {
		str := "Basdfasdf"
		document := ""

		want := true
		got := GenerateDocument(str, document)

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("test against upper/lowercase letters", func(t *testing.T) {
		str := "helloworld "
		document := "hello wOrld"

		want := false
		got := GenerateDocument(str, document)

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}

func TestFirstNonRepeatingCharacter(t *testing.T) {
	t.Run("Test a long string with the first char late in the string", func(t *testing.T) {
		str := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy"

		want := 25
		got := FirstNonRepeatingCharacter(str)

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("Test an empty string", func(t *testing.T) {
		str := ""

		want := -1
		got := FirstNonRepeatingCharacter(str)

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("Test a long string with all repeating characters", func(t *testing.T) {
		str := "aaaaaaaaaaaaaaaaaaaabbbbbbbbbbcccccccccccdddddddddddeeeeeeeeffghgh"

		want := -1
		got := FirstNonRepeatingCharacter(str)

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("Test a medium sized string", func(t *testing.T) {
		str := "ggyllaylacrhdzedddjsc"

		want := 10
		got := FirstNonRepeatingCharacter(str)

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}

func TestProductSum(t *testing.T) {
	t.Run("test a small interface with multiple interior arrays", func(t *testing.T) {
		input := SpecialArray{
			5, 2,
			SpecialArray{7, -1},
			3,
			SpecialArray{
				6,
				SpecialArray{
					-13, 8,
				},
				4,
			},
		}

		want := 12
		got := ProductSum(input)
		if want != got {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("test a small array ", func(t *testing.T) {
		input := SpecialArray{1, 2, 3, 4, 5}

		want := 15
		got := ProductSum(input)
		if want != got {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}

func TestThreeNumberSum(t *testing.T) {
	t.Run("Test functional return with attainable target", func(t *testing.T) {
		arr := []int{12, 3, 1, 2, -6, 5, -8, 6}
		target := 0

		want := [][]int{
			[]int{-8, 2, 6},
			[]int{-8, 3, 5},
			[]int{-6, 1, 5},
		}
		got := ThreeNumberSum(arr, target)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v ", want, got)
		}
	})
	t.Run("Test return with no attainable target", func(t *testing.T) {
		arr := []int{60, 12, 5, -12, -10, 4, 6}
		target := 2

		want := [][]int{}
		got := ThreeNumberSum(arr, target)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})

	t.Run("Test a large array with more than 4 conditions", func(t *testing.T) {
		arr := []int{12, 3, 1, 2, -6, 5, 0, -8, -1, 6, -5}
		target := 0

		got := ThreeNumberSum(arr, target)
		want := [][]int{
			[]int{-8, 2, 6},
			[]int{-8, 3, 5},
			[]int{-6, 0, 6},
			[]int{-6, 1, 5},
			[]int{-5, -1, 6},
			[]int{-5, 0, 5},
			[]int{-5, 2, 3},
			[]int{-1, 0, 1},
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
}

func TestSmallestDifference(t *testing.T) {
	t.Run("Array1 of small numbers - array2 with large numbers", func(t *testing.T) {
		arr1 := []int{10, 0, 20, 25, 2200}
		arr2 := []int{1005, 1006, 1014, 1032, 1031}

		want := []int{25, 1005}
		got := SmallestDifference(arr1, arr2)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
	t.Run("Test a small set of numbers in both arrays", func(t *testing.T) {
		arr1 := []int{3, 20}
		arr2 := []int{21, 5}

		want := []int{20, 21}
		got := SmallestDifference(arr1, arr2)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
}

func TestMoveElementToEnd(t *testing.T) {
	t.Run("Empty array return", func(t *testing.T) {
		arr := []int{}
		toMove := 3

		want := []int{}
		got := MoveElementToEnd(arr, toMove)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
	t.Run("An array without a moveable number", func(t *testing.T) {
		arr := []int{1, 2, 4, 5, 6}
		toMove := 3

		want := []int{1, 2, 4, 5, 6}
		got := MoveElementToEnd(arr, toMove)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
	t.Run("Large array with many numbers", func(t *testing.T) {
		arr := []int{5, 5, 5, 5, 5, 5, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12}
		toMove := 5

		want := []int{12, 11, 10, 9, 8, 7, 1, 2, 3, 4, 6, 5, 5, 5, 5, 5, 5}
		got := MoveElementToEnd(arr, toMove)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
	t.Run("Outlier for pointer values", func(t *testing.T) {
		arr := []int{2, 4, 2, 5, 6, 2, 2}
		toMove := 2

		want := []int{6, 4, 5, 2, 2, 2, 2}
		got := MoveElementToEnd(arr, toMove)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
}

func TestIsMonotonic(t *testing.T) {
	t.Run("Correct monotonic array with negative numbers", func(t *testing.T) {
		arr := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}

		want := true
		got := IsMonotonic(arr)

		if want != got {
			t.Errorf("Wanted %t, got %t", want, got)
		}
	})
	t.Run("Empty array", func(t *testing.T) {
		arr := []int{}

		want := true
		got := IsMonotonic(arr)

		if want != got {
			t.Errorf("Wanted %t, got %t", want, got)
		}
	})
	t.Run("Array with a length of one", func(t *testing.T) {
		arr := []int{1}

		want := true
		got := IsMonotonic(arr)

		if want != got {
			t.Errorf("Wanted %t, got %t", want, got)
		}
	})
	t.Run("Repeating array of the same number", func(t *testing.T) {
		arr := []int{1, 1, 1, 1, 1, 1}

		want := true
		got := IsMonotonic(arr)

		if want != got {
			t.Errorf("Wanted %t, got %t", want, got)
		}
	})
	t.Run("Incorrect monotonic array with negative numbers", func(t *testing.T) {
		arr := []int{-1, -5, -10, -1100, -900, -1101, -1102, -9001}

		want := false
		got := IsMonotonic(arr)

		if want != got {
			t.Errorf("Wanted %t, got %t", want, got)
		}
	})
}
