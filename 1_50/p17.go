package main

/*
Number letter counts
Problem 17
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/
import (
	"bytes"
	"fmt"
)

var ones = map[int]string{
	0: "",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teens = map[int]string{
	0:  "",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens = map[int]string{
	0:  "",
	10: "ten",
	20: "twenty",
	30: "thirty",
	40: "fourty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

var oneThousand = "onethousand"
var and = "and"

func splitIntoSpeakables(i int) (hundreds, tens, teens, ones int) {
	if i < 10 {
		return i
	} else if i < 100 {

	} else if i < 1000 {

	}

}

func speaknspell(i int) string {
	thousands, hundreds, tens, teens, ones := splitIntoSpeakables(i)
	if thousands {
		return oneThousand
	}
	var buffer bytes.Buffer

	if hundreds != "" {
		buffer.WriteString(hundreds)
		if tens != "" || teens != "" || ones != "" {
			buffer.WriteString(and)
		}
	}
	if tens != "" {
		buffer.WriteString(hundreds)
	}

	return buffer.String()
}

func main() {

	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += len(speaknspell(i))
	}

	fmt.Println(sum)
}
