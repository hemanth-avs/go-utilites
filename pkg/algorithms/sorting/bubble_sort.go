package sorting

import "fmt"

func Bubblesort(input []int) []int {

  input_copy := make([]int, len(input))
  copy(input_copy, input)

  for outerloop:=0; outerloop<len(input_copy); outerloop++ {
    for innerloop:=0; innerloop<len(input_copy)-1-outerloop; innerloop++ {
      if input_copy[innerloop+1] < input_copy[innerloop] {
        temp := input_copy[innerloop]
        input_copy[innerloop] = input_copy[innerloop+1]
        input_copy[innerloop+1] = temp
      }
    }
  }
  return input_copy
}
