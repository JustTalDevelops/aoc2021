package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
	"strings"
)

func main() {
	inputs := common.Inputs()
	segments := make([]segment, 0)
	for _, line := range inputs {
		unparsedPoints := strings.Split(line, " -> ")
		segments = append(segments, segment{
			start: parsePoint(unparsedPoints[0]),
			end:   parsePoint(unparsedPoints[1]),
		})
	}

	fmt.Println("Part One Answer:", countOverlaps(segments, false))
	fmt.Println("Part Two Answer:", countOverlaps(segments, true))
}
