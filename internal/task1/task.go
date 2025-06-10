package task1

import (
	"fmt"
	"math"
)

func MathTask() {
	var r float64 = 5.5
	var area float64 = math.Pi * math.Pow(r, 2)
	fmt.Printf(" Площадь круга: %0.2f ", area)
}
