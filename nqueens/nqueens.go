package nqueens

import (
	"fmt"
)

type (
	positiveDiagonalId int
	negativeDiagonalId int
)

type board struct {
	size int
	rep  [][]int
	// Map holding the columns where queen is present
	queenCols map[int]bool
	// Map holding the rows where queen is present
	queenRows map[int]bool
	// Map holding the diagonals where the queen is present
	// positive diagonal means those with positive slope,
	// negative diagonal means those with negative slope.
	queenPositiveDiagonals map[positiveDiagonalId]bool
	queenNegativeDiagonals map[negativeDiagonalId]bool
	computed               bool
	unsolvable             bool
}

func posDiagonalId(row, col int) positiveDiagonalId {
	return positiveDiagonalId(row + col)
}

func negDiagonalId(row, col int) negativeDiagonalId {
	return negativeDiagonalId(row - col)
}

func New(size int) *board {
	b := new(board)
	b.size = size
	b.queenPositiveDiagonals = make(map[positiveDiagonalId]bool)
	b.queenNegativeDiagonals = make(map[negativeDiagonalId]bool)
	b.queenCols = make(map[int]bool)
	b.queenRows = make(map[int]bool)
	// Create the internal represetation for the n*n board
	rep := make([][]int, size)
	for row := 0; row < size; row++ {
		rep[row] = make([]int, size)
	}
	b.rep = rep
	return b
}

func (b *board) HasSolution() bool {
	if !b.computed {
		b.Solve()
	}
	return !b.unsolvable
}

func (b *board) placeQueenInRow(row int, startCol int) bool {
	if startCol >= b.size && row <= 0 {
		return true
	}

	for col := startCol; col < b.size; col++ {
		if b.isSafe(row, col) {
			b.addQueenAt(row, col)
			return false
		}
	}

	// The following happens only if a safe place column for queen wasn't found

	// Backtrack to put queen at a different position in the previous row
	// If not found position, remove queen from row-1, for which
	// we have to first find out the column in the concerned row that
	// contains the queen
	var queenCol int
	// Note that there must be a queen at row row-1 before we place queen
	// at row row, thus this loop shouldn't run infinitely
	for candidate := 0; ; candidate++ {
		if b.rep[row-1][candidate] == 1 {
			queenCol = candidate
			break
		}
	}
	b.removeQueenAt(row-1, queenCol)
	// If the previous backtracking was successful, retry adding queen to the current row
	impossible := b.placeQueenInRow(row-1, queenCol+1)
	if impossible {
		return true
	}
	return b.placeQueenInRow(row, 0)
}

func (b *board) Solve() [][]int {
	if b.computed {
		return b.rep
	}
	var impossible bool
	for row := 0; row < b.size; row++ {
		impossible = b.placeQueenInRow(row, 0)
		if impossible {
			break
		}
	}
	b.unsolvable = impossible
	b.computed = true
	return b.rep
}

func (b *board) isSafe(row, col int) bool {
	posDiagonal := posDiagonalId(row, col)
	negDiagonal := negDiagonalId(row, col)
	digonalSafe := !b.queenPositiveDiagonals[posDiagonal] && !b.queenNegativeDiagonals[negDiagonal]
	rowColSafe := !b.queenCols[col] && !b.queenRows[row]
	return digonalSafe && rowColSafe
}

func (b *board) addRemoveQueen(row, col int, add bool) {
	if add {
		b.rep[row][col] = 1
	} else {
		b.rep[row][col] = 0
	}
	posDiagonal := posDiagonalId(row, col)
	negDiagonal := negDiagonalId(row, col)
	b.queenPositiveDiagonals[posDiagonal] = add
	b.queenNegativeDiagonals[negDiagonal] = add
	b.queenRows[row] = add
	b.queenCols[col] = add
}

func (b *board) addQueenAt(row, col int) {
	b.addRemoveQueen(row, col, true)
}

func (b *board) removeQueenAt(row, col int) {
	b.addRemoveQueen(row, col, false)
}

func (b *board) String() string {
	var toReturn string
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			toReturn = fmt.Sprintf("%s%v ", toReturn, b.rep[r][c])
		}
		if r+1 != b.size {
			toReturn = fmt.Sprintf("%s\n", toReturn)
		}
	}
	return toReturn
}
