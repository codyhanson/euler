package main

/*
Digit fifth powers
Problem 30
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	//"sync"
)

func splitDigits(n int) []int {
	nStr := strconv.Itoa(n)
	digits := strings.Split(nStr, "")
	intDigits := make([]int, len(digits))
	for _, d := range digits {
		digit, _ := strconv.Atoi(d)
		intDigits = append(intDigits, digit)
	}
	return intDigits
}

var gSum = 0

func checkFifthPower(n int) {
	digits := splitDigits(n)
	sum := 0.0
	for _, d := range digits {
		sum += math.Pow(float64(d), 5)
	}
	if n == int(sum) {
		gSum += n
		fmt.Println(n, gSum)
	}
}

func main() {

	n := 11
	for {
		checkFifthPower(n)
		n++
	}

}
