package sorting

type InsertionSort struct {
}

func (i InsertionSort) Sort(input []int) []int {
	for ol := 1; ol < len(input); ol++ {
		for il := ol; il > 0; il-- {
			if input[il] < input[il-1] {
				input[il], input[il-1] = input[il-1], input[il]
			}
		}
	}
	return input
}
