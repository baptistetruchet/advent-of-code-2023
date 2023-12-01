package main

import (
	"fmt"
	"time"

	day01 "github.com/baptistetruchet/advent-of-code-2023/days/01"
)

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()

	day01.SolveTwo()
}
