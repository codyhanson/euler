package main

import "fmt"

/*
If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

func main() {

	max := 0
	maxp := 0
	// bruuuuute forrrrrrrce
	for perimeter := 3; perimeter < 1000; perimeter++ {
		solutions := 0
		a := perimeter - 2
		b := 1
		c := 1

		for a > 1 {
			for b > 1 {
				for c <= perimeter-b-a {
					//fmt.Printf("%d %d %d\n", a, b, c)
					if validRightTriangle(a, b, c) {
						fmt.Printf("Found solution for %d {%d, %d, %d}\n", perimeter, a, b, c)
						solutions++
					}
					c++
				}
				b--
			}
			a--
			b = perimeter - a - 1
			c = 1
		}

		solutions /= 2

		if solutions > max {
			fmt.Printf("New Max found for p %d\n", perimeter)
			max = solutions
			maxp = perimeter
		}
	}

	fmt.Println(maxp)
}

func validRightTriangle(a, b, c int) bool {
	// Thanks, Pythagoras
	return a*a+b*b == c*c
}
