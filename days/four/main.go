package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
	"strings"
)

func main() {
	inputs := common.Inputs()
	nums := common.StringsToInts(strings.Split(inputs[0], ","))

	var boards []*board
	for ind := 2; ind < len(inputs)-2; ind += 6 {
		var lines []string
		var grid [][]int
		for i := ind; i < ind+5; i++ {
			lines = append(lines, strings.ReplaceAll(inputs[i], "  ", " "))
		}
		for _, line := range lines {
			grid = append(grid, common.StringsToInts(strings.Split(line, " ")))
		}
		if !(len(grid) == 5 && len(grid[0]) == 5) {
			panic("should never happen")
		}
		boards = append(boards, initializeBoard(grid))
	}

	var winningBoards []*board
	for _, num := range nums {
		for _, b := range boards {
			if !b.HasWinningNumber() {
				if b.Mark(num) {
					if len(winningBoards) == 0 {
						fmt.Println("Part One Answer:", b.SumUnmarked()*num)
					}
					winningBoards = append(winningBoards, b)
					b.SetWinningNumber(num)
				}
			}
		}
	}
	lastBoard := winningBoards[len(winningBoards)-1]
	fmt.Println("Part Two Answer:", lastBoard.SumUnmarked()*lastBoard.WinningNumber())
}
