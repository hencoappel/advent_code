package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func readRobot(line string) (int, int, int, int) {
	split := strings.Split(line, " ")
	pstrs := strings.Split(split[0][2:], ",")
	vstrs := strings.Split(split[1][2:], ",")
	px, py := toInt(pstrs[0]), toInt(pstrs[1])
	vx, vy := toInt(vstrs[0]), toInt(vstrs[1])
	return px, py, vx, vy
}

func quadrant(x, y int) int {
	if x < midx {
		if y < midy {
			return 0
		} else if y > midy {
			return 2
		}
	} else if x > midx {
		if y < midy {
			return 1
		} else if y > midy {
			return 3
		}
	}
	return -1
}

var bx, by = 101, 103 // bounds
// var bx, by = 11, 7 // example bounds
var midx, midy = bx / 2, by / 2

func solve(reader *bufio.Scanner, secs int) {
	quads := make([]int, 4)
	fmt.Println("mids", midx, midy)
	for reader.Scan() {
		px, py, vx, vy := readRobot(reader.Text())
		newx := (px + vx*secs) % bx
		newy := (py + vy*secs) % by
		// fmt.Println("before", newx, newy, q)
		if newx < 0 {
			newx = bx + newx
		}
		if newy < 0 {
			newy = by + newy
		}
		q := quadrant(newx, newy)
		fmt.Println("after", newx, newy, q)
		if q >= 0 {
			quads[q]++
		}
		// fmt.Println(px, py, vx, vy)
	}
	fmt.Println(quads)
	fmt.Println(quads[0] * quads[1] * quads[2] * quads[3])
}

func solve1(reader *bufio.Scanner) {
	solve(reader, 100)
}

func horizontalInRow(bots []point, n int) bool {
	ymap := make(map[int][]point, by)
	for _, b := range bots {
		if s, ok := ymap[b.y]; !ok {
			ymap[b.y] = []point{b}
		} else {

			ymap[b.y] = append(s, b)
		}
	}
	for _, points := range ymap {
		if len(points) < n {
			continue
		}
		sort.Slice(points, func(i, j int) bool {
			return points[i].x < points[j].x
		})
		currentCount := 1
		x := points[0].x
		for _, p := range points[1:] {
			if p.x == x+1 {
				currentCount++
			} else {
				currentCount = 0
			}
			x = p.x
			if currentCount == n {
				return true
			}
		}
	}
	return false
}

func printGrid(bots []point) {
	grid := make([][]byte, by)
	for y := range by {
		grid[y] = slices.Repeat([]byte{' '}, bx)
	}
	for _, b := range bots {
		if grid[b.y][b.x] == ' ' {
			grid[b.y][b.x] = 'x'
		} else {
			// grid[b.y][b.x]++
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func step(bots, vs []point) {
	for i := range bots {
		v := vs[i]
		bots[i].x = (bots[i].x + v.x) % bx
		if bx < 0 {
			bots[i].x = bx + bots[i].x
		}
		bots[i].y = (bots[i].y + v.y) % by
		if by < 0 {
			bots[i].y = by + bots[i].y
		}
	}
}

func solve2(reader *bufio.Scanner) {
	locations := make([]point, 0)
	veloctities := make([]point, 0)
	for reader.Scan() {
		px, py, vx, vy := readRobot(reader.Text())
		locations = append(locations, point{px, py})
		veloctities = append(locations, point{vx, vy})
	}
	fmt.Println(len(veloctities))
	// printGrid(locations)
	for range 10000000 {
		step(locations, veloctities)
		if horizontalInRow(locations, 4) {
			printGrid(locations)
			fmt.Println("---------------------------------------------------------")
		}
	}
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
