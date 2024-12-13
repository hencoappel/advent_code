package main

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"maps"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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
	strs := strings.Split(line, ",")
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

type set map[int]struct{}

func NewSet(ints ...int) set {
	s := set{}
	s.AddAll(slices.Values(ints))
	return s
}
func (s set) Add(i int) {
	s[i] = struct{}{}
}
func (s set) AddAll(ints iter.Seq[int]) {
	for i := range ints {
		s.Add(i)
	}
}
func (s set) Contains(i int) bool {
	_, ok := s[i]
	return ok
}
func (s set) ContainsAny(ints iter.Seq[int]) bool {
	for i := range ints {
		if s.Contains(i) {
			return true
		}
	}
	return false
}
func (s set) All() iter.Seq[int] {
	return maps.Keys(s)
}
func (s set) String() string {
	str := "{"
	first := true
	for i := range s {
		if first {
			str += strconv.Itoa(i)
			first = false
			continue
		}
		str += "," + strconv.Itoa(i)
	}
	str += "}"
	return str
}

type updates struct {
	updates [][]int
	after   map[int]set
}

func readUpdates(reader *bufio.Scanner) *updates {
	u := &updates{
		updates: make([][]int, 0),
		after:   make(map[int]set),
	}
	for reader.Scan() {
		txt := reader.Text()
		if txt == "" {
			break
		}
		vals := strings.Split(txt, "|")
		b, a := toInt(vals[0]), toInt(vals[1])
		s, ok := u.after[b]
		if !ok {
			s = NewSet()
		}
		s.Add(a)
		u.after[b] = s
	}

	for reader.Scan() {
		u.updates = append(u.updates, readIntLine(reader.Text()))
	}

	for k, v := range u.after {
		fmt.Println(k, v)
	}
	fmt.Println()
	return u
}

func updateValid(updateInfo *updates, update []int) bool {
	after := NewSet()
	for i := len(update) - 1; i >= 0; i-- {
		v := update[i]
		for a := range after.All() {
			// no number after v should contain v
			if updateInfo.after[a].Contains(v) {
				fmt.Println(update)
				fmt.Println("invalid: checking", v, a, "after for a contains v")
				fmt.Println()
				return false
			}
		}
		after.Add(v)
	}
	return true
}

func solve1(reader *bufio.Scanner) {
	updateInfo := readUpdates(reader)
	total := 0
	for _, u := range updateInfo.updates {
		if updateValid(updateInfo, u) {
			mid := int(math.Floor(float64(len(u)) / 2.0))
			fmt.Println("valid", u[mid], u)
			total += u[mid]
		}
	}
	fmt.Println("total", total)
}

func sortUpdate(updateInfo *updates, update []int) []int {
	newUpdate := slices.Clone(update)
	slices.SortStableFunc(newUpdate, func(a, b int) int {
		// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
		if updateInfo.after[a].Contains(b) {
			return -1
		}
		if updateInfo.after[b].Contains(a) {
			return 1
		}
		return 0
	})
	return newUpdate
}

func solve2(reader *bufio.Scanner) {
	updateInfo := readUpdates(reader)
	total := 0
	for _, u := range updateInfo.updates {
		if !updateValid(updateInfo, u) {
			newUpdate := sortUpdate(updateInfo, u)
			mid := int(math.Floor(float64(len(newUpdate)) / 2.0))
			fmt.Println("sorted invalid", newUpdate[mid], newUpdate)
			total += newUpdate[mid]
		}
	}
	fmt.Println("total", total)
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
