package main

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"log"
	"maps"
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type point struct {
	x int
	y int
}

type dir int

const (
	UP dir = iota
	RIGHT
	DOWN
	LEFT
)

var dirs = []dir{UP, RIGHT, DOWN, LEFT}

type grid [][]byte

func (g grid) inBounds(p point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(g[0]) && p.y < len(g)
}

func readGrid(reader *bufio.Scanner) [][]byte {
	grid := make([][]byte, 0)
	for reader.Scan() {
		row := bytes.Clone(reader.Bytes())
		for i, b := range row {
			row[i] = b - '0'
		}
		grid = append(grid, row)
	}
	return grid
}

func (g grid) get(p point) byte {
	return g[p.y][p.x]
}

func (g grid) moves(p point) []point {
	// top left is 0,0 so right is x increasing and down is y increasing
	moves := make([]point, 0)
	for _, dir := range dirs {
		newp := point{p.x, p.y}
		switch dir {
		case UP:
			newp.y -= 1
		case DOWN:
			newp.y += 1
		case RIGHT:
			newp.x += 1
		case LEFT:
			newp.x -= 1
		}
		if g.inBounds(newp) {
			moves = append(moves, newp)
		}
	}
	return moves
}

func Single[E any](e E) iter.Seq[E] {
	return func(yield func(E) bool) {
		yield(e)
	}
}

func routes(grid grid, start point, val byte) iter.Seq[point] {
	if val == 9 {
		return Single(start)
	}
	nextVal := val + 1
	tops := make(map[point]struct{})
	for _, p := range grid.moves(start) {
		if grid.get(p) == nextVal {
			for top := range routes(grid, p, nextVal) {
				tops[top] = struct{}{}
			}
		}
	}
	return maps.Keys(tops)
}

func Count[E any](it iter.Seq[E]) int {
	total := 0
	for range it {
		total++
	}
	return total
}

func solve1(reader *bufio.Scanner) {
	grid := readGrid(reader)
	total := 0
	for y, row := range grid {
		for x, b := range row {
			if b == 0 {
				total += Count(routes(grid, point{x, y}, 0))
			}
		}
	}
	fmt.Println(total)
}

func numroutes(grid grid, start point, val byte) int {
	if val == 9 {
		return 1
	}
	nextVal := val + 1
	total := 0
	for _, p := range grid.moves(start) {
		if grid.get(p) == nextVal {
			total += numroutes(grid, p, nextVal)
		}
	}
	return total
}

func solve2(reader *bufio.Scanner) {
	grid := readGrid(reader)
	total := 0
	for y, row := range grid {
		for x, b := range row {
			if b == 0 {
				total += numroutes(grid, point{x, y}, 0)
			}
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
	// solve1(scanner)
	solve2(scanner)
}
