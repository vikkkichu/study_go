package task1

import (
	"fmt"
	"math"
	"time"
)

func MathTask() {
	var r float64 = 5.5
	var area float64 = math.Pi * math.Pow(r, 2)
	fmt.Printf(" Площадь круга: %0.2f ", area)
}

func MathTask2() {
	distance := 150.5
	speed := 80.0
	t := distance / speed
	totalMinutes := t * 60
	d := time.Duration(totalMinutes) * time.Minute
	fmt.Printf("При скорости %v км/ч, расстояние %v км будет преодолено за %v минуты", speed, distance, d.Hours())
}

func MathTask3() {
	firstTemp := -5.3
	secondTemp := 2.4
	thirdTemp := 10.2
	var avg float64 = (firstTemp + secondTemp + thirdTemp) / 3
	fmt.Printf("Средняя температура: %0.1f °C", avg)
}

func MathTask5() {
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
