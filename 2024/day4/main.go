package main

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pair [2]int

var moves = []pair{
	{0, -1},  // up
	{1, -1},  // up right
	{1, 0},   // right
	{1, 1},   // down right
	{0, 1},   // down
	{-1, 1},  // down left
	{-1, 0},  // left
	{-1, -1}, // up left
}

func getMoves(x, y int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for _, move := range moves {
			if !yield(x+move[0], y+move[1]) {
				return
			}
		}
	}
}

type grid [][]byte

func (g grid) inBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(g[0]) && y < len(g)
}

func searchDirection(grid grid, start, dir pair) bool {
	x, y := start[0], start[1]
	for _, b := range []byte("MAS") {
		x, y = x+dir[0], y+dir[1]
		if !grid.inBounds(x, y) || grid[y][x] != b {
			// fmt.Println("hit wall, breaking", x, y, inBounds(x, y), b)
			// if inBounds(x, y) {
			// 	fmt.Println("failed char wall, breaking", x, y, grid[y][x], b)
			// }
			return false
		}
		// fmt.Println("found next letter at", x, y, b)
	}
	return true
}

type xmas struct {
	location  pair
	direction pair
}

func search(grid grid, startx, starty int) []xmas {
	if grid[starty][startx] != 'X' {
		return []xmas{}
	}
	// fmt.Println("found an x at", startx, starty)
	// count := 0
	found := make([]xmas, 0)
	for _, move := range moves {
		if searchDirection(grid, pair{startx, starty}, move) {
			// count++
			found = append(found, xmas{
				location:  pair{startx, starty},
				direction: move,
			})
		}
	}
	return found
}

func inword(found []xmas, x, y int) bool {
	for _, found := range found {
		checkx, checky := found.location[0], found.location[1]
		for range "XMAS" {
			if checkx == x && checky == y {
				return true
			}
			checkx, checky = checkx+found.direction[0], checky+found.direction[1]
		}
	}
	return false
}

func searchMASCross(grid grid, startx, starty int) []xmas {
	if grid[starty][startx] != 'A' {
		return []xmas{}
	}
	found := make([]xmas, 0)
	const expected = 'M' + 'S' // 77+83=160
	backslash := grid[starty-1][startx-1] + grid[starty+1][startx+1]
	forwardslash := grid[starty+1][startx-1] + grid[starty-1][startx+1]
	if backslash == expected && forwardslash == expected {
		found = append(found, xmas{
			location:  pair{startx, starty},
			direction: pair{0, 0},
		})
	}
	return found
}

func solve1(reader *bufio.Scanner) {
	grid := readBytes(reader)
	found := make([]xmas, 0)
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			found = append(found, search(grid, x, y)...)
		}
	}
	fmt.Println(len(grid), "x", len(grid[0]))
	fmt.Println(len(found))
}

func readBytes(reader *bufio.Scanner) [][]byte {
	grid := make([][]byte, 0)
	for reader.Scan() {
		grid = append(grid, bytes.Clone(reader.Bytes()))
	}
	return grid
}
func solve2(reader *bufio.Scanner) {
	grid := readBytes(reader)
	found := make([]xmas, 0)
	for x := 1; x < len(grid[0])-1; x++ {
		for y := 1; y < len(grid)-1; y++ {
			found = append(found, searchMASCross(grid, x, y)...)
		}
	}
	fmt.Println(len(grid), "x", len(grid[0]))
	fmt.Println(len(found))
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
