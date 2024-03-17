package game

import "fmt"

func NewGame(player1 string, player2 string) Game {
	return Game{
		player1: player1,
		player2: player2,
		isOver:  false,
		board:   newBoard(),
	}
}

func newBoard() Board {
	cells := make([][]Cell, 8)
	for i := range cells {
		cells[i] = make([]Cell, 8)
	}

	for row := range 8 {
		for col := range 8 {
			cells[row][col] = initCell(row, col)
		}
	}

	return Board{
		cells: cells,
	}
}

func initCell(row, col int) Cell {
	switch row {
	case 0, 2:
		if col%2 == 1 {
			return Black
		} else {
			return Empty
		}
	case 1:
		if col%2 == 0 {
			return Black
		} else {
			return Empty
		}
	case 5, 7:
		if col%2 == 0 {
			return White
		} else {
			return Empty
		}
	case 6:
		if col%2 == 1 {
			return White
		} else {
			return Empty
		}
	default:
		return Empty
	}
}

type Game struct {
	player1     string
	player2     string
	isOver      bool
	board       Board
	isWhiteTurn bool
}

func (game Game) String() string {
	result := fmt.Sprintf("%s VS %s\n\n", game.player1, game.player2)

	result += game.board.String()
	result += "\n\n"
	if game.isWhiteTurn {
		result += game.player1
	} else {
		result += game.player2
	}
	result += " turn to move!"

	return result
}

type Board struct {
	cells [][]Cell
}

func (board Board) String() string {
	var result string
	for _, row := range board.cells {
		for _, col := range row {
			result += col.String()
		}
		result += "\n"
	}
	return result
}

type Cell int32

const (
	Empty      Cell = 0
	White      Cell = 1
	Black      Cell = 2
	WhiteQueen Cell = 3
	BlackQueen Cell = 4
)

func (c Cell) String() string {
	switch c {
	case Empty:
		return "[ ]"
	case White:
		return "[w]"
	case Black:
		return "[b]"
	case WhiteQueen:
		return "[W]"
	case BlackQueen:
		return "[B]"
	default:
		return "[ ]"
	}

}
