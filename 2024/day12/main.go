package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var (
	g    grid
	seen map[point]struct{}
)

func calcCost1(start point, b byte) int {
	moves := []point{start}
	var p point
	var area, perim int
	for len(moves) > 0 { // DFS over region
		p, moves = moves[len(moves)-1], moves[:len(moves)-1]
		if _, ok := seen[p]; ok {
			continue
		}
		seen[p] = struct{}{}
		area++
		for _, m := range g.AllMoves(p) {
			if g.inBounds(m) && g.Get(m) == b {
				moves = append(moves, m)
			} else { // different plot adjacent, add fence
				perim++
			}
		}
	}
	return area * perim
}

func solve1(reader *bufio.Scanner) {
	g = readGrid(reader)
	seen = make(map[point]struct{})
	total := 0
	for p, b := range g.Coords() {
		if _, ok := seen[p]; ok {
			continue
		}
		cost := calcCost1(p, b)
		total += cost
	}
	fmt.Println(total)
}

type pointdir struct {
	p point
	d dir
}

func countEdges(perim map[dir][]point) int {
	edges := 0
	for d, points := range perim {
		seenPerim := make(map[point]struct{})
		// fmt.Println("dir", d)
		for _, p := range points {
			// fmt.Println("seen", slices.AppendSeq([]point{}, maps.Keys(seenPerim)))
			if _, ok := seenPerim[p]; ok {
				continue
			}
			// fmt.Println("counting edge", d, p)
			seenPerim[p] = struct{}{} // mark seen
			edges++
			var upLeft, downRight point
			if d == UP || d == DOWN { // original motion up/down, so left/right to follow
				upLeft.x = -1
				downRight.x = 1
			} else {
				upLeft.y = -1
				downRight.y = 1
			}
			// follow up or left
			next := p.Add(upLeft)
			for slices.Contains(points, next) {
				seenPerim[next] = struct{}{}
				next = next.Add(upLeft)
			}
			next = p.Add(downRight)
			for slices.Contains(points, next) {
				seenPerim[next] = struct{}{}
				next = next.Add(downRight)
			}
		}
	}
	return edges
}

func calcCost2(start point, b byte) int {
	moves := []point{start}
	var p point
	var area int
	perim := make(map[dir][]point, 0)
	// interior := make(map[point]struct{}, 0)
	for len(moves) > 0 { // DFS over region
		p, moves = moves[len(moves)-1], moves[:len(moves)-1]
		if _, ok := seen[p]; ok {
			continue
		}
		// interior[p] = struct{}{}
		seen[p] = struct{}{}
		area++
		for _, m := range g.AllMoves(p) {
			if g.inBounds(m) && g.Get(m) == b {
				moves = append(moves, m)
			} else { // different plot adjacent, add fence
				diff := m.Sub(p)
				var d dir
				if diff.x == 0 { // up or down
					if diff.y > 0 {
						d = DOWN
					} else {
						d = UP
					}
				} else {
					if diff.x > 0 {
						d = RIGHT
					} else {
						d = LEFT
					}
				}
				perim[d] = append(perim[d], m)
			}
		}
	}
	numEdges := countEdges(perim)
	// fmt.Println("numb edges", b, numEdges)
	return area * numEdges
}

func solve2(reader *bufio.Scanner) {
	g = readGrid(reader)
	seen = make(map[point]struct{})
	total := 0
	for p, b := range g.Coords() {
		if _, ok := seen[p]; ok {
			continue
		}
		cost := calcCost2(p, b)
		total += cost
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
