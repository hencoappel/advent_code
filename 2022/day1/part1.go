package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	maxVal := 0
	current := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			maxVal = max(maxVal, current)
			current = 0
			continue
		}
		val, _ := strconv.Atoi(l)
		current += val
	}
	fmt.Printf("highest: %d\n", maxVal)
}

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	reader := bufio.NewReader(f)
	solve(reader)
}
