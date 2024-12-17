package main

import (
	"bufio"
	"fmt"
	"math"
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

const (
	TOP_LEFT     = 0
	TOP_RIGHT    = 1
	BOTTOM_LEFT  = 2
	BOTTOM_RIGHT = 3
	NONE         = -1
)

func quadrant(x, y int) int {
	if x < midx {
		if y < midy {
			return TOP_LEFT
		} else if y > midy {
			return BOTTOM_LEFT
		}
	} else if x > midx {
		if y < midy {
			return TOP_RIGHT
		} else if y > midy {
			return BOTTOM_RIGHT
		}
	}
	return NONE
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

func safetyFactor(points []point) int {
	quads := make([]int, 4)
	for _, p := range points {
		q := quadrant(p.x, p.y)
		if q >= 0 {
			quads[q]++
		}
	}
	return quads[0] * quads[1] * quads[2] * quads[3]
}

func printGrid(bots []point) {
	grid := make([][]byte, by)
	for y := range by {
		grid[y] = slices.Repeat([]byte{' '}, bx)
	}
	for _, b := range bots {
		grid[b.y][b.x] = 'x'
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println(strings.Repeat("-", bx))
}

func step(bots, vs []point, steps int) []point {
	res := make([]point, len(bots))
	for i := range bots {
		res[i].x = (bots[i].x + vs[i].x*steps) % bx
		if res[i].x < 0 {
			res[i].x = bx + res[i].x
		}
		res[i].y = (bots[i].y + vs[i].y*steps) % by
		if res[i].y < 0 {
			res[i].y = by + res[i].y
		}
	}
	return res
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

func solve2(reader *bufio.Scanner) {
	// lots of bugs caused this to be far too hard to find the answer
	locations := make([]point, 0)
	velocities := make([]point, 0)
	for reader.Scan() {
		px, py, vx, vy := readRobot(reader.Text())
		locations = append(locations, point{px, py})
		velocities = append(velocities, point{vx, vy})
	}
	// fmt.Println(velocities)
	// fmt.Println(locations)
	// n := by * bx
	// n := 100
	n := 20_000
	minStep := 0
	minLocations := locations
	minSafetyFactor := math.MaxInt
	for i := 1; i <= n; i++ {
		newlocs := step(locations, velocities, i)
		sf := safetyFactor(newlocs)
		// fmt.Println(newlocs[0])
		if horizontalInRow(newlocs, 7) {
			fmt.Println("found", i)
			printGrid(newlocs)
		}
		if sf < minSafetyFactor {
			minStep = i
			minSafetyFactor = sf
			minLocations = newlocs
			// printGrid(minLocations)
		}
	}
	fmt.Println(len(minLocations))
	// printGrid(minLocations)
	fmt.Println(minStep)
	fmt.Println(minSafetyFactor)
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
