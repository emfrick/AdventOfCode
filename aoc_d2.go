package main

import (
	"bufio"
	"fmt"
	"strings"
)

// AocD2P1 conforms to the AdventOfCodeInterface
func AocD2P1(filename string) (string, error) {

	// Open the file
	file, err := getFile(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	// Get a scanner to read line-by-line
	scanner := bufio.NewScanner(file)

	// 2's and 3's both set to 0
	duplicates := [2]int{0, 0}

	// Read each line
	for scanner.Scan() {

		// Grab the text
		line := scanner.Text()
		letters := strings.Split(line, "")

		var twos, threes bool
		letterMap := make(map[string]int)
		for _, letter := range letters {
			letterMap[letter]++
		}

		for _, value := range letterMap {
			if value == 2 && !twos {
				duplicates[0]++
				twos = true
			}

			if value == 3 && !threes {
				duplicates[1]++
				threes = true
			}
		}
	}

	return fmt.Sprintf("Day 2, Problem 1 Solution: %d", duplicates[0]*duplicates[1]), nil
}

// AocD2P2 requires research into
// Levenshtein Distance, because this solution is garbage
func AocD2P2(filename string) (string, error) {
	// Open the file
	file, err := getFile(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	// Get a scanner to read line-by-line
	scanner := bufio.NewScanner(file)

	var words []string
	var solution string
	// Read each line
	for scanner.Scan() {

		// Grab the text
		line := scanner.Text()

		words = append(words, line)
	}

	for i := 0; i < len(words)-1; i++ {

		a := words[i]

		for j := (i + 1); j < len(words); j++ {

			b := words[j]

			aLetters := strings.Split(a, "")
			bLetters := strings.Split(b, "")

			differs := 0
			sameLetters := []string{}
			for i, letter := range aLetters {
				if letter != bLetters[i] {
					differs++
				} else {
					sameLetters = append(sameLetters, letter)
				}
			}

			if differs == 1 {
				solution = strings.Join(sameLetters, "")
			}

		}

	}

	return fmt.Sprintf("Day 2, Problem 2 Solution: %s", solution), nil
}
