package main

// segment is a line segment.
type segment struct {
	start, end point
}

// point is a 2D point.
type point struct {
	x, z int
}

// countOverlaps counts the number of overlaps between the given segments.
func countOverlaps(segments []segment, diagonals bool) int {
	points := make(map[point]int)
	for _, seg := range segments {
		pos := seg.start
		offset := offsetTowards(pos, seg.end)

		// Ignore diagonals if we're not counting them.
		if offset.x != 0 && offset.z != 0 && !diagonals {
			continue
		}

		for pos.x != seg.end.x+offset.x || pos.z != seg.end.z+offset.z {
			points[pos]++

			pos.x += offset.x
			pos.z += offset.z
		}
	}

	var overlaps int
	for _, count := range points {
		if count > 1 {
			overlaps++
		}
	}
	return overlaps
}

// offsetTowards returns the offset of this point to move towards the given point.
func offsetTowards(current, target point) point {
	return point{sign(current.x, target.x), sign(current.z, target.z)}
}

// sign ...
func sign(one, two int) int {
	if one < two {
		return 1
	} else if one > two {
		return -1
	}
	return 0
}
