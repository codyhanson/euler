package main

/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/
import (
	"fmt"
	"strconv"
	"strings"
)

func factorial(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

func sumOfDigits(i int) int {
	sum := 0

	return sum
}

func main() {

	runningTotal := 0
	n := 3
	for {
		factorialSum := 0
		nStr := strconv.Itoa(n)
		digits := strings.Split(nStr, "")
		for _, val := range digits {
			digit, _ := strconv.Atoi(val)
			factorialSum += factorial(digit)
		}

		if n == factorialSum {
			runningTotal += n
			fmt.Printf("%d: %d\n", n, runningTotal)
		}
		n++
	}

}
