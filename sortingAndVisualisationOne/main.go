package main

import (
	"fmt"
	"math"
)

func main() {
	var arrNumber = []int{1, 4, 5, 6, 8, 2}
	max := math.MinInt64
	barChartVerticaly(arrNumber)
	fmt.Println()

	// Insertion sort ascending
	fmt.Println()
	insertionSort(arrNumber)

	// Find the maximum to determine the suitable grid to plot the Histogram
	for _, e := range arrNumber {
		if e > max {
			max = e
		}
	}
	var (
		rows  = max
		colmn = len(arrNumber)
	)

	// Create grind of dimension (row * colmn)
	// Allocate rows
	histograms := make([][]string, rows)
	for i := 0; i < rows; i++ {
		// Allocate column for each row
		histograms[i] = make([]string, colmn)
	}

	// Histogram formation logic
	// use blank spaces ("	") or "  |  " to fill the grid
	for i, e := range arrNumber {
		for j := rows - 1; j >= 0; j-- {
			if j >= rows-e {
				histograms[j][i] = " | "
			} else {
				histograms[j][i] = "   "
			}
		}
	}

	// Print the histogram
	for i := 0; i < rows; i++ {
		for j := 0; j < colmn; j++ {
			fmt.Printf("%s", histograms[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < colmn; i++ {
		fmt.Printf(" %d ", arrNumber[i])
	}

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

func insertionSort(arrNum []int) {
	for i := 0; i < len(arrNum); i++ {
		tmp := arrNum[i]
		j := i
		for j > 0 && arrNum[j-1] > tmp {
			arrNum[j] = arrNum[j-1]
			j--
		}
		arrNum[j] = tmp
	}
}
