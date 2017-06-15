package main

/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

import "fmt"
import "eulertools"

// Thought: Could probably memoize getNextPrime, to avoide recomputation.

func main() {
	n := 600851475143
	//n := 13195
	var factors []int
	nextCandidateFactor := 2
	for !eulertools.IsPrime(n) {
		for {
			if n%nextCandidateFactor == 0 {
				//we found a prime factor.
				fmt.Println(nextCandidateFactor)
				factors = append(factors, nextCandidateFactor)
				n = n / nextCandidateFactor
				break
			} else {
				//this prime is not the next factor
				nextCandidateFactor = eulertools.GetNextPrime(nextCandidateFactor)
			}
		}
	}
	//last factor is now n
	fmt.Println(n)
}
