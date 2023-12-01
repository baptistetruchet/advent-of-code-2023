package day01

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
)

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r := regexp.MustCompile(`\d`)

	res := 0
	for line := range lines {
		nums := r.FindAllString(line, -1)
		a := nums[0]
		b := nums[len(nums)-1]
		res += utils.ParseInt(a + b)
	}

	fmt.Println(res)
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	mapWords := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
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
		res += utils.ParseInt(a + b)
	}

	fmt.Println(res)
}
