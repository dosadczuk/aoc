package day1_test

import (
	"testing"

	"github.com/dosadczuk/aoc/2022/internal/day1"
	"github.com/dosadczuk/aoc/2022/pkg/io"
)

func TestMain(runner *testing.T) {
	tt := []struct {
		testname string
		filepath string
		expected int
	}{
		{
			testname: "Person with one item",
			filepath: "testdata/tc1_person_item.txt",
			expected: 1000,
		},
		{
			testname: "Person with many items",
			filepath: "testdata/tc2_person_items.txt",
			expected: 3000,
		},
		{
			testname: "People with one item each",
			filepath: "testdata/tc3_people_item.txt",
			expected: 2000,
		},
		{
			testname: "People with many items each",
			filepath: "testdata/tc4_people_items.txt",
			expected: 5000,
		},
	}

	for _, tc := range tt {
		runner.Run(tc.testname, func(t *testing.T) {
			input, err := io.ReadLines(tc.filepath)
			if err != nil {
				t.Error(err)
				return
			}

			if given := day1.Solve(input); given != tc.expected {
				t.Errorf("Given: %d, Expected: %d\n", given, tc.expected)
			}
		})
	}
}
