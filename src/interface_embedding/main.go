package main

type Jumper interface{ Jump() }
type Shooter interface{ Shot() }

type Character interface {
	Jumper
	Shooter
}

type Rocket struct{}

func (Rocket) Jump() { println("jumping") }
func (Rocket) Shot() { println("shotting") }

func main() {
	var char Character = &Rocket{}
	char.Jump()
	char.Shot()
}
