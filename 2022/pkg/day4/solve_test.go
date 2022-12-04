package day4_test

import (
	"testing"

	"github.com/dosadczuk/aoc/2022/internal/io"
	"github.com/dosadczuk/aoc/2022/pkg/day4"
)

func TestSolve1(runner *testing.T) {
	tt := []struct {
		testname string
		filepath string
		expected int
	}{
		{
			testname: "Sample from website",
			filepath: "testdata/tc1_sample.txt",
			expected: 2,
		},
	}

	for _, tc := range tt {
		tc := tc

		runner.Run(tc.testname, func(t *testing.T) {
			t.Parallel()

			input, err := io.ReadLines(tc.filepath)
			if err != nil {
				t.Error(err)
				return
			}

			if given := day4.Solve1(input); given != tc.expected {
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
			testname: "Sample from website",
			filepath: "testdata/tc1_sample.txt",
			expected: 4,
		},
	}

	for _, tc := range tt {
		tc := tc

		runner.Run(tc.testname, func(t *testing.T) {
			t.Parallel()

			input, err := io.ReadLines(tc.filepath)
			if err != nil {
				t.Error(err)
				return
			}

			if given := day4.Solve2(input); given != tc.expected {
				t.Errorf("Given: %d, Expected: %d\n", given, tc.expected)
			}
		})
	}
}
