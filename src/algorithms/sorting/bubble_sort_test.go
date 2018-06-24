package sorting

import (
	"reflect"
	"testing"
)

func TestBubblesort(t *testing.T) {

	input := []int{1, 2, 3, 2, 1}
	sorted_input := []int{1, 1, 2, 2, 3}

	if !reflect.DeepEqual(Bubblesort(input), sorted_input) {
		t.Error("Not Sorted")
	}

	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sorted_input = []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	if !reflect.DeepEqual(Bubblesort(input), sorted_input) {
		t.Error("Not Sorted")
	}

	input = []int{1}
	sorted_input = []int{1}

	if !reflect.DeepEqual(Bubblesort(input), sorted_input) {
		t.Error("Not Sorted")
	}

	input = []int{}
	sorted_input = []int{}

	if !reflect.DeepEqual(Bubblesort(input), sorted_input) {
		t.Error("Not Sorted")
	}

}
