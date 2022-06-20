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

		have := firstpair[0]
		want := [][]int{{1, 1}}

		if reflect.DeepEqual(have, want) {
			t.Errorf("Have %d, want %d", have, want)
		}
	})
	t.Run("test all of the slices match up", func(t *testing.T) {
		have := pair(arr)
		want := [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}

		if reflect.DeepEqual(want, have) {
			t.Errorf("Have %d, want %d", have, want)
		}

	})

}
