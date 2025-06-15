package task8

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Tipization() {
	a := 10 / 3
	b := 10.0 / 3
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("b type: %T and value: %0.2f\n", b, b)
}

func Null() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

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

func String() {
	str := "hêllo"
	z := []rune(str)
	x := []byte(str)
	for _, i := range z {
		fmt.Print(string(i))
	}
	fmt.Print("\n")
	fmt.Println(strings.ToUpper(str))
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(string(str[2]))
	fmt.Println(str[0])
	fmt.Print(z)
	fmt.Print(x)
}

const star = "*"
const space = " "
const tab = "\t"

func Art() {
	printFirstLine()
	printSecondLine()
	printThirdLine()
	printFourLine()
	printThirdLine()
	printSecondLine()
	printFirstLine()
}

func printFirstLine() {
	printSpace(2)
	printStar(4)
	printSpace(2)
	fmt.Println()
}
func printSecondLine() {
	printSpace(1)
	printStar(1)
	printSpace(4)
	printStar(1)
	printSpace(1)
	fmt.Println()
}
func printThirdLine() {
	printStar(1)
	printSpace(1)
	printStar(1)
	printSpace(2)
	printStar(1)
	printSpace(1)
	printStar(1)
	fmt.Println()
}
func printFourLine() {
	printStar(1)
	printSpace(6)
	printStar(1)
	fmt.Println()
}
func printStar(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(star)
	}
}

func printSpace(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(space)
	}
}
