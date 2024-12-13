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

func (p pair) sub(p2 pair) pair {
	return pair{p._0 - p2._0, p._1 - p2._1}
}

func (p pair) add(p2 pair) pair {
	return pair{p._0 + p2._0, p._1 + p2._1}
}

func readAntenna(reader *bufio.Scanner) (map[byte][]pair, pair) {
	x, y := 0, 0
	antenna := make(map[byte][]pair)
	for reader.Scan() {
		line := reader.Bytes()
		x = len(line)
		for i, b := range line {
			if b != '.' {
				var list []pair
				if l, ok := antenna[b]; ok {
					list = l
				}
				antenna[b] = append(list, pair{i, y})
			}
		}
		y += 1
	}
	return antenna, pair{x, y}
}

func solve1(reader *bufio.Scanner) {
	antenna, size := readAntenna(reader)
	// fmt.Println(antenna)
	antis := make(map[pair]struct{})
	for _, list := range antenna {
		for _, a1 := range list {
			for _, a2 := range list {
				if a1 == a2 {
					continue
				}
				diff := a2.sub(a1)
				ant1 := a1.sub(diff)
				ant2 := a2.add(diff)
				if ant1._0 >= 0 && ant1._1 >= 0 && ant1._0 < size._0 && ant1._1 < size._1 {
					antis[ant1] = struct{}{}
				}
				if ant2._0 >= 0 && ant2._1 >= 0 && ant2._0 < size._0 && ant2._1 < size._1 {
					antis[ant2] = struct{}{}
				}
			}
		}
	}
	// fmt.Println(antis)
	fmt.Println(len(antis))
}

func solve2(reader *bufio.Scanner) {
	antenna, size := readAntenna(reader)
	// fmt.Println(antenna)
	antis := make(map[pair]struct{})
	for _, list := range antenna {
		for _, a1 := range list {
			for _, a2 := range list {
				if a1 == a2 {
					continue
				}
				antis[a1] = struct{}{}
				antis[a2] = struct{}{}
				diff := a2.sub(a1)
				ant1 := a1.sub(diff)
				for ant1._0 >= 0 && ant1._1 >= 0 && ant1._0 < size._0 && ant1._1 < size._1 {
					antis[ant1] = struct{}{}
					ant1 = ant1.sub(diff)
				}
				ant2 := a2.add(diff)
				for ant2._0 >= 0 && ant2._1 >= 0 && ant2._0 < size._0 && ant2._1 < size._1 {
					antis[ant2] = struct{}{}
					ant2 = ant2.add(diff)
				}
			}
		}
	}
	// fmt.Println(antis)
	fmt.Println(len(antis))
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
