package main

/*
Names scores
Problem 22
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/
import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var letterScores = make(map[string]int)

func init() {
	curLetter := 'A'
	for i := 1; i < 27; i++ {
		letterScores[string(curLetter)] = i
		curLetter++
	}

}

func nameScore(name string, index int) int {
	index++ //compensate for 0 index on the array
	score := 0
	letters := strings.Split(name, "")
	//trim double quotes
	letters = letters[1 : len(letters)-1]
	for _, letter := range letters {
		letterScore, _ := letterScores[letter]
		score += letterScore
	}
	score *= index
	fmt.Printf("%s,%d:%d\n", name, index, score)
	return score
}

func main() {

	namesData, err := ioutil.ReadFile("./p022_names.txt")
	if err != nil {
		panic(err)
	}
	names := strings.Split(string(namesData), ",")

	//sort the names first
	sort.Strings(names)

	fmt.Println(letterScores)
	total := 0
	for index, name := range names {
		total += nameScore(name, index)
	}
	fmt.Println(total)
}
