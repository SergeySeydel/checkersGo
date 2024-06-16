package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// 0, 1, 2
const (
	empty = iota
	white
	black
)

// Класс доска
type Board struct {
	Grid          [][]int `json:"grid"`
	ScoreWhite    int     `json:"scorewhite"`
	ScoreBlack    int     `json:"scoreblack"`
	CurrentPlayer int     `json:"currentplayer"`
	Move          string  `json:"move"`
	// StartRow      int     `json:"startrow"`
	// EndRow        int     `json:"endrow"`
	// StartCol      int     `json:"startcol"`
	// EndCol        int     `json:"endcol"`

	// Turn          int     `json:"turn"`
}

type BoardWhite struct {
	Grid       [][]int `json:"grid"`
	ScoreWhite int     `json:"scorewhite"`
}

type BoardBlack struct {
	Grid       [][]int `json:"grid"`
	ScoreBlack int     `json:"scoreblack"`
}

// Конструктор
func NewBoard() *Board {
	b := &Board{
		Grid:          make([][]int, 8),
		ScoreWhite:    0,
		ScoreBlack:    0,
		CurrentPlayer: 1,
		// StartRow:      0,
		// EndRow:        0,
		// StartCol:      0,
		// EndCol:        0,
		// Turn:          0,
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

func (b *BoardWhite) setup() {
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

func (b *BoardBlack) setup() {
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

func (b *Board) isValidMove(startRow, startCol, endRow, endCol int) bool {
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
	if b.CurrentPlayer == white && rowDiff == -1 && abs(colDiff) == 1 {
		return true
	} else if b.CurrentPlayer == black && rowDiff == 1 && abs(colDiff) == 1 {
		return true
	}

	// Обработка перепрыгивания через противника
	if abs(rowDiff) == 2 && abs(colDiff) == 2 {
		midRow := (startRow + endRow) / 2
		midCol := (startCol + endCol) / 2
		if b.Grid[midRow][midCol] != empty && b.Grid[midRow][midCol] != b.CurrentPlayer {
			return true
		}
	}

	return false
}

// Съесть пешку
func (b *Board) move(startRow, startCol, endRow, endCol int) {
	// Очистка доски от взятой шашки и т.д.
	if abs(endRow-startRow) == 2 && abs(endCol-startCol) == 2 {
		midRow := (startRow + endRow) / 2
		midCol := (startCol + endCol) / 2
		b.Grid[midRow][midCol] = empty
		if b.CurrentPlayer == black {
			b.ScoreBlack += 1
		} else if b.CurrentPlayer == white {
			b.ScoreWhite += 1
		}
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

func (b *Board) getAllHandler(context *gin.Context) {
	if b.CurrentPlayer == white {
		var w BoardWhite = BoardWhite{Grid: make([][]int, 8)}
		for i := range w.Grid {
			w.Grid[i] = make([]int, 8)
		}
		w.setup()
		copy(w.Grid, b.Grid)
		w.ScoreWhite = b.ScoreWhite
		context.JSON(http.StatusOK, w)
	} else if b.CurrentPlayer == black {
		var bl BoardBlack = BoardBlack{Grid: make([][]int, 8)}
		for i := range bl.Grid {
			bl.Grid[i] = make([]int, 8)
		}
		bl.setup()
		copy(bl.Grid, b.Grid)
		bl.ScoreBlack = b.ScoreBlack
		context.JSON(http.StatusOK, bl)
	}

}

func (b *Board) createHandler(context *gin.Context) {
	type Move struct {
		Move string `json:"move"`
	}
	var m Move
	if err := context.ShouldBindJSON(&m); err != nil {
		context.String(http.StatusBadRequest, err.Error())
	}
	b.createMoveForHandler(m.Move)
	context.JSON(http.StatusOK, b)
}

func (b *Board) createMoveForHandler(move string) {
	startRow, startCol, endRow, endCol, err := parseMove(move)
	if err != nil {
		fmt.Println("Invalid move format. Use format A1-B2")
		// continue
	}

	if b.isValidMove(startRow, startCol, endRow, endCol) {
		b.move(startRow, startCol, endRow, endCol)
		if b.CurrentPlayer == white {
			b.CurrentPlayer = black
		} else {
			b.CurrentPlayer = white
		}
	} else {
		fmt.Println("Invalid move. Try again.")
	}
}

// func (b *Board) serialize(f *os.File) {
// 	bytes, err := json.Marshal(b)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, errwr := f.Write(bytes)
// 	if errwr != nil {
// 		log.Fatal(err)
// 	}
// 	b.Turn += 1
// }

func main() {
	board := NewBoard()
	// player := white
	// file, err := os.OpenFile("serialized.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	router := gin.Default()

	router.GET("/", board.getAllHandler)
	router.POST("/", board.createHandler)

	router.Run("127.0.0.1:8080")

	for {
		board.Print()
		var move string
		// if board.CurrentPlayer == white {
		// 	fmt.Printf("White's score: %d\n", board.ScoreWhite)
		// 	fmt.Printf("Black's score: %d\n", board.ScoreBlack)
		// 	fmt.Print("White's move: ")

		// } else {
		// 	// board.CurrentPlayer = 2
		// 	fmt.Printf("White's score: %d\n", board.ScoreWhite)
		// 	fmt.Printf("Black's score: %d\n", board.ScoreBlack)
		// 	fmt.Print("Black's move: ")
		// }
		// board.serialize(file)
		fmt.Scanln(&move)
		// startRow, startCol, endRow, endCol, err := parseMove(move)
		// board.StartRow = startRow
		// board.StartCol = startCol
		// board.EndRow = endCol
		// board.EndCol = endRow
		// if err != nil {
		// 	fmt.Println("Invalid move format. Use format A1-B2")
		// 	continue
		// }

		// if board.isValidMove(startRow, startCol, endRow, endCol) {
		// 	board.move(startRow, startCol, endRow, endCol)
		// 	if board.CurrentPlayer == white {
		// 		board.CurrentPlayer = black
		// 		// board.CurrentPlayer = black
		// 	} else {
		// 		board.CurrentPlayer = white
		// 		// board.CurrentPlayer = white
		// 	}
		// } else {
		// 	fmt.Println("Invalid move. Try again.")
		// }

	}

}
