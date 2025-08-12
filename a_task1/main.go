package main

type Human struct{}

func (h Human) Walk()

type Action struct {
	Human
}

func main() {
	Action{}.Human.Walk()
}
