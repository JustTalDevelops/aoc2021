package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
	"strings"
)

func main() {
	inputs := common.Inputs()
	populations := make([]int, 9)
	for _, input := range strings.Split(inputs[0], ",") {
		populations[common.Atoi(input)]++
	}

	fmt.Println("Part One Answer:", population(common.CopyInts(populations), 80))
	fmt.Println("Part Two Answer:", population(common.CopyInts(populations), 256))
}

// population returns the population after the given number of days.
func population(populations []int, days int) int {
	// Tick each day.
	for i := 0; i < days; i++ {
		// Current population.
		currentPopulation := populations[0]

		// Tick down a day.
		for j := 0; j < len(populations)-1; j++ {
			populations[j] = populations[j+1]
		}

		// Add the new current population.
		populations[8] = currentPopulation
		// Reset the current population.
		populations[6] += currentPopulation
	}
	return common.Sum(populations)
}
