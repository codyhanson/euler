package main

/*
10001st prime
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

import (
	"eulertools"
	"fmt"
)

func main() {
	curPrime := 1
	for i := 0; i < 10001; i++ {
		//hell yeah code reuse
		curPrime = eulertools.GetNextPrime(curPrime)
	}
	fmt.Println(curPrime)
}
