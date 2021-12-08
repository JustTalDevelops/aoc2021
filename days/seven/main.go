package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
	"strings"
)

func main() {
	positions := common.StringsToInts(strings.Split(common.Inputs()[0], ","))

	median := common.Median(positions)
	average := common.Average(positions)

	var partOneCost, partTwoCost int
	for _, pos := range positions {
		partOneDiff, partTwoDiff := common.Abs(pos-median), common.Abs(pos-average)

		partOneCost += partOneDiff
		partTwoCost += ((partTwoDiff * partTwoDiff) + partTwoDiff) / 2
	}

	fmt.Println("Part One Answer:", partOneCost)
	fmt.Println("Part Two Answer:", partTwoCost)
}
