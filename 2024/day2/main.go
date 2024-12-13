package main

import (
	"bufio"
	"fmt"
	"iter"
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
		res[i] = toInt(s)
	}
	return res
}

func isSafe(report []int) bool {
	incCount := 0
	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]
		if a < b {
			incCount++
		}
		diff := math.Abs(float64(a - b))
		if diff == 0 || diff > 3 {
			return false
		}
	}
	return incCount == 0 || incCount == len(report)-1
}

func solve1(reader *bufio.Scanner) {
	reports := make([][]int, 0)
	for reader.Scan() {
		reports = append(reports, readIntLine(reader.Text()))
	}
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	fmt.Println(count)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pair [2]int

func zipAdjacentExcluding(s []int, exclude int) iter.Seq[pair] {
	return func(yield func(pair) bool) {
		for i := 0; i < len(s)-1; i++ {
			if i == exclude {
				continue
			}
			i2 := i + 1
			if i2 == exclude {
				i2++
			}
			if i2 >= len(s) {
				break
			}
			if !yield(pair{s[i], s[i2]}) {
				return
			}
		}
	}
}

func isSafeExcluding(report []int, exclude int) bool {
	incCount := 0
	for p := range zipAdjacentExcluding(report, exclude) {
		a := p[0]
		b := p[1]
		if a < b {
			incCount++
		}
		diff := math.Abs(float64(a - b))
		if diff == 0 || diff > 3 {
			return false
		}
	}
	return incCount == 0 || incCount == len(report)-2
}

func isSafeDampened(report []int) bool {
	if isSafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		if isSafeExcluding(report, i) {
			return true
		}
	}
	return false
}

func isSafeDampened2(report []int) bool {
	incdampens := 0
	decdampens := 0
	gapdampens := 0
	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]
		if a < b {
			incCount++
		}
		diff := math.Abs(float64(a - b))
		if diff == 0 || diff > 3 {
			return false
		}
	}
}

func solve2(reader *bufio.Scanner) {
	reports := make([][]int, 0)
	for reader.Scan() {
		reports = append(reports, readIntLine(reader.Text()))
	}
	count := 0
	for _, report := range reports {
		if isSafeDampened(report) {
			count++
			// fmt.Println("safe ", i, report)
		} else {
			// fmt.Println("not safe ", i, report)
		}
	}
	fmt.Println(count)
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
	// solve1(scanner)
	solve2(scanner)
}
