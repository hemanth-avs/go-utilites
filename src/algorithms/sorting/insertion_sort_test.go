package sorting

import "testing"

func TestInsertionSort(t *testing.T) {

    t.Parallel()
	sort := InsertionSort{}
	SortTester(sort, t)

}
