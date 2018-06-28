package sorting

import (
	"reflect"
	"testing"
)

func SortTester(sort SortAlgorithm, t *testing.T) {

	checkSortResults := func(t *testing.T, sort SortAlgorithm, testInput []int, expectedOutput []int) {

		if !reflect.DeepEqual(sort.Sort(testInput), expectedOutput) {
			t.Error("Sort Failed")
		}
	}

	input := []int{1, 2, 3, 2, 1, 0}
	sorted_input := []int{1, 1, 2, 2, 3}
	checkSortResults(t, sort, input, sorted_input)

	input = []int{6, 1, 2, 3, 4, 5}
	sorted_input = []int{1, 2, 3, 4, 5, 6}
	checkSortResults(t, sort, input, sorted_input)

	input = []int{1, 2, 3, 4, 5, 6, 0}
	sorted_input = []int{0, 1, 2, 3, 4, 5, 6}
	checkSortResults(t, sort, input, sorted_input)

	input = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sorted_input = []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	checkSortResults(t, sort, input, sorted_input)

	input = []int{1}
	sorted_input = []int{1}
	checkSortResults(t, sort, input, sorted_input)

	input = []int{}
	sorted_input = []int{}
	checkSortResults(t, sort, input, sorted_input)

}
