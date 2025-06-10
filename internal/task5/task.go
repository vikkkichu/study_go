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
