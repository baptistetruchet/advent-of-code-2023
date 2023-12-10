package day09

import (
	"fmt"
	"os"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	"github.com/samber/lo"
)

func predict(h []int, b bool) int {
	if lo.EveryBy(h, func(i int) bool { return i == 0 }) {
		return 0
	}

	n := make([]int, len(h)-1)
	for i := 0; i < len(h)-1; i++ {
		n[i] = h[i+1] - h[i]
	}
	if b {
		return h[len(h)-1] + predict(n, b)
	} else {
		return h[0] - predict(n, b)
	}
}

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan []string)

	go utils.ReadLines(filePath, lines, utils.SplitBy(" "))

	res := 0
	for line := range lines {
		d := lo.Map(line, func(s string, _ int) int { return utils.ParseInt(s) })
		res += predict(d, true)
	}
	fmt.Println(res)
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan []string)

	go utils.ReadLines(filePath, lines, utils.SplitBy(" "))

	res := 0
	for line := range lines {
		d := lo.Map(line, func(s string, _ int) int { return utils.ParseInt(s) })
		res += predict(d, false)
	}
	fmt.Println(res)
}
