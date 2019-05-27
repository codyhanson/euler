package main

import (
	"fmt"
	"github.com/codyhanson/euler/eulertools"
	"github.com/codyhanson/euler/slicetools"
)

/*
Truncatable Primes

The number 3797 has an interesting property. Being prime itself,
it is possible to continuously remove digits from left to right, and remain prime at each stage:
3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

func isTruncatable(digits []int, goRight bool) bool {
	n := slicetools.CombineDigits(digits)
	if !eulertools.IsPrimeFaster(n) {
		return false
	} else if len(digits) == 1 {
		return true
	}

	if goRight {
		return isTruncatable(digits[1:], goRight)
	} else {
		return isTruncatable(digits[:len(digits)-1], goRight)
	}
}

func pointersAtMax(pointers []int) bool {
	for _, p := range pointers {
		if p < 4 {
			return false
		}
	}
	return true
}

func generatePrimes(primesToCheck chan<- int) {
	//Primes less than 100 are trickier, so explicitly generate
	p := 11
	for p < 100 {
		p =  eulertools.GetNextPrime(p)
		primesToCheck <- p
	}
	//Above 3 digits we can stick to 1/3/5/7/9
	length := 3
	primeDigits := [5]int{1, 3, 5, 7, 9}
	for {
		pointers := make([]int, length)
		for !pointersAtMax(pointers) {
			digits := make([]int, length)
			for i, digitPointer := range pointers {
				digits[i] = primeDigits[digitPointer]
			}
			primesToCheck <- slicetools.CombineDigits(digits)
			// bump pointers by one.
			for i := length - 1; i >= 0; i-- {
				if pointers[i] < 4 {
					pointers[i] += 1
					if (length > 2 && pointers[i] == 2) {
						//Skip 5's after length > 2
						pointers[i] += 1
					}

					break
				} else {
					pointers[i] = 0
				}
			}
		}
		length += 1
	}
}

func checkPrimes(workerNum int, foundPrimes chan<- int, primesToCheck <-chan int) {
	for nextCandidate := range primesToCheck {
		//fmt.Println("Testing:", workerNum, nextCandidate)
		digits := slicetools.SplitDigits(nextCandidate)
		if isTruncatable(digits, true) && isTruncatable(digits, false) {
			foundPrimes <- nextCandidate
		}
	}

}

func collectResults(done chan<- bool, foundPrimes <-chan int) {
	truncatables := make([]int, 0, 11)

	for nextCandidate := range foundPrimes {
		fmt.Println("Found:", nextCandidate)
		truncatables = append(truncatables, nextCandidate)
		if len(truncatables) == 11 {
			break
		}
		//close channels?
	}

	sum := 0
	for i := 0; i < len(truncatables); i++ {
		sum = sum + truncatables[i]
	}
	fmt.Println("sum:", sum)
	done <- true
}

func main() {
	primesToCheck := make(chan int, 10000000)
	foundPrimes := make(chan int, 8)
	done := make(chan bool)
	go generatePrimes(primesToCheck)
	go collectResults(done, foundPrimes)
	numWorkers := 8
	for i := 0; i < numWorkers; i++ {
		go checkPrimes(i, foundPrimes, primesToCheck)
	}
	<-done //wait for the finish
}
