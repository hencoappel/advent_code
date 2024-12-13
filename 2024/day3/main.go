package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func solve1(reader *bufio.Scanner) {
	str := ""
	for reader.Scan() {
		str += reader.Text()
	}
	rgx, err := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	if err != nil {
		log.Fatalln("Failed to compile regex", err)
	}
	matches := rgx.FindAllStringSubmatch(str, -1)
	sum := 0
	for _, match := range matches {
		sum += toInt(match[1]) * toInt(match[2])
	}
	fmt.Println(sum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pair [2]int

func eq(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func matchInt(buffer []byte, i int) (int, int) {
	var res int
	starti := i
	for ; i < i+3 && i < len(buffer); i++ {
		b := buffer[i]
		if b >= '0' && b <= '9' {
			res = res*10 + int(b-'0')
		} else {
			return res, i - starti
		}
	}
	// check if 4th digit number, only 3 digit numbers allowed
	if i+3 < len(buffer) && buffer[i+3] >= '0' && buffer[i+3] <= '9' {
		return 0, 0
	}
	return res, i - starti
}

func solve2(reader *bufio.Reader) {
	buffer, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalln("failed to read", err)
	}
	sum := 0
	enabled := true
	for i := 0; i < len(buffer); i++ {
		b := buffer[i]
		switch b {
		case 'm':
			if !enabled {
				continue
			}
			// mul(1,1) <- min length 8 so check i+7
			if i+7 < len(buffer) && eq(buffer[i:i+4], []byte("mul(")) {
				i = i + 4
				val1, move := matchInt(buffer, i)
				if move == 0 { // didn't read int, continue
					continue
				}
				i += move
				if buffer[i] != ',' {
					continue
				}
				i++
				val2, move := matchInt(buffer, i)
				if move == 0 { // didn't read int, continue
					continue
				}
				i += move
				if buffer[i] != ')' {
					continue
				}
				sum += val1 * val2
			}
		case 'd':
			if i+6 < len(buffer) && eq(buffer[i:i+7], []byte("don't()")) {
				enabled = false
			}
			if i+3 < len(buffer) && eq(buffer[i:i+4], []byte("do()")) {
				enabled = true
			}
		}
	}
	fmt.Println("sum:", sum)
}

func main() {
	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	reader := bufio.NewReader(f)
	// scanner := bufio.NewScanner(reader)
	// solve1(scanner)
	solve2(reader)
}
