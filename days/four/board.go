package main

// board contains information about the current state of the bingo board.
type board struct {
	hasWinningNumber bool
	winningNumber    int

	markedGrid [][]bool
	grid       [][]int
}

// initializeBoard initializes the board with the given grid.
func initializeBoard(grid [][]int) *board {
	markedGrid := make([][]bool, len(grid))
	for i := range markedGrid {
		markedGrid[i] = make([]bool, len(grid[i]))
	}
	return &board{grid: grid, markedGrid: markedGrid}
}

// Bingo returns true if the board contains a bingo.
func (b *board) Bingo() (won bool) {
	for x := 0; x < len(b.grid); x++ {
		won = true
		for z := 0; z < len(b.grid[x]); z++ {
			if _, marked := b.At(x, z); !marked {
				won = false
				break
			}
		}
		if won {
			return
		}
	}
	for y := 0; y < len(b.grid[0]); y++ {
		won = true
		for x := 0; x < len(b.grid); x++ {
			if _, marked := b.At(x, y); !marked {
				won = false
				break
			}
		}
		if won {
			return
		}
	}
	return
}

// Grid is a 5x5 grid of [][]int.
func (b *board) Grid() [][]int {
	return b.grid
}

// SetWinningNumber sets the winning number.
func (b *board) SetWinningNumber(winningNumber int) {
	b.hasWinningNumber = true
	b.winningNumber = winningNumber
}

// WinningNumber returns the winning number.
func (b *board) WinningNumber() int {
	return b.winningNumber
}

// HasWinningNumber returns true if the board's winning number has been set.
func (b *board) HasWinningNumber() bool {
	return b.hasWinningNumber
}

// Mark marks the given row and column on the board.
func (b *board) Mark(num int) bool {
	for x := 0; x < len(b.grid); x++ {
		for z := 0; z < len(b.grid[x]); z++ {
			if b.grid[x][z] == num {
				b.markedGrid[x][z] = true
			}
		}
	}
	return b.Bingo()
}

// SumUnmarked returns the sum of all unmarked numbers in the board.
func (b *board) SumUnmarked() (sum int) {
	for x := 0; x < len(b.grid); x++ {
		for z := 0; z < len(b.grid[x]); z++ {
			if _, marked := b.At(x, z); !marked {
				sum += b.grid[x][z]
			}
		}
	}
	return
}

// At returns the value at a position on a board.
func (b *board) At(x, z int) (int, bool) {
	return b.grid[x][z], b.markedGrid[x][z]
}
