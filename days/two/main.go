package main

import (
	"fmt"
	"github.com/justtaldevelops/aoc2021/common"
	"strings"
)

func main() {
	input := common.Inputs()

	var horizontalPosition, depth, aim int
	for _, in := range input {
		s := strings.Split(in, " ")
		command, change := s[0], common.Atoi(s[1])

		switch command {
		case "forward":
			horizontalPosition += change
		case "down":
			depth += change
		case "up":
			depth -= change
		}
	}
	fmt.Println("Part One Answer:", horizontalPosition*depth)
	horizontalPosition, depth = 0, 0

	for _, in := range input {
		s := strings.Split(in, " ")
		command, change := s[0], common.Atoi(s[1])

		switch command {
		case "down":
			aim += change
		case "up":
			aim -= change
		case "forward":
			horizontalPosition += change
			depth += aim * change
		}
	}
	fmt.Println("Part Two Answer:", horizontalPosition*depth)
}
