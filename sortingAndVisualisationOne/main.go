package main

import (
	"fmt"
	"math"
)

func main() {
	var arrNumber = []int{1, 4, 5, 6, 8, 2}
	max := math.MinInt64

	// Find the maximum to determine the suitable grid to plot the Histogram
	for _, e := range arrNumber {
		if e > max {
			max = e
		}
	}
	var (
		row = max
		col = len(arrNumber)
	)

	// Create grind of dimension (row * colmn)
	// Allocate rows
	histogram := make([][]string, row)
	for i := 0; i < row; i++ {
		// Allocate column for each row
		histogram[i] = make([]string, col)
	}

	// Histogram formation logic
	// use blank spaces ("	") or "  |  " to fill the grid
	for i, e := range arrNumber {
		for j := row - 1; j >= 0; j-- {
			if j >= row-e {
				histogram[j][i] = " | "
			} else {
				histogram[j][i] = "   "
			}
		}
	}

	// Print the histogram
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%s", histogram[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < col; i++ {
		fmt.Printf(" %d ", arrNumber[i])
	}
	fmt.Println()
}
