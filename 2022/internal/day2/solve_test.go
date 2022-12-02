package day2_test

import (
	"testing"

	"github.com/dosadczuk/aoc/2022/internal/day2"
	"github.com/dosadczuk/aoc/2022/pkg/io"
)

func TestSolve1(runner *testing.T) {
	tt := []struct {
		testname string
		filepath string
		expected int
	}{
		{
			testname: "Sample strategy from website",
			filepath: "testdata/tc1_sample.txt",
			expected: 15,
		},
	}

	for _, tc := range tt {
		runner.Run(tc.testname, func(t *testing.T) {
			input, err := io.ReadLines(tc.filepath)
			if err != nil {
				t.Error(err)
				return
			}

			if given := day2.Solve1(input); given != tc.expected {
				t.Errorf("Given: %d, Expected: %d\n", given, tc.expected)
			}
		})
	}
}

func TestSolve2(runner *testing.T) {
	tt := []struct {
		testname string
		filepath string
		expected int
	}{
		{
			testname: "Sample strategy from website",
			filepath: "testdata/tc1_sample.txt",
			expected: 12,
		},
	}

	for _, tc := range tt {
		runner.Run(tc.testname, func(t *testing.T) {
			input, err := io.ReadLines(tc.filepath)
			if err != nil {
				t.Error(err)
				return
			}

			if given := day2.Solve2(input); given != tc.expected {
				t.Errorf("Given: %d, Expected: %d\n", given, tc.expected)
			}
		})
	}
}
