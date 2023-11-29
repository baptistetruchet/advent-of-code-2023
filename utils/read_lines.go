package utils

import (
	"bufio"
	"os"
)

func ReadLines[C any](filePath string, lines chan<- C, fn func(string) C) {
	defer close(lines)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- fn(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
