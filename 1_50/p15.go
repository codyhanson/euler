package main

/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

https://projecteuler.net/project/images/p015.gif

How many such routes are there through a 20×20 grid?
*/

import (
	"fmt"
	"time"
)

var numPaths = 0
var numRows = 1
var numCols = 1

func atTheEnd(row, col int) bool {
	return row == numRows && col == numCols
}

func tryNextStep(row, col int) {
	if atTheEnd(row, col) {
		//this path is done
		numPaths++
		return
	}

	if row < numRows {
		tryNextStep(row+1, col)
	}
	if col < numCols {
		tryNextStep(row, col+1)
	}
}

func main() {
	for i := 0; i <= 20; i++ {
		numRows++
		numCols++
		numPaths = 0
		start := time.Now()
		tryNextStep(0, 0)
		elapsed := time.Since(start)
		fmt.Printf("%d x %d: %d, %v\n", numRows, numCols, numPaths, elapsed)
	}
}
