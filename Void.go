package rogue

import "log"

//Void is the cell that cannot be stepped on
type Void struct {
	Cell
}

func NewVoid() *Void {
	return &Void{NewCell(VOID)}
}

func (void *Void) OnEnter() {
	log.Fatal("CANNOT ENTER TO A VOID CELL")
}
