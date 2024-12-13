package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var whitespace, _ = regexp.Compile(" +")

func readMove(line string) point {
	splits := strings.SplitN(line, "+", 3)
	x, _ := strconv.Atoi(splits[1][:strings.IndexByte(splits[1], ',')])
	y, _ := strconv.Atoi(splits[2])
	return point{x, y}
}
func readPrize(line string) point {
	splits := strings.SplitN(line, "=", 3)
	x, _ := strconv.Atoi(splits[1][:strings.IndexByte(splits[1], ',')])
	y, _ := strconv.Atoi(splits[2])
	return point{x, y}
}

func readMachine(reader *bufio.Scanner) (point, point, point) {
	a := readMove(reader.Text())
	reader.Scan()
	b := readMove(reader.Text())
	reader.Scan()
	target := readPrize(reader.Text())
	reader.Scan()
	return a, b, target
}

func divisible(num1, denom1, num2, denom2 int) (int, bool) {
	// f := float64(num1) / float64(denom1)
	// i, frac := math.Modf(f)
	frac := num1 % denom1
	if frac != 0 {
		return 0, false
	}
	i := num1 / denom1
	if denom2*int(i) != num2 {
		return 0, false
	}
	return int(i), true
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(reader *bufio.Scanner, shiftTarget int) int {
	total := 0
	for reader.Scan() {
		a, b, target := readMachine(reader)
		// fmt.Println(a, b, target)
		target = target.Add(point{shiftTarget, shiftTarget})
		currentx, currenty := target.x, target.y
		bcount := 0
		for currentx > 0 && currenty > 0 {
			// fmt.Println(current)
			// diff := target.Sub(current)
			acount, ok := divisible(currentx, a.x, currenty, a.y)
			// fmt.Println(currentx, currenty)
			if ok {
				cost := acount*3 + bcount*1
				total += cost
				fmt.Println("found one for", acount, bcount)
				break
			}
			currentx -= b.x
			currenty -= b.y
			bcount++
		}
		fmt.Println("finished ", a, b, target)
	}
	return total
}

func solve1(reader *bufio.Scanner) {
	fmt.Println(solve(reader, 0))
}

func solveLinear(reader *bufio.Scanner, shiftTarget int) int {
	total := 0
	for reader.Scan() {
		n1, n2, target := readMachine(reader)
		a := float64(n1.x)
		b := float64(n2.x)
		c := float64(shiftTarget + target.x)
		d := float64(n1.y)
		e := float64(n2.y)
		f := float64(shiftTarget + target.y)
		y := (c*d - f*a) / (b*d - e*a)
		x := (c - b*y) / a
		if math.Trunc(x) == x && math.Trunc(y) == y {
			cost := int(x*3 + y*1)
			total += cost
		}
	}
	return total
}

func solve2(reader *bufio.Scanner) {
	// 240202054
	// 10000000000000
	// Button A: X+94, Y+34
	// Button B: X+22, Y+67
	// Prize: X=8400, Y=5400
	// 8400/94=89,8400%9436
	// a, b := math.Mod(float64(8400), float64(22))
	// return
	fmt.Println(solveLinear(reader, 10000000000000))
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
