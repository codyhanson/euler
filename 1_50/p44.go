package main

/*
Pentagon numbers
Problem 44
Pentagonal numbers are generated by the formula, Pn=n(3n-1)/2. The first ten pentagonal numbers are:

1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70 - 22 = 48, is not pentagonal.

Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and D = |Pk - Pj| is minimised; what is the value of D?
*/

import (
	"fmt"
	"math"
)

func makePentagon(n int) int {
	return n * (3*n - 1) / 2
}

func isPentagon(p int) bool {
	//a := math.Sqrt((float64(p*2 + )))
	a := p*2 +
		fmt.Println(a)
	return a != 0
}

var pentagons []int

var minDValue = 999999999

func checkD() {

}

func main() {
	fmt.Println(isPentagon(22))
	fmt.Println(isPentagon(23))
	fmt.Println(isPentagon(145))
	numPentagons := 1000000
	fmt.Printf("Trying %d Pentagonal numbers\n", numPentagons)
	for n := 1; n <= numPentagons; n++ {
		pentagons = append(pentagons, makePentagon(n))
	}
	for i := 0; i < len(pentagons); i++ {
		for j := i + 1; j < len(pentagons); j++ {
			//			if (isPentagon
			return
		}
	}

}
