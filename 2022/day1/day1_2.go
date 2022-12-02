package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
)

// An IntHeap is a min-heap of ints.
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

func (h *IntHeap) pop() int {
	return heap.Pop(h).(int)
}

func solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	h := &IntHeap{}
	heap.Init(h)
	current := 0
	for scanner.Scan() {
		var maxVals []int = *h
		fmt.Printf("heap: %v\n", maxVals)
		l := scanner.Text()
		if l == "" {
			heap.Push(h, current)
			if h.Len() > 3 {
				heap.Pop(h)
			}
			current = 0
			continue
		}
		val, _ := strconv.Atoi(l)
		current += val
	}
	var sum int = h.pop() + h.pop() + h.pop()
	fmt.Printf("highest: %d\n", sum)
}

func main() {
	f, err := os.Open("part1.in")
	defer f.Close()
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	reader := bufio.NewReader(f)
	solve(reader)
}
