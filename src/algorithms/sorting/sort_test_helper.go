package sorting

import (
	"math/rand"
	"reflect"
	gosort "sort"
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
	randArray100Sorted := make([]int, len(randArray100))
	copy(randArray100Sorted, randArray100)
	gosort.Ints(randArray100Sorted)

	randArray150000 := generateRandArray(150000)
	randArray150000Sorted := make([]int, len(randArray150000))
	copy(randArray150000Sorted, randArray150000)
	gosort.Ints(randArray150000Sorted)

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
		{"testcase7", randArray100, randArray100Sorted},
		{"testcase8", randArray150000, randArray150000Sorted},
	}

	for _, tc := range testCases {
		checkSortResults(t, sort, tc.given, tc.expected)
	}

}
