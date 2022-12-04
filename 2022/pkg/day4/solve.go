package day4

import (
	"strconv"
	"strings"
)

func Solve1(input []string) int {
	pairs := 0

	for _, entry := range input {
		elves := strings.Split(entry, ",")

		from1, to1 := getSections(elves[0])
		from2, to2 := getSections(elves[1])

		if areSectionsOverlap(from1, to1, from2, to2) {
			pairs += 1
		}
	}

	return pairs
}

func Solve2(input []string) int {
	return 0
}

func getSections(assigment string) (int, int) {
	sections := strings.Split(assigment, "-")

	s1, _ := strconv.Atoi(sections[0])
	s2, _ := strconv.Atoi(sections[1])

	return s1, s2
}

func areSectionsOverlap(f1, t1, f2, t2 int) bool {
	return (f1 >= f2 && t1 <= t2) || (f2 >= f1 && t2 <= t1)
}
