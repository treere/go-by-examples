package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got '%d' expected '%d' give, %v", want, got, numbers)
	}

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int ) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{3, 4})
		want := []int{3, 4}

		checkSum(t, got, want)
	})

	t.Run("empty slices", func (t *testing.T) {
		got := SumAllTails([]int{}, []int{1,2,3})
		want := []int{0,5}

		checkSum(t, got, want)
	})
}
