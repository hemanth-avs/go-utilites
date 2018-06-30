package sorting

import (
	"math/rand"
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

	checkSortResults := func(t *testing.T, sort SortAlgorithm, testInput []int) {
		sortedArray := sort.Sort(testInput)
		sortedArrayIntSlice := gosort.IntSlice(sortedArray[0:])

		if !gosort.IsSorted(sortedArrayIntSlice) {
			t.Error("Sort Failed")
		}
	}

	randArray100 := generateRandArray(100)
	randArray150000 := generateRandArray(150000)

	testCases := []struct {
		name  string
		given []int
	}{
		{"testcase1", []int{1, 2, 3, 2, 1}},
		{"testcase2", []int{6, 1, 2, 3, 4, 5}},
		{"testcase3", []int{1, 2, 3, 4, 5, 6, 0}},
		{"testcase4", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"testcase5", []int{1}},
		{"testcase6", []int{}},
		{"testcase7", randArray100},
		{"testcase8", randArray150000},
		{"testcase9", randArray300000},
		{"testcase10", []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}},

	}

	for _, tc := range testCases {
		checkSortResults(t, sort, tc.given)
	}

}
