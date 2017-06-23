package main

/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

https://projecteuler.net/project/images/p015.gif

How many such routes are there through a 20×20 grid?
*/

import (
	"fmt"
)

var numPaths = 0
var numRows = 20
var numCols = 20

func atTheEnd(row, col int) bool {
	return row == numRows && col == numCols
}

func tryNextStep(row, col int, downOrRight string) {

	//advance a step
	if downOrRight == "d" {
		row++
	} else {
		col++
	}

	if atTheEnd(row, col) {
		//this path is done
		numPaths++
		return
	}
	if downOrRight == "d" {
		if row == numRows && col < numCols {
			//cant go down more. just try right.
			tryNextStep(row, col, "r")
		} else {
			//try going down one more
			tryNextStep(row, col, "d")
			//also try going right, from here
			if col < numCols {
				tryNextStep(row, col, "r")
			}
		}

	} else {
		//going right
		if col == numCols && row < numRows {
			//cant go right more, just try down
			tryNextStep(row, col, "d")
		} else {
			//go right one more
			tryNextStep(row, col, "r")
			//also try going down from here
			if row < numRows {
				tryNextStep(row, col, "d")
			}
		}
	}
}

func main() {
	tryNextStep(0, 0, "d")
	//tryNextStep(0, 0, "r")
	fmt.Println(numPaths * 2)
}
