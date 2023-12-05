package day04

import (
	"fmt"
	"math"
	"os"
	"regexp"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	"github.com/samber/lo"
)

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r1 := regexp.MustCompile(`Card\s+\d+: ([\d\s]+) \| ([\d\s]+)`)
	r2 := regexp.MustCompile(`(\d+)`)

	var res float64 = 0

	for line := range lines {
		m := r1.FindAllStringSubmatch(line, -1)[0]
		w := lo.Map(r2.FindAllStringSubmatch(m[1], -1), func(a []string, i int) string { return a[0] })
		d := lo.Map(r2.FindAllStringSubmatch(m[2], -1), func(a []string, i int) string { return a[0] })
		n := len(lo.Intersect(w, d))
		if n > 0 {
			res += 1 * math.Pow(2, float64(n-1))
		}
	}

	fmt.Println(res)
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r1 := regexp.MustCompile(`Card\s+(\d+): ([\d\s]+) \| ([\d\s]+)`)
	r2 := regexp.MustCompile(`(\d+)`)

	c := make(map[int]int)

	for line := range lines {
		m := utils.FindAllStringGroup(r1, line)[0]
		g := utils.ParseInt(m[0])
		w := lo.Flatten(utils.FindAllStringGroup(r2, m[1]))
		d := lo.Flatten(utils.FindAllStringGroup(r2, m[2]))
		n := len(lo.Intersect(w, d))

		c[g] = c[g] + 1

		for i := 1; i <= n; i++ {
			c[g+i] = c[g+i] + c[g]
		}
	}

	fmt.Println(lo.Sum(lo.Values(c)))
}
