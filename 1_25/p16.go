package main

/*
Power digit sum
Problem 16
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	result := big.NewInt(2)
	oneK := big.NewInt(1000)
	result.Exp(result, oneK, nil)
	resultChars := strings.Split(result.String(), "")
	sum := 0
	for c := range resultChars {
		val, _ := strconv.Atoi(resultChars[c])
		sum += val
	}
	fmt.Println(sum)
}
