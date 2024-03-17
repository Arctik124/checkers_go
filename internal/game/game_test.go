package game

import (
	"fmt"
	"testing"
)

func TestInitCell(t *testing.T) {
	c := initCell(0, 1)
	if c != Black {
		t.Errorf("c is not black, %d", c)
	}
}

func TestNewBoard(t *testing.T) {
	board := newBoard()
	fmt.Println(board)
}
