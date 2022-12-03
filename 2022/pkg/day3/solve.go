package day3

import (
	"unicode"

	"github.com/dosadczuk/aoc/2022/internal/ds"
)

func Solve1(input []string) int {
	priority := 0

	for _, entry := range input {
		var half int = len(entry) / 2

		l := getCharStats(entry[:half])
		r := getCharStats(entry[half:])

		i := ds.SetIntersection(l, r)

		p := 0
		for _, v := range i.Vals() {
			if unicode.IsLower(v) {
				// Lowercase item types a through z have priorities 1 through 26.
				p += int(v - 'a' + 01)
			} else {
				// Uppercase item types A through Z have priorities 27 through 52.
				p += int(v - 'A' + 27)
			}
		}

		priority += p
	}

	return priority
}

func Solve2(input []string) int {
	return 0
}

func getCharStats(value string) *ds.Set[rune] {
	stats := ds.NewSet[rune]()

	for _, char := range value {
		stats.Add(char)
	}

	return stats
}
