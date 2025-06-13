package task9

import (
	"fmt"
)

func Rune() {
	var char rune = 'A'
	str := string(char)
	fmt.Printf("%v, %s", char, str)
}

func ToBool() {
	num := 0
	newBool := num == 0
	fmt.Println(newBool)
}

const (
	Monday = iota + 1
	Tuesday
	Wendsday
)

func Const() {
	fmt.Printf("Понедельник = %v\n", Monday)
}
