package main

/*
Summation of primes
Problem 10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

import (
	"eulertools"
	"fmt"
)

func main() {
	sum := 0
	curPrime := 2
	for curPrime < 2000000 {
		sum += curPrime
		fmt.Println(curPrime)
		curPrime = eulertools.GetNextPrimeFaster(curPrime)
	}
	fmt.Println(sum)
}
