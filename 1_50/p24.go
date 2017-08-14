package main

/*Lexicographic permutations
Problem 24
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

import (
	"bytes"
	"fmt"
	"sort"
)

func concat(digits []string) string {
	var currentP bytes.Buffer
	for j := 0; j < len(digits); j++ {
		currentP.WriteString(digits[j])
	}
	return currentP.String()
}

var permutations = make([]string, 0, 4000000)

func gather(s []string) {
	p := concat(s)
	permutations = append(permutations, p)
}

//https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permute(n int, s []string) {
	if n == 1 {
		gather(s)
		return
	}
	for i := 0; i < n; i++ {
		permute(n-1, s)
		if n%2 == 0 {
			s[0], s[n-1] = s[n-1], s[0]
		} else {
			s[i], s[n-1] = s[n-1], s[i]
		}
	}
}

func main() {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	permute(10, digits)
	sort.Strings(permutations)
	fmt.Println("-------")
	fmt.Println(permutations[:12])
	fmt.Printf("%q\n", permutations[0])
	fmt.Println(permutations[999999])
	fmt.Println(permutations[1000000])
}
