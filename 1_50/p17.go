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

var numMap = map[int]string{
	0:   "",
	1:   "one",
	2:   "two",
	3:   "three",
	4:   "four",
	5:   "five",
	6:   "six",
	7:   "seven",
	8:   "eight",
	9:   "nine",
	11:  "eleven",
	12:  "twelve",
	13:  "thirteen",
	14:  "fourteen",
	15:  "fifteen",
	16:  "sixteen",
	17:  "seventeen",
	18:  "eighteen",
	19:  "nineteen",
	10:  "ten",
	20:  "twenty",
	30:  "thirty",
	40:  "forty",
	50:  "fifty",
	60:  "sixty",
	70:  "seventy",
	80:  "eighty",
	90:  "ninety",
	100: "onehundred",
	200: "twohundred",
	300: "threehundred",
	400: "fourhundred",
	500: "fivehundred",
	600: "sixhundred",
	700: "sevenhundred",
	800: "eighthundred",
	900: "ninehundred",
}

var oneThousand = "onethousand"
var and = "and"

func splitIntoSpeakables(i int) (hundreds, tens, teens, ones string) {
	h := 0
	ten := 0
	teen := 0
	one := 0
	if i < 10 {
		one = i
	} else if i < 20 {
		teen = i
	} else if i < 100 {
		one = i % 10
		ten = (i / 10) * 10
	} else if i < 1000 {
		h = (i / 100) * 100
		t := i % 100
		if t < 20 && t > 10 {
			teen = t
		} else {
			ten = (t / 10) * 10
			one = t % 10
		}
	}
	//fmt.Printf("%d, %d, %d, %d, %d\n", i, h, ten, teen, one)
	hundreds = numMap[h]
	tens = numMap[ten]
	teens = numMap[teen]
	ones = numMap[one]
	return
}

func speaknspell(i int) string {
	hundreds, tens, teens, ones := splitIntoSpeakables(i)
	if i == 1000 {
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
		buffer.WriteString(tens)
	}
	if teens != "" {
		buffer.WriteString(teens)
	}
	if ones != "" {
		buffer.WriteString(ones)
	}

	return buffer.String()
}

func main() {

	sum := 0
	for i := 1; i <= 1000; i++ {
		str := speaknspell(i)
		sum += len(str)
		fmt.Printf("%s, %d, %d\n", str, len(str), sum)
	}

	fmt.Println(sum)
}
