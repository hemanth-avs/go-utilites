package sorting

import (
	"math/rand"
	"reflect"
	"testing"
)

func generateRandArray(n int) []int {
	randArray := make([]int, n)
	for i := 0; i < n; i++ {
		randArray[i] = rand.Int()
	}
	return randArray
}

func SortTester(sort SortAlgorithm, t *testing.T) {

	checkSortResults := func(t *testing.T, sort SortAlgorithm, testInput []int, expectedOutput []int) {

		if !reflect.DeepEqual(sort.Sort(testInput), expectedOutput) {
			t.Error("Sort Failed")
		}
	}

	randArray100 := generateRandArray(100)
	randArray500 := generateRandArray(500)
	randArray1000 := generateRandArray(1000)

	testCases := []struct {
		name     string
		given    []int
		expected []int
	}{
		{"testcase1", []int{1, 2, 3, 2, 1}, []int{1, 1, 2, 2, 3}},
		{"testcase2", []int{6, 1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6}},
		{"testcase3", []int{1, 2, 3, 4, 5, 6, 0}, []int{0, 1, 2, 3, 4, 5, 6}},
		{"testcase4", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}},
		{"testcase5", []int{1}, []int{1}},
		{"testcase6", []int{}, []int{}},
		{"testcase7", randArray100, sort.Sort(randArray100)},
		{"testcase8", randArray500, sort.Sort(randArray500)},
		{"testcase9", randArray1000, sort.Sort(randArray1000)},
	}

	for _, tc := range testCases {
		checkSortResults(t, sort, tc.given, tc.expected)
	}

}
