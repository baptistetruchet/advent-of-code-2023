package day06

import (
	"fmt"
	"math"
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

	r1 := regexp.MustCompile(`(\d+)`)
	time := lo.Flatten(utils.FindAllStringGroup(r1, <-lines))
	distance := lo.Flatten(utils.FindAllStringGroup(r1, <-lines))

	res := 1

	for i, s := range time {
		t := utils.ParseInt(s)
		d := utils.ParseInt(distance[i])
		n := 0

		for i := 0; i <= t; i++ {
			v := i
			if v*(t-i) > d {
				n += 1
			}
		}

		res *= n
	}

	fmt.Println(res)
}

func SolveTwo() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, utils.Identity)

	r1 := regexp.MustCompile(`(\d+)`)
	t := utils.ParseInt(strings.Join(lo.Flatten(utils.FindAllStringGroup(r1, <-lines)), ""))
	d := utils.ParseInt(strings.Join(lo.Flatten(utils.FindAllStringGroup(r1, <-lines)), ""))

	a := -1
	b := t
	c := -d

	det := b*b - 4*a*c

	x1 := (float64(-b) + math.Sqrt(float64(det))) / (2 * float64(a))
	x2 := (float64(-b) - math.Sqrt(float64(det))) / (2 * float64(a))

	fmt.Println(int(math.Floor(x2) - math.Ceil(x1) + 1))
}
