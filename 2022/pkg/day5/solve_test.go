package day5_test

import (
	"testing"

	"github.com/dosadczuk/aoc/2022/internal/io"
	"github.com/dosadczuk/aoc/2022/pkg/day5"
)

func TestSolve1(runner *testing.T) {
	tt := []struct {
		testname string
		filepath string
		expected string
	}{
		{
			testname: "Sample from website",
			filepath: "testdata/tc1_sample.txt",
			expected: "CMZ",
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

			if given := day5.Solve1(input); given != tc.expected {
				t.Errorf("Given: %s, Expected: %s\n", given, tc.expected)
			}
		})
	}
}

func TestSolve2(runner *testing.T) {
	runner.Skip("TODO")

	tt := []struct {
		testname string
		filepath string
		expected string
	}{
		{
			testname: "Sample from website",
			filepath: "testdata/tc1_sample.txt",
			expected: "MCD",
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

			if given := day5.Solve2(input); given != tc.expected {
				t.Errorf("Given: %s, Expected: %s\n", given, tc.expected)
			}
		})
	}
}
