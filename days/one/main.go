package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
)

func main() {
	inputs := common.InputsAsInts()

	fmt.Println("Part One Answer:", increases(inputs))
	fmt.Println("Part Two Answer:", increases(threeMeasurementSums(inputs)))
}

// increases returns the number of increases in a slice of ints.
func increases(inputs []int) int {
    var lastNum, increased int
    for i, num := range inputs {
        if i > 0 {
            if lastNum < num {
                increased++
            }
        }
        lastNum = num
    }

    return increased
}

// threeMeasurementSums returns an int of three-measurement sums from the given slice of ints.
func threeMeasurementSums(inputs []int) []int {
	var sums []int
	for i := 0; i < len(inputs) - 2; i++ {
		if len(inputs) < i + 2 {
			break
		}

		sums = append(sums, inputs[i] + inputs[i + 1] + inputs[i + 2])
	}
	return sums
}
