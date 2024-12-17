package main

import (
	"bufio"
	"fmt"
	"os"
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

func doubleWidth(g grid) ([]point, []point, point) {
	boxes := make([]point, 0)
	walls := make([]point, 0)
	var start point
	for y, row := range g {
		for x, b := range row {
			var b1, b2 byte
			switch b {
			case '#':
				walls = append(walls, point{x * 2, y}, point{x*2 + 1, y})
			case 'O':
				boxes = append(boxes, point{x * 2, y})
			case '@':
				start = point{x * 2, y}
		}
	}
	return boxes, walls, start
}

func step2(g grid, start point, moves []byte) {
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

func solve2(reader *bufio.Scanner) {
	g := readGrid(reader)
	g.Print()
	g = doubleWidth(g)
	var start point
	for y, row := range g {
		for x, b := range row {
			if b == '@' {
				start = point{x, y}
			}
		}
	}
	g.Print()
	fmt.Println(start)
	// moves := readMoves(reader)
	// fmt.Println(string(moves))
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
