package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int %v", err)
	}
	return i
}

type Counter map[int]int

func (c Counter) Add(i int) {
	count, _ := c[i]
	c[i] = count + 1
}

func (c Counter) GetCount(i int) int {
	count, _ := c[i]
	return count
}

func solve1(reader *bufio.Scanner) {
	h1 := &IntHeap{}
	h2 := &IntHeap{}
	whitespace, _ := regexp.Compile(" +")
	for reader.Scan() {
		line := whitespace.Split(reader.Text(), 2)
		heap.Push(h1, toInt(line[0]))
		heap.Push(h2, toInt(line[1]))
	}
	sum := 0
	for h1.Len() > 0 {
		i1 := heap.Pop(h1).(int)
		i2 := heap.Pop(h2).(int)
		diff := i1 - i2
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	fmt.Println(sum)
}

func solve2(reader *bufio.Scanner) {
	list := make([]int, 0)
	counter := Counter{}
	whitespace, _ := regexp.Compile(" +")
	for reader.Scan() {
		line := whitespace.Split(reader.Text(), 2)
		list = append(list, toInt(line[0]))
		counter.Add(toInt(line[1]))
	}
	sum := 0
	for _, i := range list {
		sum += i * counter.GetCount(i)
	}
	fmt.Println(sum)
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
