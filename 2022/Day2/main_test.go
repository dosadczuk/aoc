package main

import "testing"

func TestMain(runner *testing.T) {
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
			input, err := readInputFile(tc.filepath)
			if err != nil {
				t.Error(err)
				return
			}

			given := totalScoreForStrategy(input)
			if given != tc.expected {
				t.Errorf("\nScore: %d\nExpected: %d\n\n", given, tc.expected)
				return
			}
		})
	}

}
