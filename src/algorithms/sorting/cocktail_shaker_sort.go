package sorting

type CocktailShakerSort struct {
}

func (c CocktailShakerSort) Sort(input []int) []int {

	isSwapped := false
	lindex := 0
	rindex := len(input)

	for lindex < rindex {

		isSwapped = false
		for innerloop := lindex; innerloop < rindex-1; innerloop++ {

			if input[innerloop] > input[innerloop+1] {
				input[innerloop], input[innerloop+1] = input[innerloop+1], input[innerloop]
				isSwapped = true
			}
		}

		rindex--

		for outerloop := rindex; outerloop > lindex; outerloop-- {

			if input[outerloop] < input[outerloop-1] {
				input[outerloop-1], input[outerloop] = input[outerloop], input[outerloop-1]
				isSwapped = true
			}
		}

		lindex++

		if !isSwapped {
			return input
		}

	}
	return input
}
