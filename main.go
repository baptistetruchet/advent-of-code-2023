package main

import (
	"fmt"

	"github.com/baptistetruchet/aoc2023/utils"
	_ "github.com/samber/lo"
)

func main() {
	lines := make(chan string)

	go utils.ReadLines("input.txt", lines, func(line string) string { return line })

	fmt.Println(utils.ChanToSlice(lines))
}
