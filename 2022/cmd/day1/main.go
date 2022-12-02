package main

import (
	"fmt"
	"log"

	"github.com/dosadczuk/aoc/2022/internal/day1"
	"github.com/dosadczuk/aoc/2022/pkg/io"
)

func main() {
	input, err := io.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", day1.Solve(input))
}
