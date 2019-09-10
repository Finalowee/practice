package main

import (
	"day06"
	"math/rand"
)

func main() {
	//day05.Copy("./src/day05/copy.go", "./src/day05/copy-bak.go")
	//day06.TestGoSched()

	day06.Producer(rand.Intn(10))
	day06.Consumer()
	select {}
}
