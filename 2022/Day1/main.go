package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readInputFile("submit.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Calories: %d\n", mostCaloriesTheElfIsCarrying(input))
}

func mostCaloriesTheElfIsCarrying(input []string) int {
	sumRunning := 0
	sumHighest := 0

	for _, entry := range input {
		if entry == "" {
			if sumRunning > sumHighest {
				sumHighest = sumRunning
			}

			sumRunning = 0
			continue
		}

		calories, _ := strconv.Atoi(entry)
		sumRunning += calories
	}

	if sumRunning > sumHighest {
		return sumRunning
	} else {
		return sumHighest
	}
}

func readInputFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil
}
