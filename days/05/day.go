package day05

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	"github.com/samber/lo"
)

func findClosestLowerValue(arr []int, v int) (res int, j int, f bool) {
	for i, e := range arr {
		if e <= v && (!f || e > res) {
			f = true
			res = e
			j = i
		}
	}
	return
}

func mapNext(s []int, m [][]int) []int {
	c := lo.Map(m, func(b []int, _ int) int { return b[1] })
	return lo.Map(s, func(a int, i int) int {
		v, j, ok := findClosestLowerValue(c, a)
		if ok && (a < v+(m[j][2])) {
			return m[j][0] + (a - v)
		}
		return a
	})
}

func SolveOne() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	s := lo.Map(lo.Flatten(utils.FindAllStringGroup(regexp.MustCompile(`(\d+)`), <-lines)), func(a string, _ int) int { return utils.ParseInt(a) })
	<-lines
	<-lines
	var m [][]int

	for line := range lines {
		if line == "" {
			s = mapNext(s, m)
			m = [][]int{}
		}
		g := utils.FindAllStringGroup(regexp.MustCompile(`(\d+) (\d+) (\d+)`), line)
		if len(g) > 0 {
			m = append(m, lo.Map(g[0], func(a string, _ int) int { return utils.ParseInt(a) }))
		}
	}
	s = mapNext(s, m)

	fmt.Println(lo.Min(s))
}

type Range struct {
	start  int
	length int
}

type RangeMap struct {
	startDestination int
	startSource      int
	length           int
}

func mapNext2(s []Range, m []RangeMap) (res []Range) {
	sort.SliceStable(m, func(i, j int) bool {
		return m[i].startSource < m[j].startSource
	})

	for _, r := range s {
		var lastRm *RangeMap

		for _, rm := range m {
			if r.start >= rm.startSource+rm.length {
				continue
			}

			if rm.startSource >= r.start+r.length {
				break
			}

			lastRm = &rm

			diffStart := r.start - rm.startSource
			if diffStart >= 0 {
				res = append(res, Range{rm.startDestination + diffStart, lo.Min([]int{rm.length - diffStart, r.length})})
			} else {
				res = append(res, Range{r.start, -diffStart}, Range{rm.startDestination, lo.Min([]int{rm.length, r.length + diffStart})})
			}
		}

		if lastRm == nil {
			res = append(res, r)
		} else {
			diffEnd := r.start + r.length - lastRm.startSource - lastRm.length
			if diffEnd > 0 {
				res = append(res, Range{r.start + r.length - diffEnd, diffEnd})
			}
		}
	}

	return
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	s := lo.Map(utils.FindAllStringGroup(regexp.MustCompile(`(\d+) (\d+)`), <-lines), func(a []string, _ int) Range {
		return Range{utils.ParseInt(a[0]), utils.ParseInt(a[1])}
	})

	<-lines
	<-lines
	var m []RangeMap

	for line := range lines {
		if line == "" {
			s = mapNext2(s, m)
			m = []RangeMap{}
		}
		g := utils.FindAllStringGroup(regexp.MustCompile(`(\d+) (\d+) (\d+)`), line)
		if len(g) > 0 {
			m = append(m, RangeMap{utils.ParseInt(g[0][0]), utils.ParseInt(g[0][1]), utils.ParseInt(g[0][2])})
		}
	}
	s = mapNext2(s, m)

	fmt.Println(lo.Min(lo.Map(s, func(r Range, _ int) int { return r.start })))
}
