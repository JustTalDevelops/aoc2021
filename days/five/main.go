package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
)

func main() {
	inputs := common.Inputs()
	segments := make([]segment, 0)
	for _, line := range inputs {
		var seg segment

		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &seg.start.x, &seg.start.z, &seg.end.x, &seg.end.z)
		if err != nil {
			panic(err)
		}
		segments = append(segments, seg)
	}

	fmt.Println("Part One Answer:", countOverlaps(segments, false))
	fmt.Println("Part Two Answer:", countOverlaps(segments, true))
}
