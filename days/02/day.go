package day02

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	"github.com/samber/lo"
)

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	ids := []int{}

	r1 := regexp.MustCompile(`Game (\d+): (.+)`)

	for line := range lines {
		match := r1.FindStringSubmatch(line)
		games := strings.Split(match[2], "; ")
		isValid := true
		for _, game := range games {
			m := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			for _, p := range strings.Split(game, ", ") {
				g := strings.Split(p, " ")
				n, l := utils.ParseInt(g[0]), g[1]
				m[l] += n
			}
			if m["red"] > 12 || m["green"] > 13 || m["blue"] > 14 {
				isValid = false
				break
			}
		}
		if isValid {
			ids = append(ids, utils.ParseInt(match[1]))
		}
	}

	fmt.Println(lo.Sum(ids))
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	powers := []int{}

	r1 := regexp.MustCompile(`Game (\d+): (.+)`)

	for line := range lines {
		match := r1.FindStringSubmatch(line)
		games := strings.Split(match[2], "; ")
		max := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, game := range games {
			m := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			for _, p := range strings.Split(game, ", ") {
				g := strings.Split(p, " ")
				n, l := utils.ParseInt(g[0]), g[1]
				m[l] += n
			}
			max["red"] = lo.Max([]int{max["red"], m["red"]})
			max["blue"] = lo.Max([]int{max["blue"], m["blue"]})
			max["green"] = lo.Max([]int{max["green"], m["green"]})
		}
		powers = append(powers, max["red"]*max["blue"]*max["green"])
	}

	fmt.Println(lo.Sum(powers))
}
