package sorting

type InsertionSort struct {
}

func (i InsertionSort) Sort(input []int) []int {

	for ol := 1; ol < len(input); ol++ {
		indexElement := input[ol]
		innerLoop := ol
		for innerLoop > 0 && input[innerLoop-1] > indexElement {
			input[innerLoop] = input[innerLoop-1]
			innerLoop--
		}
		input[innerLoop] = indexElement

	}
	return input
}
