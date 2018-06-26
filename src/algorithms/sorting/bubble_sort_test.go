package sorting

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	sort := BubbleSort{}
	SortTester(sort, t)
}
