package main

import (
	"bufio"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
)

type state struct {
	p point
	d dir
}

func requiredTurns(d1, d2 dir) int {
	if d1 == d2 {
		return 0
	}
	if (d1 == UP && d2 == DOWN) ||
		(d1 == DOWN && d2 == UP) ||
		(d1 == LEFT && d2 == RIGHT) ||
		(d1 == RIGHT && d2 == LEFT) {
		return 2
	}
	return 1
}

func bfs(g grid, start, end point, d dir) (int, map[state]int) {
	queue := []state{{start, d}}
	visited := make(map[state]int)
	visited[state{start, d}] = 0
	for len(queue) > 0 {
		p := queue[0]
		// fmt.Println(p)
		queue = queue[1:]
		currentScore := visited[p]
		for _, d := range dirs {
			next := p.p.Move(d)
			turns := requiredTurns(p.d, d)
			if turns == 2 {
				continue
			}
			nextScore := currentScore + turns*1000 + 1
			nextState := state{next, d}
			if g.Get(next) == '#' {
				continue
			}
			s, ok := visited[nextState]
			if !ok || nextScore < s {
				visited[nextState] = nextScore
				queue = append(queue, nextState)
			}
		}
	}
	// fmt.Println(visited)
	minScore := math.MaxInt
	for _, d := range dirs {
		if s, ok := visited[state{end, d}]; ok {
			minScore = minInt(minScore, s)
		}
	}
	return minScore, visited
}

func solve1(reader *bufio.Scanner) {
	g := readGrid(reader)
	var s, e point
	for p, b := range g.Coords() {
		if b == 'S' {
			s = p
		} else if b == 'E' {
			e = p
		}
	}
	g.Print()
	// fmt.Println(s, e)
	score, _ := bfs(g, s, e, RIGHT)
	fmt.Println(score)
}

type route struct {
	r     []point
	turns int
}

func dfsRec(g grid, r route, currentDir dir) []route {
	pos := r.r[len(r.r)-1]
	// fmt.Println(pos)
	routes := make([]route, 0)
	for _, d := range dirs {
		turns := requiredTurns(currentDir, d)
		if turns == 2 {
			continue
		}
		next := pos.Move(d)
		if g.Get(next) == '#' {
			continue
		}
		for _, p := range r.r {
			if p == next {
				continue
			}
		}
		newRoute := route{append(slices.Clone(r.r), next), r.turns + turns}
		if g.Get(next) == 'E' {
			fmt.Println("Found end")
			return []route{newRoute}
		}
		routes = append(routes, dfsRec(g, newRoute, d)...)
	}
	return routes
}

func dfs(g grid, start, end point, d dir) int {
	r := route{[]point{start}, 0}
	routes := dfsRec(g, r, d)
	minScore := math.MaxInt
	for _, r := range routes {
		minScore = minInt(minScore, r.turns*1000+len(r.r))
	}
	return minScore
}

func states(p point) []state {
	return []state{
		{p, UP},
		{p, DOWN},
		{p, LEFT},
		{p, RIGHT},
	}
}

func (p point) MoveBack(d dir) point {
	switch d {
	case UP:
		return p.Move(DOWN)
	case DOWN:
		return p.Move(UP)
	case LEFT:
		return p.Move(RIGHT)
	case RIGHT:
		return p.Move(LEFT)
	}
	return p
}

func backtrackRoutes(visited map[state]int, end point, minScore int) []point {
	points := make(map[point]struct{}, 0)
	queue := make([]state, 0)
	for _, s := range states(end) {
		if score, ok := visited[s]; ok && score == minScore {
			queue = append(queue, s)
		}
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		points[p.p] = struct{}{}
		for _, d := range dirs {
			next := state{p.p.MoveBack(d), d}
			if score, ok := visited[next]; ok && score == visited[p]-1 {
				queue = append(queue, next)
			}
		}
	}
	return slices.Collect(maps.Keys(points))
}

func solve2(reader *bufio.Scanner) {
	g := readGrid(reader)
	var s, e point
	for p, b := range g.Coords() {
		if b == 'S' {
			s = p
		} else if b == 'E' {
			e = p
		}
	}
	g.Print()
	// fmt.Println(s, e)
	score, visited := bfs(g, s, e, RIGHT)
	points := backtrackRoutes(visited, e, score)
	fmt.Println(len(points))
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
