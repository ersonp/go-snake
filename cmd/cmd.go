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
	fmt.Printf("Please Enter the %q of the board.\n", measurement)
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		fmt.Scan(&s)
		i, err = strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Please enter a valid %q in int.\n", measurement)
		} else {
			break
		}
	}
	return i
}

// GetBoardSize is to get the width and height from the user
func GetBoardSize() (int, int) {
	fmt.Println("Go-Snake")
	height := readInt("height")
	width := readInt("width")
	return height, width
}
