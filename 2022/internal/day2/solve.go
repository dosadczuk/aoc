package day2

import "strings"

type shape int

const (
	rock shape = iota
	paper
	scissors
)

var (
	myShape = map[string]shape{"X": rock, "Y": paper, "Z": scissors}
	urShape = map[string]shape{"A": rock, "B": paper, "C": scissors}
)

const (
	lose = 0
	draw = 3
	win  = 6
)

var (
	scoreForShape = map[shape]int{rock: 1, paper: 2, scissors: 3}
	scoreForRound = map[shape]map[shape]int{
		rock:     {rock: draw, paper: lose, scissors: win},
		paper:    {rock: win, paper: draw, scissors: lose},
		scissors: {rock: lose, paper: win, scissors: draw},
	}
)

func Solve(input []string) int {
	total := 0

	for _, entry := range input {
		round := strings.Split(entry, " ")

		my := myShape[round[1]]
		ur := urShape[round[0]]

		score := scoreForShape[my]
		score += scoreForRound[my][ur]

		total += score
	}

	return total
}
