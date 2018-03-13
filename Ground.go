package rogue

import "fmt"

//Ground is the most common steppable cell
type Ground struct {
	Cell
}

//NewGround returns a pointer to a new Ground Cell
func NewGround() *Ground {
	return &Ground{NewCell(GROUND)}
}

func (ground *Ground) OnEnter() {
	fmt.Println("Entered to a ground cell")
}

func (ground *Ground) CanEnter() bool {
	return true
}
