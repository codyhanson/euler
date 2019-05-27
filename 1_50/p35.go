package main

import (
	"fmt"
	"github.com/codyhanson/euler/eulertools"
	"github.com/codyhanson/euler/slicetools"
)

/*
Circular primes

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

func isCircular(primeCandidates <-chan int, circulars chan<- int) {
	for p := range primeCandidates {
		digits := slicetools.SplitDigits(p)
		circular := true
		for i := 0; i < len(digits); i++ {
			slicetools.RotateInt(digits, 1)
			if !eulertools.IsPrimeFaster(slicetools.CombineDigits(digits)) {
				circular = false
				break
			}
		}
		if circular {
			circulars <- p
		}
	}
	close(circulars)
}

func generateCandidates(primeCandidates chan<- int) {
	p := 2
	for p < 1000000 {
		primeCandidates <- p
		p = eulertools.GetNextPrimeFaster(p)
	}
	close(primeCandidates)
}

func main() {
	circularCount := 0
	primeCandidates := make(chan int, 1000000)
	circulars := make(chan int, 100)

	//workers
	go isCircular(primeCandidates, circulars)
	go generateCandidates(primeCandidates)

	for c := range circulars {
		fmt.Println("Found:", c)
		circularCount += 1
	}
	fmt.Println("Number of Circular Primes:", circularCount)
}
