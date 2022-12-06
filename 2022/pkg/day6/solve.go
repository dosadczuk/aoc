package day6

import "github.com/dosadczuk/go-ds/set"

func Solve1(input []string) int {
	return getMarkerAfterNthUniqueCharacter(input[0], 4)
}

func Solve2(input []string) int {
	return getMarkerAfterNthUniqueCharacter(input[0], 14)
}

func getMarkerAfterNthUniqueCharacter(entry string, n int) int {
	chars := set.NewHashSet[byte]()
	for i, j := 0, 0; j < len(entry); j += 1 {
		for chars.Has(entry[j]) {
			chars.Remove(entry[i])
			i += 1
		}

		chars.Add(entry[j])

		if chars.Len() == n {
			return j + 1
		}
	}

	return -1
}
