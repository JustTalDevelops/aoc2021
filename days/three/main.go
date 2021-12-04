package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
	"strings"
)

func main() {
	columns := common.Inputs()
	firstColumn := columns[0]

	var gammaRayBits, epsilonRateBits strings.Builder
	for column := 0; column < len(firstColumn); column++ {
		if positiveBits, negativeBits := bits(columns, column); positiveBits > negativeBits {
			gammaRayBits.WriteRune('1')
			epsilonRateBits.WriteRune('0')
		} else {
			gammaRayBits.WriteRune('0')
			epsilonRateBits.WriteRune('1')
		}
	}

	oxygenGeneratorRatingBits := shorten(common.Copy(columns), func(positiveBits, negativeBits int) bool {
		return positiveBits < negativeBits
	})
	carbonDioxideScrubberRatingBits := shorten(common.Copy(columns), func(positiveBits, negativeBits int) bool {
		return negativeBits < positiveBits || negativeBits == positiveBits
	})

	gammaRay, epsilonRate := common.IntFromBinary(gammaRayBits.String()), common.IntFromBinary(epsilonRateBits.String())
	oxygenGeneratorRating, carbonDioxideScrubberRating := common.IntFromBinary(oxygenGeneratorRatingBits), common.IntFromBinary(carbonDioxideScrubberRatingBits)

	fmt.Println("Part One Answer:", gammaRay*epsilonRate)
	fmt.Println("Part Two Answer:", oxygenGeneratorRating*carbonDioxideScrubberRating)
}

// shorten removes the columns that aren't a part of the solution, or don't pass the check.
func shorten(columns []string, check func(positiveBits, negativeBits int) bool) string {
	var bitPosition int
	for len(columns) > 1 {
		bit := uint8('0')
		if check(bits(columns, bitPosition)) {
			bit = '1'
		}

		for i := 0; i < len(columns); i++ {
			if columns[i][bitPosition] == bit {
				columns = append(columns[:i], columns[i+1:]...)
				i--
			}
		}
		bitPosition++
	}
	return columns[0]
}

// bits returns the positive and negative bits of a column.
func bits(columns []string, column int) (int, int) {
	var positiveBits, negativeBits int
	for _, input := range columns {
		if input[column] == '1' {
			positiveBits++
		} else {
			negativeBits++
		}
	}
	return positiveBits, negativeBits
}