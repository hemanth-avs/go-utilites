package sorting

import (
	"testing"
)

func TestCocktailShakerSort(t *testing.T) {

    t.Parallel()
	sort := CocktailShakerSort{}
	SortTester(sort, t)
}
