package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arrNumber := []int{1, 4, 5, 6, 8, 2}

	// Vertical bar chart
	barChartVerticaly(arrNumber)
	fmt.Println()

	// Insertion sort ascending
	fmt.Println()
	insertionSortAsc(arrNumber)
	fmt.Println()

	// Insertion sort Desc
	fmt.Println()
	insertionSortDesc(arrNumber)

}

func barChartVerticaly(arrNumber []int) {
	// Find the maximum to determine the suitable grid to plot the Histogram
	max := math.MinInt64
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
}

func insertionSortAsc(arrNum []int) {
	for i := 0; i < len(arrNum); i++ {
		tmp := arrNum[i]
		j := i
		for j > 0 && arrNum[j-1] > tmp {
			arrNum[j] = arrNum[j-1]
			j--
		}
		arrNum[j] = tmp
	}
	barChartVerticaly(arrNum)
}

func insertionSortDesc(arrNumber []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(arrNumber)))
	barChartVerticaly(arrNumber)
}
