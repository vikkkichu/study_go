package task7

import "fmt"

func MathTask3() {
	var length float64
	var width float64
	square := lenght * width

	fmt.Print("Введите длину:")
	fmt.Scan(&length)

	fmt.Print("Введите ширину:")
	fmt.Scan(&width)

	fmt.Printf("Площадь: %0.2f", square)
}
