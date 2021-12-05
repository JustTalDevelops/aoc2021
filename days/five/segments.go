package main

// segment is a line segment.
type segment struct {
	start, end point
}

// countOverlaps counts the number of overlaps between the given segments.
func countOverlaps(segments []segment, diagonals bool) int {
	points := make(map[point]int)
	for _, seg := range segments {
		pos := seg.start
		offset := offsetForTarget(pos, seg.end)

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

// sign ...
func sign(one, two int) int {
	if one < two {
		return 1
	} else if one > two {
		return -1
	}
	return 0
}
