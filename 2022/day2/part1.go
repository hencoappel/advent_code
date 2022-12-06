package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var moveScore map[string]int = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func calcScore(a, b string) int {
	score := moveScore[b]
	switch a {
	case "A":
		switch b {
		case "X":
			score += 3
		case "Y":
			score += 6
		case "Z":
			score += 0
		}
	case "B":
		switch b {
		case "X":
			score += 0
		case "Y":
			score += 3
		case "Z":
			score += 6
		}
	case "C":
		switch b {
		case "X":
			score += 6
		case "Y":
			score += 0
		case "Z":
			score += 3
		}
	}
	return score
}

func solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	score := 0
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		score += calcScore(moves[0], moves[1])
	}
	fmt.Printf("%d\n", score)
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
