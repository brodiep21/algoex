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
	arr := []int{1, 2}
	t.Run("test the first pair in the first slice", func(t *testing.T) {
		firstpair := pair(arr)

		got := firstpair[0]
		want := [][]int{{1, 1}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Have %d, want %d", got, want)
		}
	})
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
	t.Run("test an array that MUST be in ascending order to pass", func(t *testing.t) {
		redshirt := []int{5, 6}
		blueshirt := []int{5, 4}

		got := ClassPhotos(redshirt, blueshirt)
		want := true

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}
