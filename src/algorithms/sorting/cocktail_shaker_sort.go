package sorting

type CocktailShakerSort struct {
}

func (c CocktailShakerSort) Sort(input []int) []int {

	isSwapper := false
	for outerloop := 0; outerloop < (len(input)+1)/2; outerloop++ {

		isSwapper = false
		for innerloop := outerloop; innerloop < len(input)-1-outerloop; innerloop++ {

			if input[innerloop] > input[innerloop+1] {
				input[innerloop], input[innerloop+1] = input[innerloop+1], input[innerloop]
				isSwapper = true
			}
		}

		if !isSwapper {
			return input
		}

		isSwapper = false
		for innerloop := len(input) - outerloop - 1; innerloop >= outerloop+1; innerloop-- {

			if input[innerloop] < input[innerloop-1] {
				input[innerloop-1], input[innerloop] = input[innerloop], input[innerloop-1]
				isSwapper = true
			}
		}

		if !isSwapper {
			return input
		}

	}
	return input
}
