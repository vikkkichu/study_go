package task10

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

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
