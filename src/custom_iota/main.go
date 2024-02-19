package main

type Level byte

func (l Level) Name() string {
	return [...]string{"Gold", "Silver", "Platinum"}[l]
}

const (
	Gold Level = iota
	Silver
	Platinum
)

func main() {
	println(Gold.Name())
	println(Silver.Name())
	println(Platinum.Name())
}
