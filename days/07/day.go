package day07

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	"github.com/samber/lo"
)

// type
// 7 -> Five of a kind
// 6 -> Four of a kind
// 5 -> Full house
// 4 -> Three of a kind
// 3 -> Two pairs
// 2 -> One pair
// 1 -> Highest card

func findType(h string) string {
	m := map[string]int{}
	for _, c := range h {
		m[string(c)] += 1
	}
	v := lo.Values(m)
	if lo.Contains(v, 5) {
		return "7"
	}
	if lo.Contains(v, 4) {
		return "6"
	}
	if lo.Contains(v, 3) && lo.Contains(v, 2) {
		return "5"
	}
	if lo.Contains(v, 3) {
		return "4"
	}
	if lo.Contains(v, 2) && len(v) == 3 {
		return "3"
	}
	if lo.Contains(v, 2) {
		return "2"
	}
	return "1"
}

var cardMap = map[string]string{
	"A": "14",
	"K": "13",
	"Q": "12",
	"J": "11",
	"T": "10",
	"9": "09",
	"8": "08",
	"7": "07",
	"6": "06",
	"5": "05",
	"4": "04",
	"3": "03",
	"2": "02",
}

func calcScore(h string) string {
	return findType(h) + cardMap[string(h[0])] + cardMap[string(h[1])] + cardMap[string(h[2])] + cardMap[string(h[3])] + cardMap[string(h[4])]
}

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r1 := regexp.MustCompile(`(.{5}) (\d+)`)

	s := map[string]string{}
	b := map[string]int{}
	c := []string{}

	for line := range lines {
		g := utils.FindAllStringGroup(r1, line)[0]
		s[g[0]] = calcScore(g[0])
		b[g[0]] = utils.ParseInt(g[1])
		c = append(c, g[0])
	}

	sort.Slice(c, func(i, j int) bool {
		return s[c[i]] < s[c[j]]
	})

	res := 0
	for i, h := range c {
		res += (i + 1) * b[h]
	}
	fmt.Println(res)
}

func findType2(h string) string {
	m := map[string]int{}
	for _, c := range h {
		m[string(c)] += 1
	}
	j := m["J"]
	if j > 0 {
		delete(m, "J")
		max, l := 0, ""
		for k, v := range m {
			if v > max {
				max = v
				l = k
			}
		}
		m[l] += j
	}
	v := lo.Values(m)
	if lo.Contains(v, 5) {
		return "7"
	}
	if lo.Contains(v, 4) {
		return "6"
	}
	if lo.Contains(v, 3) && lo.Contains(v, 2) {
		return "5"
	}
	if lo.Contains(v, 3) {
		return "4"
	}
	if lo.Contains(v, 2) && len(v) == 3 {
		return "3"
	}
	if lo.Contains(v, 2) {
		return "2"
	}
	return "1"
}

var cardMap2 = map[string]string{
	"A": "14",
	"K": "13",
	"Q": "12",
	"T": "10",
	"9": "09",
	"8": "08",
	"7": "07",
	"6": "06",
	"5": "05",
	"4": "04",
	"3": "03",
	"2": "02",
	"J": "01",
}

func calcScore2(h string) string {
	return findType2(h) + cardMap2[string(h[0])] + cardMap2[string(h[1])] + cardMap2[string(h[2])] + cardMap2[string(h[3])] + cardMap2[string(h[4])]
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r1 := regexp.MustCompile(`(.{5}) (\d+)`)

	s := map[string]string{}
	b := map[string]int{}
	c := []string{}

	for line := range lines {
		g := utils.FindAllStringGroup(r1, line)[0]
		s[g[0]] = calcScore2(g[0])
		b[g[0]] = utils.ParseInt(g[1])
		c = append(c, g[0])
	}

	sort.Slice(c, func(i, j int) bool {
		return s[c[i]] < s[c[j]]
	})

	res := 0
	for i, h := range c {
		res += (i + 1) * b[h]
	}
	fmt.Println(res)
}
