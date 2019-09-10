package day06

import (
	"fmt"
	"runtime"
	"time"
)

func Task(i interface{}) {
	fmt.Println(i)
}

func TestGoSched() {
	go func() {
		for i := 0; true; i++ {
			time.Sleep(time.Second)
			runtime.Gosched()
			fmt.Println("routine1:", i)
		}
	}()
	go func() {
		for i := 0; true; i++ {
			time.Sleep(time.Second)
			fmt.Println("routine2:", i)
		}
	}()
	select {}
}
