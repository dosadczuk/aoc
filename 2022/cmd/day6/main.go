package main

import (
	"fmt"
	"log"

	"github.com/dosadczuk/aoc/2022/internal/io"
	"github.com/dosadczuk/aoc/2022/pkg/day6"
)

func main() {
	input, err := io.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result (1): %d\n", day6.Solve1(input))
	fmt.Printf("Result (2): %d\n", day6.Solve2(input))
}
