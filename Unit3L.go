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
	var colour string = ""

	var reader = bufio.NewReader(os.Stdin)

	// input colour
	fmt.Print("Please choose a sweater colour from the available choices: Blue, Black, Red, White.\n")
	colour, _ = reader.ReadString('\n')
	colour = strings.TrimSpace(colour)

	// check colour choice
	if strings.ToLower(colour) == "black" || strings.ToLower(colour) == "white" {
		fmt.Println("You have enough sweaters in this colour.")
	} else if strings.ToLower(colour) == "red" {
		fmt.Println("This colour will look good on you.")
		fmt.Println("How about a pair of jeans to go with the sweater?")
	} else if strings.ToLower(colour) == "blue" {
		fmt.Println("This colour doesn’t go well with your eyes.")
	} else {
		fmt.Println("Your colour choice is invalid.")
	}

	fmt.Println("\nDone.")
}
