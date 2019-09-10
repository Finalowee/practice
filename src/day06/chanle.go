package day06

import (
	"fmt"
	"os"
	"time"
)

var ch1 chan int

func init() {
	ch1 = make(chan int, 0)
}

func Producer(i int) {
	go func() {
		for {
			fmt.Println("Producer:", os.Getpid(), i+1)
			ch1 <- -(i + 1)
			time.Sleep(time.Second)
		}
	}()
}

func Consumer() {
	go func() {
		for {
			v := <-ch1
			fmt.Printf("Consumer: %d, %T, %+v\n", os.Getpid(), v, v)
		}
	}()
}
