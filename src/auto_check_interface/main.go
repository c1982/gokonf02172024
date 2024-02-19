package main

type Character interface {
	Level(int)
}

type Rocket struct{}

func (c *Rocket) Level(l int) { println("Level:", l) }

var _ Character = (*Rocket)(nil)

func main() {
	rocket := new(Rocket)

	if c, ok := any(rocket).(Character); ok {
		c.Level(10)
	}
}
