package cmd

import (
	"fmt"
	"strconv"
)

//readInt reads from the cmd and validates if its a int
func readInt(measurement string) int {
	var s string
	var i int
	var err error
	fmt.Printf("Please Enter the %q of the board between 10-80.\n", measurement)
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		fmt.Scan(&s)
		i, err = strconv.Atoi(s)
		if err != nil || i < 10 || i > 80 {
			fmt.Printf("Please enter a valid %q in int between 10-80.\n", measurement)
		} else {
			break
		}
	}
	return i
}

// GetBoardSize is to get the width and height from the user
func GetBoardSize() (int, int) {
	height := readInt("height")
	width := readInt("width")
	return height, width
}
