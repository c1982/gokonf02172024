package main

type Character struct {
	_      [0]func()
	Level  float64
	Health float64
}

func main() {
	c1 := Character{Level: 1, Health: 100}
	c2 := Character{Level: 1, Health: 100}
	println(c1 == c2)
}
