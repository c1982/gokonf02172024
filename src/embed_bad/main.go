package main

import (
	"os"
)

var data []byte

func init() {
	var err error
	data, err = os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
}

func main() {}
