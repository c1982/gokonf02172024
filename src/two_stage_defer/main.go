package main

import (
	"fmt"
	"time"
)

func metric() func() {
	begin := time.Now()

	return func() {
		end := time.Now().Sub(begin)
		fmt.Printf("execution time: %v", end)
	}
}

func main() {
	defer metric()()
	time.Sleep(time.Second)
}
