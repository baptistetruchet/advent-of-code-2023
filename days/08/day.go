package day08

import (
	"fmt"
	"os"
	"regexp"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	"github.com/samber/lo"
)

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r1 := regexp.MustCompile(`(\w{3}) = .(\w{3}), (\w{3}).`)
	m := make(map[string][]string)

	inst := utils.Split(<-lines)
	<-lines

	for line := range lines {
		g := utils.FindAllStringGroup(r1, line)[0]
		m[g[0]] = []string{g[1], g[2]}
	}

	pos := "AAA"
	steps := 0
	i := 0

	for {
		if pos == "ZZZ" {
			break
		}
		pos = m[pos][lo.Ternary(inst[i] == "L", 0, 1)]
		i = (i + 1) % len(inst)
		steps++
	}

	fmt.Println(steps)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, ints ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(ints); i++ {
		result = lcm(result, ints[i])
	}

	return result
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r1 := regexp.MustCompile(`(\w{3}) = .(\w{3}), (\w{3}).`)
	m := make(map[string][]string)

	inst := utils.Split(<-lines)
	pos := []string{}
	<-lines

	for line := range lines {
		g := utils.FindAllStringGroup(r1, line)[0]
		m[g[0]] = []string{g[1], g[2]}
		if string(g[0][2]) == "A" {
			pos = append(pos, g[0])
		}
	}

	steps := make([]int, len(pos))

	for j, p := range pos {
		step := 0
		i := 0
		q := p
		for {
			if string(q[2]) == "Z" {
				break
			}
			q = m[q][lo.Ternary(inst[i] == "L", 0, 1)]
			i = (i + 1) % len(inst)
			step++
		}
		steps[j] = step
	}

	fmt.Println(lcm(steps[0], steps[1], steps[2:]...))
}
