package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var winScore map[string]int = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var moveScore map[string]int = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func calcScore(a, b string) int {
	score := winScore[b]
	switch a {
	case "A":
		switch b {
		case "X":
			score += moveScore["C"]
		case "Y":
			score += moveScore["A"]
		case "Z":
			score += moveScore["B"]
		}
	case "B":
		switch b {
		case "X":
			score += moveScore["A"]
		case "Y":
			score += moveScore["B"]
		case "Z":
			score += moveScore["C"]
		}
	case "C":
		switch b {
		case "X":
			score += moveScore["B"]
		case "Y":
			score += moveScore["C"]
		case "Z":
			score += moveScore["A"]
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
