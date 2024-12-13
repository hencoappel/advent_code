package main

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"os"
	"strconv"
	"time"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type num = int

type grid [][]byte

func (g grid) inBounds(p pair) bool {
	return p._0 >= 0 && p._1 >= 0 && p._0 < num(len(g[0])) && p._1 < num(len(g))
}

func (g grid) Coords() iter.Seq2[pair, byte] {
	return func(yield func(pair, byte) bool) {
		for y := range len(g) {
			for x := range len(g[0]) {
				p := pair{num(x), num(y)}
				if !yield(p, g.Get(p)) {
					return
				}
			}
		}
	}
}
func (g grid) Get(p pair) byte {
	// need to access as y/x
	return g[p._1][p._0]
}
func (g grid) Set(p pair, b byte) {
	g[p._1][p._0] = b
}
func readGrid(reader *bufio.Scanner) grid {
	grid := make([][]byte, 0)
	for reader.Scan() {
		grid = append(grid, bytes.Clone(reader.Bytes()))
	}
	return grid
}

type pair struct {
	_0 num
	_1 num
}

func (p pair) String() string {
	return strconv.FormatInt(int64(p._0), 10) + "," + strconv.FormatInt(int64(p._1), 10)
}

func readMap(reader *bufio.Scanner) (grid, pair) {
	grid := readGrid(reader)
	start := pair{}
	for p, b := range grid.Coords() {
		if b == '^' {
			start = p
			break
		}
	}
	return grid, start
}

type dir int

const (
	UP dir = iota
	RIGHT
	DOWN
	LEFT
)

func turn90(dir dir) dir {
	return (dir + 1) % 4
}

func (p pair) move(dir dir) pair {
	// top left is 0,0 so right is x increasing and down is y increasing
	newp := pair{p._0, p._1}
	switch dir {
	case UP:
		newp._1 -= 1
	case DOWN:
		newp._1 += 1
	case RIGHT:
		newp._0 += 1
	case LEFT:
		newp._0 -= 1
	}
	return newp
}

func fullPath(grid grid, start pair) []locdir {
	seenDir := make([]locdir, 0)
	p := start
	currentDir := UP
	for {
		seenDir = append(seenDir, locdir{p, currentDir})
		nextp := p.move(currentDir)
		if !grid.inBounds(nextp) {
			break
		}
		for grid.Get(nextp) == '#' { // doesn't handle nextp being out of bounds technically
			currentDir = turn90(currentDir)
			nextp = p.move(currentDir)
		}
		p = nextp
	}
	return seenDir
}

func solve1(reader *bufio.Scanner) {
	grid, start := readMap(reader)
	seenDir := fullPath(grid, start)
	seen := make(map[pair]struct{})
	for _, p := range seenDir {
		seen[p.p] = struct{}{}
	}
	fmt.Println(len(seen))
}

type locdir struct {
	p pair
	d dir
}

func causeLoop(grid grid, start locdir, block pair) bool {
	// seen := make(map[pair]struct{})
	seenDir := make(map[locdir]struct{})
	p := start.p
	currentDir := start.d
	for {
		// seen[p] = struct{}{}
		curLocDir := locdir{p, currentDir}
		if _, ok := seenDir[curLocDir]; ok { // been here in current direction, i.e. loop
			return true
		}
		seenDir[curLocDir] = struct{}{}
		nextp := p.move(currentDir)
		if !grid.inBounds(nextp) {
			return false
		}

		for grid.Get(nextp) == '#' || nextp == block {
			currentDir = turn90(currentDir)
			nextp = p.move(currentDir)
		}
		p = nextp
	}
}

func solve2(reader *bufio.Scanner) {
	t1 := time.Now()
	grid, start := readMap(reader)
	t2 := time.Now()
	seenDir := fullPath(grid, start)
	t3 := time.Now()
	count := 0
	seen := make(map[pair]struct{})
	seen[start] = struct{}{}
	for i := 1; i < len(seenDir); i++ {
		prevp := seenDir[i-1]
		p := seenDir[i]
		if _, ok := seen[p.p]; ok { // already checked, handles start
			continue
		}
		seen[p.p] = struct{}{}
		if causeLoop(grid, prevp, p.p) {
			count++
		}
	}
	t4 := time.Now()
	fmt.Println(t2.Sub(t1), t3.Sub(t2), t4.Sub(t3))
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
	scanner.Buffer(make([]byte, 130*131+10), bufio.MaxScanTokenSize)
	// solve1(scanner) // 41 - 5409
	solve2(scanner) // 6 - 2022
}
