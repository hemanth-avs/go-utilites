package sorting

type SelectionSort struct {
}

func (s SelectionSort) Sort(input []int) []int {

	min := 0
	minIndex := 0

	for index := 0; index < len(input); index++ {

		min = input[index]
		minIndex = index

		for loop := index; loop < len(input)-1; loop++ {
			if input[loop+1] < min {
				min = input[loop+1]
				minIndex = loop + 1
			}
		}
		input[minIndex], input[index] = input[index], input[minIndex]
	}

	return input
}
