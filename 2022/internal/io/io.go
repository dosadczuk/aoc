package io

import (
	"bufio"
	"os"
)

// ReadLines reads lines from file located in given path.
func ReadLines(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines, s.Err()
}
