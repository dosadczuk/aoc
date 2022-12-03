package day1

import (
	"strconv"

	"github.com/dosadczuk/aoc/2022/internal/ds"
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
	vals := ds.NewMaxHeap[int]()

	for _, entry := range input {
		if entry == "" { // new Elf is coming
			vals.Push(sum)
			sum = 0

			continue
		}

		val, _ := strconv.Atoi(entry)
		sum += val
	}
	vals.Push(sum) // add last Elf

	sum = 0
	for i := 1; i <= 3 && !vals.Empty(); i++ {
		sum += vals.Pop()
	}

	return sum
}
