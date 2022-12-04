package day3

import (
	"unicode"

	"github.com/dosadczuk/go-ds/ds"
)

func Solve1(input []string) int {
	total := 0

	for _, entry := range input {
		var half int = len(entry) / 2

		chars := ds.HashSetIntersection(
			getUniqueChars(entry[:half]), // left half
			getUniqueChars(entry[half:]), // right half
		)

		total += getCharPriority(chars)
	}

	return total
}

func Solve2(input []string) int {
	total := 0

	chars := ds.NewHashSet[rune]()
	for i, entry := range input {
		if chars.Empty() { // is new group
			chars = getUniqueChars(entry)
		} else {
			chars = ds.HashSetIntersection(getUniqueChars(entry), chars)
		}

		if ord := i + 1; ord%3 == 0 || ord == len(input) {
			total += getCharPriority(chars)
			chars = ds.NewHashSet[rune]()
		}
	}

	return total
}

func getUniqueChars(v string) *ds.HashSet[rune] {
	chars := ds.NewHashSet[rune]()
	for _, char := range v {
		chars.Add(char)
	}

	return chars
}

func getCharPriority(s *ds.HashSet[rune]) int {
	for _, v := range s.Vals() { // first value only
		if unicode.IsLower(v) {
			return int(v - 'a' + 01) // priority 01 through 26
		} else {
			return int(v - 'A' + 27) // priority 27 through 52
		}
	}

	return 0 // default int
}
