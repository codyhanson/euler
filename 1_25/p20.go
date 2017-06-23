package main

/*
Factorial digit sum
Problem 20
n! means n × (n - 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/
import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func factorial(n int64) *big.Int {
	f := big.NewInt(n)
	for i := n - 1; i > 0; i-- {
		bigI := big.NewInt(i)
		f.Mul(bigI, f)
	}
	return f
}

func main() {
	result := factorial(100)
	resultStr := result.String()
	resultChars := strings.Split(resultStr, "")
	sum := 0
	for c := range resultChars {
		val, _ := strconv.Atoi(resultChars[c])
		sum += val
	}
	fmt.Println(sum)
}
