package main

/*
Special Pythagorean triplet
Problem 9
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

import "fmt"

func isPythagoreanTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

/*
a = 1
b,c 998, 1
	997, 2
*/

func main() {

	for a := 1; a < 998; a++ {
		for b := 999 - a; b > a; b-- {
			c := 1000 - a - b
			fmt.Printf("%d %d %d %d\n", a, b, c, a+b+c)
			if a+b+c == 1000 && isPythagoreanTriplet(a, b, c) {
				//got it.
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
