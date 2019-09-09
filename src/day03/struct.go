package day03

import "fmt"

type Student struct {
	id     int
	name   string
	gender int8
}

func Run() {
	var lee Student
	lee.id = 1
	lee.gender = 0
	lee.name = "lee"
	p := &lee
	(&lee).name = "LoLi"

	fmt.Printf("%T %T", lee, p)
}
