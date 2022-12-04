package day4

import (
	"strconv"
	"strings"
)

func Solve1(input []string) int {
	pairs := 0

	for _, entry := range input {
		elves := strings.Split(entry, ",")

		from1, to1 := getRange(elves[0])
		from2, to2 := getRange(elves[1])

		if areRangesOverlappingFully(from1, to1, from2, to2) {
			pairs += 1
		}
	}

	return pairs
}

func Solve2(input []string) int {
	pairs := 0

	for _, entry := range input {
		elves := strings.Split(entry, ",")

		from1, to1 := getRange(elves[0])
		from2, to2 := getRange(elves[1])

		if areRangesOverlapping(from1, to1, from2, to2) {
			pairs += 1
		}
	}

	return pairs
}

func getRange(assigment string) (from int, to int) {
	sections := strings.Split(assigment, "-")

	from, _ = strconv.Atoi(sections[0])
	to, _ = strconv.Atoi(sections[1])

	return
}

// areRangesOverlappingFully checks if the ranges overlap fully.
//
// Examples:
//
//   - Second range is fully contained in the first one:
//     .2345678.  2-8
//     ..34567..  3-7
//
//   - First range is fully contained in the second one
//     .....6...  6-6
//     ...456...  4-6
func areRangesOverlappingFully(f1, t1, f2, t2 int) bool {
	return (f1 >= f2 && t1 <= t2) || (f2 >= f1 && t2 <= t1)
}

// areRangesOverlapping checks if the ranges overlap.
//
// Examples:
//
//   - Second range is fully contained in the first one:
//     .2345678.  2-8
//     ..34567..  3-7
//
//   - First range is fully contained in the second one
//     .....6...  6-6
//     ...456...  4-6
//
//   - First range is overlapping with the second one partially (a):
//     ....567..  5-7
//     ......789  7-9
//
//   - First range is overlapping with the second one partially (b):
//     ....567..  5-7
//     ...456...  4-6
func areRangesOverlapping(f1, t1, f2, t2 int) bool {
	return (f1 >= f2 && t1 <= t2) ||
		(f2 >= f1 && t2 <= t1) ||
		(f1 <= t2 && t1 >= f2) ||
		(f2 <= t1 && t2 >= f1)
}
