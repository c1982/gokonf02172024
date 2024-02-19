package main

import "unsafe"

type Character1 struct {
	isAlive  bool    // 1 byte   (7 byte)
	Health   int64   // 8 bytes  (0 byte)
	isBanned bool    // 1 bytes  (7 byte)
	Name     string  // 16 bytes (0 byte)
	Coin     float32 // 4 bytes  (4 byte)
}

type Character2 struct {
	Name     string  // 16 bytes (0 byte)
	Health   int64   // 8 bytes  (0 byte)
	Coin     float32 // 4 bytes  (2 byte)
	isAlive  bool    // 1 byte
	isBanned bool    // 1 bytes
}

func main() {
	c1, c2 := Character1{}, Character2{}
	println("Character 1:", unsafe.Sizeof(c1))
	println("Character 2:", unsafe.Sizeof(c2))
}

//END OMIT
