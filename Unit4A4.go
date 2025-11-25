// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-24
// Fileoverview: This program prompts the user for a starting and ending ASCII value (between 32 and 126)
// and displays each number with its corresponding character.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var startValue int
	var endValue int
	var input string

	reader := bufio.NewReader(os.Stdin)

	// get starting value
	fmt.Print("Please enter a number larger than 32, and less than 126: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	startValue, _ = strconv.Atoi(input)

	// get ending value
	for {
		fmt.Printf("Please enter a number larger than %d and less than 126: ", startValue)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		endValue, _ = strconv.Atoi(input)
		if endValue > startValue && endValue < 127 {
			break
		}
		fmt.Println("Invalid ending value. It must be larger than start value and less than 126.")
	}

	// display ASCII values and characters
	fmt.Println("\nASCII values and their characters:")
	for i := startValue; i <= endValue; i++ {
		fmt.Printf("%d = %c\n", i, rune(i))
	}

	fmt.Println("\nDone.")
}
