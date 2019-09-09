package day01

import "fmt"

func Say(msg string) {
	fmt.Println(msg)
}

func Declare() {
	var a int = 1
	var b int
	b = 2
	c := 3
	fmt.Printf("%T %T %T\n", a, b, c)
	fmt.Printf("%d %d %d\n", a, b, c)
}
