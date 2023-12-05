package day03

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	"github.com/samber/lo"
)

func nextToSym(grid []string, j, i1, i2 int) bool {
	s := ""
	rMin := lo.Max([]int{i1 - 1, 0})
	rMax := lo.Min([]int{i2 + 1, len(grid[0]) - 1})

	if j > 0 {
		s += grid[j-1][rMin:rMax]
	}
	s += grid[j][rMin:rMax]
	if j < len(grid)-1 {
		s += grid[j+1][rMin:rMax]
	}

	r := regexp.MustCompile(`[^\d^\.]`)

	return r.MatchString(s)
}

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	grid := utils.ChanToSlice(lines)

	r := regexp.MustCompile(`(\d+)`)

	sum := 0

	for j, line := range grid {
		m := r.FindAllStringSubmatchIndex(line, -1)
		for _, rg := range m {
			if nextToSym(grid, j, rg[0], rg[1]) {
				sum += utils.ParseInt(line[rg[0]:rg[1]])
			}
		}
	}

	fmt.Println(sum)
}

func findAdj(grid []string, i, j int) (nums []int) {
	lines := []string{}
	if i > 0 {
		lines = append(lines, grid[i-1])
	}
	lines = append(lines, grid[i])
	if i < len(grid)-1 {
		lines = append(lines, grid[i+1])
	}

	r := regexp.MustCompile(`(\d+)`)

	for _, line := range lines {
		m := r.FindAllStringSubmatchIndex(line, -1)
		for _, rg := range m {
			if (rg[1]-1) >= j-1 && rg[0] <= j+1 {
				nums = append(nums, utils.ParseInt(line[rg[0]:rg[1]]))
			}
		}
	}

	return
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	grid := utils.ChanToSlice(lines)

	sum := 0

	for i, line := range grid {
		for j, c := range strings.Split(line, "") {
			if c == "*" {
				nums := findAdj(grid, i, j)
				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Println(sum)
}
