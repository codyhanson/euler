package main

/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

import "fmt"
import "eulertools"

var factors = [...]int{}

func main() {
	n := 600851475143
	//for target % nextPrime == 0
	nextCandidateFactor := 2
	for {

		//this prime is not a factor
		nextCandidateFactor = eulertools.GetNextPrime(nextCandidateFactor)
		fmt.Println(i)
	}
	fmt.Println(factors[len(factors)-1])
}
