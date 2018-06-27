package sorting

import "testing"

func TestSelectionSort(t *testing.T) {

	sort := SelectionSort{}
	SortTester(sort, t)

}
