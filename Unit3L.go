// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-15
// Fileoverview: This program chooses a sweater colour and prints a message based on the choice.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// variable declaration
	var colour string

	reader := bufio.NewReader(os.Stdin)

	// input colour
	fmt.Print("Please choose a sweater colour from the available choices: Blue, Black, Red, White.\n")
	colour, _ = reader.ReadString('\n')
	colour = strings.TrimSpace(colour)

	// check colour choice
	switch strings.ToLower(colour) {
	case "black", "white":
		fmt.Println("You have enough sweaters in this colour.")
	case "red":
		fmt.Println("This colour will look good on you.")
		fmt.Println("How about a pair of jeans to go with the sweater?")
	case "blue":
		fmt.Println("This colour doesn’t go well with your eyes.")
	default:
		fmt.Println("Your colour choice is invalid.")
	}

	fmt.Println("\nDone.")
}
