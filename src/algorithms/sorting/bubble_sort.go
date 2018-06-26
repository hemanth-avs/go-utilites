package sorting

type BubbleSort struct {
}

func (b BubbleSort) Sort(input []int) []int {

	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	isSwapped := false
	for outerloop := 0; outerloop < len(inputCopy); outerloop++ {

		isSwapped = false

		for innerloop := 0; innerloop < len(inputCopy)-1-outerloop; innerloop++ {

			if inputCopy[innerloop+1] < inputCopy[innerloop] {
				temp := inputCopy[innerloop]
				inputCopy[innerloop] = inputCopy[innerloop+1]
				inputCopy[innerloop+1] = temp
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}
	return inputCopy
}
