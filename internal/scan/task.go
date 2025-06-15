package task5

import (
	"fmt"
	"time"
)

func Scan() {
	var birthYear int
	var age int
	currentYear := time.Now().Year()
	age = currentYear - birthYear
	fmt.Print("Введите год рождения: ")
	fmt.Scan(&birthYear)
	fmt.Printf("Ваш возраст: %d лет\n", age)
}
func TimeConverter() {
	var minutes int
	var hours int
	var totalMinutes int

	fmt.Print("Введите минуты:")
	fmt.Scan(&totalMinutes)

	hours = totalMinutes / 60
	minutes = totalMinutes % 60
	fmt.Printf("Результат: %d ч. %d мин.", hours, minutes)
}

func MathTask3() {
	var length float64
	var width float64

	fmt.Print("Введите длину:")
	fmt.Scan(&length)

	fmt.Print("Введите ширину:")
	fmt.Scan(&width)

	square := length * width
	fmt.Printf("Площадь: %0.2f", square)
}
