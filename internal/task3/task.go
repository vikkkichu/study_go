package task3

import (
	"fmt"
	"strconv"
)

func MskToSch() {
	city := "Сочи"
	distance := 1620
	str := strconv.Itoa(distance)
	fmt.Printf("Расстояние от Москвы до %s — %s км.", city, str)
}
