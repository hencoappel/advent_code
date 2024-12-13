package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var whitespace, _ = regexp.Compile(" +")

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int %v", err)
	}
	return i
}

func readIntLine(line string) []int {
	strs := whitespace.Split(line, -1)
	res := make([]int, len(strs))
	for i, s := range strs {
		if s[len(s)-1] == ':' {
			s = s[:len(s)-1]
		}
		res[i] = toInt(s)
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pair [2]int

func canMake(target int, nums []int) bool {
	totals := []int{nums[0]}
	for _, num := range nums[1:] {
		newtotals := make([]int, 0, len(totals)*2)
		for _, t := range totals {
			newtotals = append(newtotals, t+num, t*num)
		}
		totals = newtotals
	}
	for _, t := range totals {
		if t == target {
			return true
		}
	}
	return false
}

func solve1(reader *bufio.Scanner) {
	total := 0
	for reader.Scan() {
		line := readIntLine(reader.Text())
		target, nums := line[0], line[1:]
		if canMake(target, nums) {
			total += target
		}
	}
	fmt.Println(total)
}

func concat(a, b int) int {
	magOfB := (int(math.Floor(math.Log10(float64(b)))) + 1)
	return int(math.Pow10(magOfB))*a + b
}

func canMake2(target int, nums []int) bool {
	totals := []int{nums[0]}
	for _, num := range nums[1:] {
		newtotals := make([]int, 0, len(totals)*2)
		for _, t := range totals {
			if t <= target {
				newtotals = append(newtotals, t+num, t*num, concat(t, num))
			}
		}
		totals = newtotals
	}
	for _, t := range totals {
		if t == target {
			return true
		}
	}
	return false
}

func solve2(reader *bufio.Scanner) {
	total := 0
	for reader.Scan() {
		line := readIntLine(reader.Text())
		target, nums := line[0], line[1:]
		if canMake2(target, nums) {
			total += target
		}
	}
	fmt.Println(total)
}

func main() {
	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 130*131+10), bufio.MaxScanTokenSize)
	// solve1(scanner)
	solve2(scanner)
}
