package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for input != "EXIT" {
		fmt.Print("What integer would you like to convert to a hex value? (Type 'EXIT' to leave)")
		scanner.Scan()
		input = scanner.Text()

		if input != "EXIT" {
			integer, error := strconv.Atoi(input)
			if error != nil {
				fmt.Print("Please enter an integer...\n")
			} else {
				hex := fmt.Sprintf("%x", integer)
				fmt.Printf("Hex conversion of '%d' is '%s'\n", integer, hex)
			}
		}
	}
}
