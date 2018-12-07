package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// AocD1P1 conforms to the AdventOfCodeInterface
func AocD1P1(filename string) (string, error) {

	// Open the file
	file, err := getFile(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	// Get a scanner to read line-by-line
	scanner := bufio.NewScanner(file)

	// Read each line
	var sum int = 0
	for scanner.Scan() {

		// Grab the text and convert to int
		line := scanner.Text()
		num, err := strconv.Atoi(line)

		if err != nil {
			return "", fmt.Errorf("Could not parse int '%s'", line)
		}

		// Add the number to the running total
		sum += num
	}

	// Return the running total
	return fmt.Sprintf("Day 1, Problem 1 Solution: %d", sum), nil
}

// AocD1P2 conforms to the AdventOfCodeInterface
func AocD1P2(filename string) (string, error) {

	frequency := 0
	frequencyMap := make(map[int]int)
	frequencyMap[frequency]++

	for {

		// Open the file
		file, err := getFile(filename)

		if err != nil {
			return "", err
		}

		defer file.Close()

		// Get a scanner to read line-by-line
		scanner := bufio.NewScanner(file)

		// For each line
		for scanner.Scan() {

			// Grab the text and convert it to a number
			line := scanner.Text()
			num, err := strconv.Atoi(line)

			if err != nil {
				return "", fmt.Errorf("Could not parse int '%s'", line)
			}

			// Add the number to our running frequency
			// and increment the number of occurrances of that frequency
			frequency += num
			frequencyMap[frequency]++

			// If we have more than one particular frequency, we've found our solution,
			// so return that with nil error
			if frequencyMap[frequency] > 1 {
				return fmt.Sprintf("Day 1, Problem 2 Solution: %d", frequency), nil
			}
		}

		// If we've reached this point, we have not found a repeating frequency,
		// So close the file and start over (repeating frequencies can be found
		// after several loops through the file)
		file.Close()
	}
}

func getFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("Unable to open file '%s'", filename)
	}

	return file, nil
}
