package sorting

import (
	"testing"
)

func TestCocktailShakerSort(t *testing.T) {

	sort := CocktailShakerSort{}
	SortTester(sort, t)
}
