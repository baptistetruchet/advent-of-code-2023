package main

import (
	"fmt"
	"os"

	"github.com/baptistetruchet/advent-of-code-2023/utils"
	_ "github.com/samber/lo"
)

func main() {
	filePath := os.Args[2]
	lines := make(chan string)

	go utils.ReadLines(filePath, lines, func(line string) string { return line })

	fmt.Println(utils.ChanToSlice(lines))
}
