package main

import (
	"fmt"
)

// 0, 1, 2
const (
	empty = iota
	white
	black
)

// Класс доска
type Board struct {
	Grid [][]int `json:"grid"`
}

type ApiExtension struct {
	Id         int     `json:"id"`
	ScoreWhite int     `json:"scorewhite"`
	ScoreBlack int     `json:"scoreblack"`
	Board      [][]int `json:"board"`
}

// func NewApiExt() *ApiExtension {
// 	a := &ApiExtension{}

// }

// Конструктор
func NewBoard() *Board {
	b := &Board{
		Grid: make([][]int, 8),
	}
	for i := range b.Grid {
		b.Grid[i] = make([]int, 8)
	}
	b.setup()
	return b
}

// Установка шашек
func (b *Board) setup() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 1 && i < 3 {
				b.Grid[i][j] = black
			} else if (i+j)%2 == 1 && i > 4 {
				b.Grid[i][j] = white
			} else {
				b.Grid[i][j] = empty
			}
		}
	}
}

// Печать в консоль
func (b *Board) Print() {
	fmt.Print("  ")
	for i := 0; i < 8; i++ {
		r := int('A')
		si := r + i
		fmt.Print(string(rune(si)))
		fmt.Print(" ")
	}
	fmt.Println()
	for i := 0; i < 8; i++ {
		fmt.Print(8 - i)
		fmt.Print(" ")
		for j := 0; j < 8; j++ {
			switch b.Grid[i][j] {
			case white:
				fmt.Print("W ")
			case black:
				fmt.Print("B ")
			default:
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) isValidMove(startRow, startCol, endRow, endCol int, player int) bool {
	//Если значения не вылезают за пределы
	if startRow < 0 || startRow >= 8 || startCol < 0 || startCol >= 8 || endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
		return false
	}
	//Если место пустое
	if b.Grid[endRow][endCol] != empty {
		return false
	}
	//Разница строк
	rowDiff := endRow - startRow
	//Разница столбцов
	colDiff := endCol - startCol

	// white идет вверх
	if player == white && rowDiff == -1 && abs(colDiff) == 1 {
		return true
	} else if player == black && rowDiff == 1 && abs(colDiff) == 1 {
		return true
	}

	// Проверка взятия шашки
	if abs(rowDiff) == 2 && abs(colDiff) == 2 {
		midRow := (startRow + endRow) / 2
		midCol := (startCol + endCol) / 2
		if b.Grid[midRow][midCol] != empty && b.Grid[midRow][midCol] != player {
			return true
		}
	}

	return false
}

func (b *Board) move(startRow, startCol, endRow, endCol int) {
	// Проверка взятия шашки
	if abs(endRow-startRow) == 2 && abs(endCol-startCol) == 2 {
		midRow := (startRow + endRow) / 2
		midCol := (startCol + endCol) / 2
		b.Grid[midRow][midCol] = empty
	}
	b.Grid[endRow][endCol] = b.Grid[startRow][startCol]
	b.Grid[startRow][startCol] = empty
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseMove(move string) (int, int, int, int, error) {
	if len(move) != 5 || move[2] != '-' {
		return 0, 0, 0, 0, fmt.Errorf("invalid move format")
	}
	startCol := int(move[0] - 'A')
	startRow := 8 - int(move[1]-'0')
	endCol := int(move[3] - 'A')
	endRow := 8 - int(move[4]-'0')
	return startRow, startCol, endRow, endCol, nil
}

func main() {
	board := NewBoard()
	player := white

	for {
		board.Print()
		var move string
		if player == white {
			fmt.Print("White's move: ")
		} else {
			fmt.Print("Black's move: ")
		}
		fmt.Scanln(&move)

		startRow, startCol, endRow, endCol, err := parseMove(move)
		if err != nil {
			fmt.Println("Invalid move format. Use format A1-B2")
			continue
		}

		if board.isValidMove(startRow, startCol, endRow, endCol, player) {
			board.move(startRow, startCol, endRow, endCol)
			if player == white {
				player = black
			} else {
				player = white
			}
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}
