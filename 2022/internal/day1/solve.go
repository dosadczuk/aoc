package day1

import (
	"strconv"

	"github.com/dosadczuk/aoc/2022/pkg/math"
)

func Solve(input []string) int {
	// How many Calories are those Elves carrying in total?

	caloriesSum := 0
	caloriesMax := 0

	for _, entry := range input {
		if entry == "" { // new Elf is coming
			caloriesMax = math.MaxInt(caloriesMax, caloriesSum)
			caloriesSum = 0

			continue
		}

		calories, _ := strconv.Atoi(entry)
		caloriesSum += calories
	}

	return math.MaxInt(caloriesSum, caloriesMax)
}
