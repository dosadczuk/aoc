package day5

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dosadczuk/aoc/2022/internal/ds"
	"github.com/dosadczuk/aoc/2022/internal/math"
)

func Solve1(input []string) string {
	stacks := getStacks(&input)

	for _, entry := range input {
		qty, fromIdx, toIdx := parseAction(entry)

		from, to := stacks[fromIdx], stacks[toIdx]
		for ; qty > 0 && !from.Empty(); qty -= 1 {
			to.Push(from.Pop())
		}
	}

	return result(stacks)
}

func Solve2(input []string) string {
	stacks := getStacks(&input)

	for _, entry := range input {
		qty, fromIdx, toIdx := parseAction(entry)

		from, to := stacks[fromIdx], stacks[toIdx]
		temp := ds.NewStack[string]()

		for ; qty > 0 && !from.Empty(); qty -= 1 {
			temp.Push(from.Pop())
		}

		for !temp.Empty() {
			to.Push(temp.Pop())
		}
	}

	return result(stacks)
}

func getStacks(input *[]string) []*ds.Stack[string] {
	board := getCratesBoard(input)
	count := getCratesCount(&board)

	stacks := make([]*ds.Stack[string], count)
	for i, _ := range stacks {
		stacks[i] = ds.NewStack[string]()
	}

	for row := len(board) - 1; row >= 0; row-- {
		level := board[row]

		stkIdx := 0
		for col := 0; col < len(level); col, stkIdx = col+4, stkIdx+1 {
			crate := level[col:math.MinInt(col+4, len(level))]
			crate = regexp.MustCompile("\\w").FindString(crate)
			if crate == "" {
				continue
			}

			stacks[stkIdx].Push(crate)
		}
	}

	return stacks
}

func getCratesBoard(input *[]string) (board []string) {
	row := 0
	for ; row < len(*input); row++ {
		entry := (*input)[row]
		if entry == "" {
			break // end of the board
		}

		board = append(board, entry)
	}

	*input = (*input)[row+1:] // remove the board

	return
}

func getCratesCount(board *[]string) (count int) {
	n := len(*board)

	labels := regexp.MustCompile("\\d+").FindAllString((*board)[n-1], -1)
	count, _ = strconv.Atoi(labels[len(labels)-1])

	*board = (*board)[:n-1]

	return
}

func parseAction(action string) (qty, fromIdx, toIdx int) {
	parts := regexp.MustCompile("\\d+").FindAllString(action, -1)

	// quantity
	qty, _ = strconv.Atoi(parts[0])

	// "from stack" index
	fromIdx, _ = strconv.Atoi(parts[1])
	fromIdx -= 1 // index

	// "to stack" index
	toIdx, _ = strconv.Atoi(parts[2])
	toIdx -= 1 // index

	return
}

func result(stacks []*ds.Stack[string]) string {
	var out strings.Builder
	for _, s := range stacks {
		out.WriteString(s.Peek())
	}

	return out.String()
}
