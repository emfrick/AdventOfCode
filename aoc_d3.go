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
	maxX, maxY := 0, 0
	for scanner.Scan() {

		// Grab the text
		line := scanner.Text()

		square := tokenize(line)

		x := square.x + square.width
		y := square.y + square.height

		if maxX <= x {
			maxX = x
		}

		if maxY <= y {
			maxY = y
		}
	}

	file.Close()

	// Open the file
	file, err = getFile(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	// Create the 2D array
	fabric := make([][]int, maxX+1)
	for i := 0; i <= maxX; i++ {
		fabric[i] = make([]int, maxY+1)
		for j := 0; j <= maxY; j++ {
			fabric[i][j] = 0
		}
	}

	// Get a scanner to read line-by-line
	scanner = bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {

		// Grab the text
		line := scanner.Text()

		square := tokenize(line)

		for i := 0; i < square.width; i++ {
			for j := 0; j < square.height; j++ {
				fabric[square.x+i][square.y+j]++
			}
		}
	}

	sum := 0
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if fabric[i][j] >= 2 {
				sum++
			}
		}
	}

	return fmt.Sprintf("Day 3, Problem 1 Solution: %d", sum), nil
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
