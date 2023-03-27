package main

import (
	"fmt"
	"testing"
)

func errorString(want, got int) string {
	return fmt.Sprintf("There was an error. Wanted %d, got %d\n", want, got)
}

func TestAdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	var want int
	for _, number := range numbers {
		want += number
	}
	got := Add(numbers)

	if want != got {
		t.Error(errorString(want, got))
	}
}

func TestSubs(t *testing.T) {
	want := 20
	got := Substract(65, 45)

	if want != got {
		t.Error(errorString(want, got))
	}
}
