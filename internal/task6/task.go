package task6

import "fmt"

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
