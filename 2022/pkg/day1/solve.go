package day1

import (
	"sort"
	"strconv"

	"github.com/dosadczuk/aoc/2022/internal/math"
)

func Solve1(input []string) int {
	// How many Calories are those Elves carrying in total?

	sum := 0
	max := 0

	for _, entry := range input {
		if entry == "" { // new Elf is coming
			max = math.MaxInt(max, sum)
			sum = 0

			continue
		}

		val, _ := strconv.Atoi(entry)
		sum += val
	}

	return math.MaxInt(sum, max)
}

func Solve2(input []string) int {
	// How many Calories are those Elves carrying in total?

	sum := 0
	vals := []int{}

	for _, entry := range input {
		if entry == "" { // new Elf is coming
			vals = append(vals, sum)
			sum = 0

			continue
		}

		val, _ := strconv.Atoi(entry)
		sum += val
	}

	// add last Elf
	vals = append(vals, sum)

	sort.Slice(vals, func(i, j int) bool {
		return vals[j] < vals[i] // sort descending
	})

	sum = 0
	for i := 0; i < math.MinInt(3, len(vals)); i++ {
		sum += vals[i]
	}

	return sum
}
