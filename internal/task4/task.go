package task4

import "fmt"

func MathTask2() {
	firstTemp := -5.3
	secondTemp := 2.4
	thirdTemp := 10.2
	var avg float64 = (firstTemp + secondTemp + thirdTemp) / 3
	fmt.Printf("Средняя температура: %0.1f °C", avg)
}

func MathTask4() {
	firstTemp := -5.3
	secondTemp := 2.4
	thirdTemp := 10.2
	allTemp := []float64{firstTemp, secondTemp, thirdTemp, 8, 7, 11, 20}
	var avg float64
	sum := 0.0
	for _, t := range allTemp {
		sum = t + sum
	}
	avg = sum / float64(len(allTemp))

	fmt.Printf("Средняя температура: %0.1f °C", avg)
}
