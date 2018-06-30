package sorting

import "testing"

func TestSelectionSort(t *testing.T) {

    t.Parallel()
	sort := SelectionSort{}
	SortTester(sort, t)

}
