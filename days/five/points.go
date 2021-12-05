package main

import (
	"github.com/justtaldevelops/aoc2021/common"
	"strings"
)

// point is a 2D point.
type point struct {
	x, z int
}

// parsePoint parses a point from a string.
func parsePoint(s string) point {
	sp := strings.Split(s, ",")
	return point{x: common.Atoi(sp[0]), z: common.Atoi(sp[1])}
}

// offsetForTarget returns the offset of this point to move towards the given point.
func offsetForTarget(current, target point) point {
	return point{sign(current.x, target.x), sign(current.z, target.z)}
}
