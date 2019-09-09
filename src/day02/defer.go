package day02

import "fmt"

func Order() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
