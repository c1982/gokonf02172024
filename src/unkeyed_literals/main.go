package main

import (
	"fmt"
)

type Character struct {
	Level  int16
	Health int16
	_      struct{}
}

func main() {

	//c := Character{1, 100}
	c := Character{Level: 1, Health: 100}
	fmt.Printf("%+v\n", c)
}
