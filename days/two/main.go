package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
)

func main() {
	input := common.InputsSplit(" ")

	var positionOne, positionTwo, depthOne, depthTwo, aim int
	for _, in := range input {
		command, change := in[0], common.Atoi(in[1])

		switch command {
		case "forward":
			positionOne += change
			positionTwo += change
			depthTwo += aim * change
		case "down":
			depthOne += change
			aim += change
		case "up":
			depthOne -= change
			aim -= change
		}
	}
	fmt.Println("Part One Answer:", positionOne*depthOne)
	fmt.Println("Part Two Answer:", positionTwo*depthTwo)
}
