package task8

import "fmt"

func Tipization() {
	a := 10 / 3
	b := 10.0 / 3
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("b type: %T and value: %0.2f\n", b, b)
}
