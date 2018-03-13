package rogue

import (
	"testing"
)

func TestNewCell(t *testing.T) {
	newCell := NewCell(GROUND)
	newCell.ShowInfo()

}
