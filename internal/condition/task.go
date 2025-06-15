package task2

import "fmt"

func WorkDay() {
	isWorkingDay := true
	isWeekend := false
	var answer string
	if isWorkingDay && !isWeekend {
		answer = "да"
	} else {
		answer = "нет"
	}
	fmt.Println("Идти на работу?", answer)
}
