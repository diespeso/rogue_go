package rogue

import "fmt"

type Water struct {
	Cell
}

func NewWater() *Water {
	return &Water{NewCell(WATER)}
}

func (water *Water) OnEnter() {
	fmt.Println("You entered on water")
}
