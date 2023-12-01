package day01

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/baptistetruchet/advent-of-code-2023/parsing"
	"github.com/baptistetruchet/advent-of-code-2023/utils"
)

func main() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, parsing.Identity)

	mapWords := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	words := make([]string, 0, len(mapWords))
	rWords := make([]string, 0, len(mapWords))
	for k := range mapWords {
		words = append(words, k)
		rWords = append(rWords, utils.ReverseString(k))
	}

	r1 := regexp.MustCompile(fmt.Sprintf("(%v)", strings.Join(words, "|")))
	r2 := regexp.MustCompile(fmt.Sprintf("(%v)", strings.Join(rWords, "|")))

	res := 0
	for line := range lines {
		a := mapWords[r1.FindString(line)]
		b := mapWords[utils.ReverseString(r2.FindString(utils.ReverseString(line)))]

		res += utils.ParseInt(fmt.Sprintf("%v%v", a, b))
	}

	fmt.Println(res)
}
