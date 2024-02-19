package main

import "fmt"

func main() {
	s := "Hello, playground"
	for i := 0; i < 10; i++ {
		s += "," + s
	}
	fmt.Println("Hello, playground")
}
