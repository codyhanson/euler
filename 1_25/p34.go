package main

/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/
import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	f := big.NewInt(n)
	for i := n - 1; i > 0; i-- {
		bigI := big.NewInt(i)
		f.Mul(bigI, f)
	}
	return f
}

func sumOfDigits(i int) {

}

func main() {

	for {

	}

}
