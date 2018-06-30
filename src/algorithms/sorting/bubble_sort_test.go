package sorting

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

    t.Parallel()
	sort := BubbleSort{}
	SortTester(sort, t)
}
