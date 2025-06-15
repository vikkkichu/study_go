package arrays

import (
	"fmt"
)

func Sum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	// 	for _, t := range numbers {
	// 		sum = t + sum
	// 	}
	// 	fmt.Print(sum)
	for i := 0; i < len(numbers); i++ {
		currentNum := numbers[i]
		sum = currentNum + sum
	}
	fmt.Print(sum)
}
