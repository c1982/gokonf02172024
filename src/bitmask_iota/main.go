package main

import "fmt"

type States uint8

const (
	Jump     States = 1 << iota // 1 << 0 00000001
	Freeze                      // 1 << 1 00000010
	Fire                        // 1 << 2 00000100
	IsGround                    // 1 << 3 00001000
)

func main() {
	input := Jump | Fire // 00000101

	fmt.Printf("%d=%08b\n", input, input)

	if input&Jump == 1 {
		fmt.Println("Jump is set")
	}

	if input&IsGround != 0 {
		fmt.Println("Ground is set")
	}
}

// END OMIT
