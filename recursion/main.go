package main

//Find all unique path from top left box to bottom right box 
// only down or right moves are allowed

import (
	"fmt"
)
func main() {
	fmt.Println("Starting main...")
	m := 3
	n := 4

	res := findAllPaths(m-1, n-1)
	fmt.Println("Result is: ", res)
}

func findAllPaths(row, col int) int {
	//base case
	if row == 0 || col == 0 {
		return 1
	}
	fmt.Println("Calling with row: ", row, "and col: ", col)
	return findAllPaths(row, col-1) + findAllPaths(row-1, col)

	//return 100
}

func init() {
	fmt.Println("setting up some state...")
}