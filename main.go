package main

import (
	"fmt"
	"time"

	day "github.com/baptistetruchet/advent-of-code-2023/days/06"
)

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()

	day.SolveTwo()
}
