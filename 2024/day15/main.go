package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func sumBoxCoords(g grid) int {
	sum := 0
	for y, row := range g {
		for x, b := range row {
			if b == 'O' {
				sum += x + 100*y
			}
		}
	}
	return sum
}

func motion(m byte) point {
	switch m {
	case '<': // left
		return point{-1, 0}
	case '>': // right
		return point{1, 0}
	case '^': // up
		return point{0, -1}
	case 'v': // down
		return point{0, 1}
	}
	panic("invalid motion")
}

func nextEmpty(g grid, start point, dir point) *point {
	p := start
	for {
		p = p.Add(dir)
		b := g.Get(p)
		if b == '#' {
			return nil
		}
		if g.Get(p) == '.' {
			return &p
		}
	}
}

func step(g grid, start point, moves []byte) {
	p := start
	for _, b := range moves {
		m := motion(b)
		empty := nextEmpty(g, p, m)
		if empty == nil { // can't move
			// fmt.Println("can't move")
			continue
		}
		next := p.Add(m)
		// fmt.Println("moving to", next)
		g.Set(p, '.')
		g.Set(next, '@')
		if *empty != next {
			g.Set(*empty, 'O')
		}
		p = next
		// g.Print()
	}
}

func readMoves(reader *bufio.Scanner) []byte {
	b := make([]byte, 0)
	for reader.Scan() {
		b = append(b, reader.Bytes()...)
	}
	return b
}

func solve1(reader *bufio.Scanner) {
	g := readGrid(reader)
	var start point
	for y, row := range g {
		for x, b := range row {
			if b == '@' {
				start = point{x, y}
			}
		}
	}
	// fmt.Println(start)
	moves := readMoves(reader)
	// g.Print()
	step(g, start, moves)
	fmt.Println(sumBoxCoords(g))
}

func doubleWidth(g grid) plan {
	boxes := make([]point, 0)
	walls := make([]point, 0)
	var start point
	for y, row := range g {
		for x, b := range row {
			switch b {
			case '#':
				walls = append(walls, point{x * 2, y}, point{x*2 + 1, y})
			case 'O':
				boxes = append(boxes, point{x * 2, y})
			case '@':
				start = point{x * 2, y}
			}
		}
	}
	return plan{boxes, walls, start}
}

type plan struct {
	boxes []point
	wall  []point
	pos   point
}

func (p plan) boxAt(pos point) *point {
	if slices.Contains(p.boxes, pos) {
		return &pos
	}
	posLeft := pos.Add(point{-1, 0})
	if slices.Contains(p.boxes, posLeft) {
		return &posLeft
	}
	return nil
}

func (p plan) canPush(pos, dir point) bool {
	if slices.Contains(p.wall, pos) {
		return false
	}
	b := p.boxAt(pos)
	if b == nil {
		return true // empty spot
	}
	next := (*b).Add(dir)
	nextRight := next.Add(point{1, 0})
	if dir.x == 0 { // up/down
		return p.canPush(next, dir) && p.canPush(nextRight, dir)
	} else if dir.x == -1 { // left
		return p.canPush(next, dir)
	} else { // right
		return p.canPush(nextRight, dir)
	}
}

func (p plan) print() {
	maxx, maxy := 0, 0
	for _, w := range p.wall {
		if w.x > maxx {
			maxx = w.x
		}
		if w.y > maxy {
			maxy = w.y
		}
	}
	g := make([][]byte, maxy+1)
	for i := range len(g) {
		g[i] = make([]byte, maxx+1)
		for j := range len(g[i]) {
			g[i][j] = '.'
		}
	}
	for _, w := range p.wall {
		g[w.y][w.x] = '#'
	}
	for _, b := range p.boxes {
		g[b.y][b.x] = '['
		g[b.y][b.x+1] = ']'
	}
	g[p.pos.y][p.pos.x] = '@'
	for _, row := range g {
		fmt.Println(string(row))
	}
	fmt.Println(strings.Repeat("-", len(g[0])))
}

func (p plan) getBoxesToPush(pos, dir point) []int {
	fmt.Println("getBoxesToPush", pos, dir)
	// fmt.Println(pos, dir)
	b := p.boxAt(pos)
	if b == nil {
		fmt.Println("no boxes")
		return []int{}
	}
	fmt.Println("found box", *b)
	res := []int{slices.Index(p.boxes, *b)}
	next := (*b).Add(dir)
	if dir.x == 0 { // up/down - check left and right
		res = append(res, p.getBoxesToPush(next, dir)...)
		res = append(res, p.getBoxesToPush(next.Add(point{1, 0}), dir)...)
	} else if dir.x == 1 { // right - shift 2
		next = next.Add(dir)
		res = append(res, p.getBoxesToPush(next, dir)...)
	} else { // left - just check left
		res = append(res, p.getBoxesToPush(next, dir)...)
	}
	return res
}

func (p plan) push(next, dir point) {
	topush := p.getBoxesToPush(next, dir)
	slices.Sort(topush)
	topush = slices.Compact(topush)
	for _, i := range topush {
		p.boxes[i] = p.boxes[i].Add(dir)
	}
}

func (p plan) step(moves []byte) {
	// p.print()
	for _, b := range moves {
		m := motion(b)
		next := p.pos.Add(m)
		if p.canPush(next, m) {
			p.push(next, m)
			p.pos = next
		}
		// p.print()
	}
}

func solve2(reader *bufio.Scanner) {
	g := readGrid(reader)
	// g.Print()
	fmt.Println(len(g))
	plan := doubleWidth(g)
	moves := readMoves(reader)
	// fmt.Println(string(moves))
	plan.step(moves)
	sum := 0
	for _, b := range plan.boxes {
		sum += b.x + 100*b.y
	}
	fmt.Println("su", sum)
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
