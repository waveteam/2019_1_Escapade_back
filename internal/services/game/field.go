package game

import (
	"escapade/internal/models"
	"math/rand"
)

type Field struct {
	Matrix    [][]int
	Width     int
	Height    int
	CellsLeft int
	// Open   int
	// Mines int
	// Flags int
}

/*
type RoomSettings struct {
	Width  int
	Heigth int
	Percent   float32
}
*/

func (f *Field) randomCell() *models.Cell {
	cell := &models.Cell{
		X: rand.Intn(f.Width),
		Y: rand.Intn(f.Height),
	}
	return cell
}

// setFlag add flag to matrix
func setFlag(matrix *[][]int, x int, y int, id int, width int, height int) {
	mine := 9

	// if there was a mine lets reduce dangerous value near it
	if (*matrix)[x][y] == mine {
		for i := x - 1; i <= x+1; i++ {
			if i > 0 && i < width {
				for j := y - 1; j <= y+1; j++ {
					// < mine, not == mine because there can be another flag
					if j > 0 && j < height && (*matrix)[i][j] < mine {
						(*matrix)[i][j]--
					}
				}
			}
		}
	}

	// To identifier which flag we see, lets set id
	// add 10 to id, because if id = 3 we can think that there are 3 mines around
	// we cant use -id, becase in future there will be a lot of conditions with
	// something < 9 (to find not mine places)
	(*matrix)[x][y] = id + 10
}

// setMine add mine to matrix and increase dangerous value in cells near mine
func setMine(matrix *[][]int, x int, y int, width int, height int) {

	mine := 9
	(*matrix)[x][y] = mine
	for i := x - 1; i <= x+1; i++ {
		if i > 0 && i < width {
			for j := y - 1; j <= y+1; j++ {
				if j > 0 && j < height && (*matrix)[i][j] != mine {
					(*matrix)[i][j]++
				}
			}
		}
	}
}

func deleteCell(cells *[]models.Cell, i int) {
	last := len(*cells) - 1
	(*cells)[i] = (*cells)[last] // Copy last element to index i.
	*cells = (*cells)[:last]     // Truncate slice.
}

// fill matrix with mines
func fill(matrix *[][]int, width int, height int, mines int, mineProbability int) {
	freeCells := make([]models.Cell, width*height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cell := freeCells[x+y]
			cell.X = x
			cell.Y = y
		}
	}

	for mines > 0 && len(freeCells) > 0 {
		for deleteIndex, cell := range freeCells {
			if rand.Intn(100) > mineProbability {
				setMine(matrix, cell.X, cell.Y, width, height)
				deleteCell(&freeCells, deleteIndex)
			}
		}
	}
}

// generate matrix
func generate(rs *models.RoomSettings) (mines int, matrix [][]int) {
	width := rs.Width
	height := rs.Height

	matrix = make([][]int, height)
	mines = int(float32(width*height) * rs.Percent)

	fill(&matrix, width, height, mines, int(100*rs.Percent))
	return
}

// NewField create new instance of field
func NewField(rs *models.RoomSettings) *Field {
	mines, matrix := generate(rs)
	field := &Field{
		Matrix:    matrix,
		Width:     rs.Width,
		Height:    rs.Height,
		CellsLeft: rs.Width*rs.Height - mines,
	}
	return field
}

func (f *Field) SetFlag(x int, y int, id int) {
	setFlag(&f.Matrix, x, y, id, f.Width, f.Height)
}
