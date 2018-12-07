package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Square struct {
	x      int
	y      int
	width  int
	height int
}

// AocD3P1 conforms to the AdventOfCodeInterface
func AocD3P1(filename string) (string, error) {

	// Open the file
	file, err := getFile(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	// Get a scanner to read line-by-line
	scanner := bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {

		// Grab the text
		line := scanner.Text()

		square := tokenize(line)

		fmt.Println(square)
	}

	return fmt.Sprintf("Day 3, Problem 1 Solution: "), nil
}

// AocD3P2 conforms to the AdventOfCodeInterface
func AocD3P2(filename string) (string, error) {

	return fmt.Sprintf("Day 3, Problem 2 Solution: "), nil
}

func tokenize(line string) Square {

	tokens := strings.Split(line, " ")

	tmp := strings.TrimRight(tokens[2], ":")
	coords := strings.Split(tmp, ",")

	x, err := strconv.Atoi(coords[0])
	y, err := strconv.Atoi(coords[1])

	area := strings.Split(tokens[3], "x")
	w, err := strconv.Atoi(area[0])
	h, err := strconv.Atoi(area[1])

	if err != nil {
		fmt.Println("Error conversion")
	}

	return Square{x, y, w, h}
}
