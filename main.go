package main

import (
	"fmt"
	"log"
)

func main() {

	// Loop through each struct within the AOC_Functions array
	for _, aoc := range AocFunctions {

		// Call the function pointer with the input file
		result, err := aoc.run(aoc.file)

		// Bail if we ever fail
		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		// Otherwise, print the result
		fmt.Println(result)
	}
}
