package sorting

type BubbleSort struct {
}

func (b BubbleSort) Sort(input []int) []int {

	isSwapped := false
	for outerloop := 0; outerloop < len(input); outerloop++ {

		isSwapped = false

		for innerloop := 0; innerloop < len(input)-1-outerloop; innerloop++ {

			if input[innerloop+1] < input[innerloop] {

				input[innerloop], input[innerloop+1] = input[innerloop+1], input[innerloop]
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}
	return input
}
