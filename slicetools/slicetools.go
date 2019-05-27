package slicetools

import (
	"bytes"
	"strconv"
	"strings"
)

func Reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Rotate(slc []string, rot int) {
	Reverse(slc[rot:])
	Reverse(slc[:rot])
	Reverse(slc)
}

func SplitDigits(n int) []int {
	nStr := strconv.Itoa(n)
	digits := strings.Split(nStr, "")
	intDigits := make([]int, len(digits))
	for i, d := range digits {
		digit, _ := strconv.Atoi(d)
		intDigits[i] = digit
	}
	return intDigits
}

func CombineDigits(digits []int) int {
	var b bytes.Buffer
	for _, d := range digits {
		b.WriteString(strconv.Itoa(d))
	}
	i, err := strconv.Atoi(b.String())
	if err != nil {
		panic("Bad Strconv")
	}
	return i
}
