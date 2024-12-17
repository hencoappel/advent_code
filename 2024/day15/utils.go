package main

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"regexp"
	"strconv"
)

var whitespace, _ = regexp.Compile(" +")

func toInt(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type point struct {
	x int
	y int
}

func (p point) Add(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}
func (p point) Sub(p2 point) point {
	return point{p.x - p2.x, p.y - p2.y}
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

func readGrid(reader *bufio.Scanner) grid {
	grid := make([][]byte, 0)
	for reader.Scan() {
		b := reader.Bytes()
		if len(b) == 0 {
			break
		}
		row := bytes.Clone(b)
		for i, b := range row {
			row[i] = b
		}
		grid = append(grid, row)
	}
	return grid
}

func (g grid) Coords() iter.Seq2[point, byte] {
	return func(yield func(point, byte) bool) {
		for y := range len(g) {
			for x := range len(g[0]) {
				p := point{int(x), int(y)}
				if !yield(p, g.Get(p)) {
					return
				}
			}
		}
	}
}
func (g grid) Get(p point) byte {
	// need to access as y/x
	return g[p.y][p.x]
}
func (g grid) Set(p point, b byte) {
	g[p.y][p.x] = b
}

func (g grid) AllMoves(p point) []point {
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
		moves = append(moves, newp)
	}
	return moves
}

func (g grid) Moves(p point) []point {
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

func (g grid) Print() {
	for y := range len(g) {
		fmt.Println(string(g[y]))
	}
}
