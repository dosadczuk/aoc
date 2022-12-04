package main

import (
	"fmt"
	"log"

	"github.com/dosadczuk/aoc/2022/internal/io"
	"github.com/dosadczuk/aoc/2022/pkg/day4"
)

func main() {
	input, err := io.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result (1): %d\n", day4.Solve1(input))
	//fmt.Printf("Result (2): %d\n", day4.Solve2(input))
}
