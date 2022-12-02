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
	myResult = map[string]int{"X": lose, "Y": draw, "Z": win}

	// shapeForResult is the mapping for the result and the shape.
	//
	// It answers the question:
	// "I expect the result and my opponent's choice is the shape,
	// what shape do I need to choose?"
	shapeForResult = map[int]map[shape]shape{
		lose: {rock: scissors, paper: rock, scissors: paper},
		draw: {rock: rock, paper: paper, scissors: scissors},
		win:  {rock: paper, paper: scissors, scissors: rock},
	}
)

var (
	scoreForShape = map[shape]int{rock: 1, paper: 2, scissors: 3}

	// scoreForResult is the mapping for the shape and the shape.
	//
	// It answers the question:
	// "What's the result of the battle between me, with the shape,
	// and my opponent with the shape."
	scoreForResult = map[shape]map[shape]int{
		rock:     {rock: draw, paper: lose, scissors: win},
		paper:    {rock: win, paper: draw, scissors: lose},
		scissors: {rock: lose, paper: win, scissors: draw},
	}
)

func Solve1(input []string) int {
	total := 0

	for _, entry := range input {
		round := strings.Split(entry, " ")

		my := myShape[round[1]]
		ur := urShape[round[0]]

		score := scoreForShape[my]
		score += scoreForResult[my][ur]

		total += score
	}

	return total
}

func Solve2(input []string) int {
	total := 0

	for _, entry := range input {
		round := strings.Split(entry, " ")

		res := myResult[round[1]]

		ur := urShape[round[0]]
		my := shapeForResult[res][ur]

		score := res
		score += scoreForShape[my]

		total += score
	}

	return total
}
