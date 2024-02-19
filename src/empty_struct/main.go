package main

import "os"

type State struct{}

func (State) Load() (any, error) {
	return os.ReadFile("game.dat")
}

func (State) Save(s any) error {
	return os.WriteFile("game.dat", s.([]byte), 0644)
}

func main() {
	var stor = State{}
	stor.Save([]byte("state-data"))
}
