package day6_test

import (
	"testing"

	"github.com/dosadczuk/aoc/2022/internal/io"
	"github.com/dosadczuk/aoc/2022/pkg/day6"
)

func TestSolve1(runner *testing.T) {
	tt := []struct {
		testname string
		filepath string
		expected int
	}{
		{
			testname: "Sample 1",
			filepath: "testdata/tc1.txt",
			expected: 7,
		},
		{
			testname: "Sample 2",
			filepath: "testdata/tc2.txt",
			expected: 5,
		},
		{
			testname: "Sample 3",
			filepath: "testdata/tc3.txt",
			expected: 6,
		},
		{
			testname: "Sample 4",
			filepath: "testdata/tc4.txt",
			expected: 10,
		},
		{
			testname: "Sample 5",
			filepath: "testdata/tc5.txt",
			expected: 11,
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

			if given := day6.Solve1(input); given != tc.expected {
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
			testname: "Sample 1",
			filepath: "testdata/tc1.txt",
			expected: 19,
		},
		{
			testname: "Sample 2",
			filepath: "testdata/tc2.txt",
			expected: 23,
		},
		{
			testname: "Sample 3",
			filepath: "testdata/tc3.txt",
			expected: 23,
		},
		{
			testname: "Sample 4",
			filepath: "testdata/tc4.txt",
			expected: 29,
		},
		{
			testname: "Sample 5",
			filepath: "testdata/tc5.txt",
			expected: 26,
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

			if given := day6.Solve2(input); given != tc.expected {
				t.Errorf("Given: %d, Expected: %d\n", given, tc.expected)
			}
		})
	}
}
