package day5

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dosadczuk/aoc/2022/internal/ds"
	"github.com/dosadczuk/aoc/2022/internal/math"
)

func Solve1(input []string) string {
	stacks := parseStacks(&input)

	for _, entry := range input {
		qty, srcIdx, tgtIdx := parseAction(entry)

		src, tgt := stacks[srcIdx], stacks[tgtIdx]
		for ; qty > 0 && !src.Empty(); qty -= 1 {
			tgt.Push(src.Pop())
		}
	}

	return result(stacks)
}

func Solve2(input []string) string {
	stacks := parseStacks(&input)

	for _, entry := range input {
		qty, srcIdx, tgtIdx := parseAction(entry)

		src, tgt := stacks[srcIdx], stacks[tgtIdx]
		tmp := ds.NewStackWithSizeOf[string](qty)

		for ; qty > 0 && !src.Empty(); qty -= 1 {
			tmp.Push(src.Pop())
		}

		for !tmp.Empty() {
			tgt.Push(tmp.Pop())
		}
	}

	return result(stacks)
}

// parseStacks retrieves creates board from the input and converts it
// into slice of Stack data structure.
//
// Example:
//
//	Input:
//		    [D]
//		[N] [C]
//		[Z] [M] [P]
//		1   2   3
//
//		...
//
//	Output:
//		[
//		  Stack{ Z, N },
//		  Stack{ M, C, D },
//		  Stack{ P },
//		]
func parseStacks(input *[]string) []*ds.Stack[string] {
	board := parseBoard(input)

	stacks := make([]*ds.Stack[string], len(board[0])/4+1)
	for i := range stacks {
		stacks[i] = ds.NewStack[string]()
	}

	for row := len(board) - 2; row >= 0; row-- {
		cols := board[row]
		size := len(cols)

		for col, stk := 0, 0; col < size; col, stk = col+4, stk+1 {
			end := math.MinInt(col+4, size)

			// e.g. "[A] " -> "A"
			crate := regexp.MustCompile("\\w").FindString(cols[col:end])
			if crate == "" {
				continue
			}

			stacks[stk].Push(crate)
		}
	}

	return stacks
}

// parseBoard retrieves crates board from input lines.
//
// Example:
//
//	Input:
//		    [D]
//		[N] [C]
//		[Z] [M] [P]
//		1   2   3
//
//		move 1 from 2 to 1
//		move 3 from 1 to 3
//		move 2 from 2 to 1
//		move 1 from 1 to 2
//
//	Output:
//
//		    [D]
//		[N] [C]
//		[Z] [M] [P]
//		1   2   3
func parseBoard(input *[]string) (board []string) {
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

// parseAction retrieves most important values from the line:
// 1. a quantity of crates is moved from one stack to a different stack,
// 2. from which stack,
// 3. to which stack.
//
// Example:
//
//	Input:
//		move 1 from 2 to 1
//
//	Output:
//		qty: 1
//		src: 1 (-1 because it's index)
//		tgt: 0 (-1 because it's index)
func parseAction(action string) (qty, src, tgt int) {
	parts := regexp.MustCompile("\\d+").FindAllString(action, -1)

	// quantity
	qty, _ = strconv.Atoi(parts[0])

	// "source stack" index
	src, _ = strconv.Atoi(parts[1])
	src -= 1 // index

	// "target stack" index
	tgt, _ = strconv.Atoi(parts[2])
	tgt -= 1 // index

	return
}

// result returns the result expected by the input.
// It peeks and concatenates top element from each stack.
func result(stacks []*ds.Stack[string]) string {
	var out strings.Builder
	for _, s := range stacks {
		out.WriteString(s.Peek())
	}

	return out.String()
}
