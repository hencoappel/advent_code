package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
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

func readDisk(line []byte) []int {
	total := 0
	for _, b := range line {
		total += int(b - '0')
	}
	disk := make([]int, total)
	file := true
	var id int = 0
	i := 0
	for _, b := range line {
		size := int(b - '0')
		for j := 0; j < size; j++ {
			if file {
				disk[i] = id
			} else {
				disk[i] = -1
			}
			i++
		}
		if file {
			id++
		}
		file = !file
	}
	return disk
}

func solve1(reader *bufio.Scanner) {
	reader.Scan()
	disk := readDisk(reader.Bytes())
	fmt.Println(disk)
	s, e := 0, len(disk)-1
	for s != e {
		for disk[s] != -1 && s != e {
			s++
		}
		for disk[e] == -1 && s != e {
			e--
		}
		disk[s] = disk[e]
		disk[e] = -1
	}
	fmt.Println(disk)
	sum := 0
	for i, v := range disk {
		if v == -1 {
			break
		}
		sum += i * v
	}
	fmt.Println(sum)
}

type file struct {
	start int
	size  int
}

func readFiles(line []byte) ([]*file, []*file) {
	spaces := make([]*file, 0, int(math.Floor(float64(len(line)/2.0))))
	files := make([]*file, 0, int(math.Ceil(float64(len(line)/2.0))))
	isfile := true
	start := 0
	for _, b := range line {
		size := int(b - '0')
		if isfile {
			files = append(files, &file{start, size})
		} else {
			spaces = append(spaces, &file{start, size})
		}
		start += size
		isfile = !isfile
	}
	return files, spaces
}

func solve2(reader *bufio.Scanner) {
	reader.Scan()
	files, spaces := readFiles(reader.Bytes())
	// for i, f := range files {
	// 	fmt.Println(i, f)
	// }
	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		// shifted := false
		for _, s := range spaces {
			if f.size <= s.size && s.start < f.start {
				// fmt.Printf("shifting %d %v to space %v\n", i, f, s)
				s.size -= f.size
				f.start = s.start
				s.start += f.size
				// fmt.Printf("new start/size: file %v - space %v\n", f, s)
				// shifted = true
				break
			}
		}
		// if !shifted {
		// 	fmt.Printf("No spaces for file %d %v\n", i, f)
		// }
	}
	sum := 0
	sum2 := 0
	for i, f := range files {
		// fmt.Println(i, f)
		sum += f.size * (f.start + (f.start + f.size - 1)) / 2 * i
		for j := range f.size {
			sum2 += (f.start + j) * i
		}
	}
	// printdisk(files)
	fmt.Println(sum)
	fmt.Println(sum2)
}

func printdisk(files []*file) {
	idmap := make(map[int]int)
	for i, f := range files {
		idmap[f.start] = i
	}
	files = slices.SortedFunc(slices.Values(files), func(f1, f2 *file) int {
		return f1.start - f2.start
	})

	idx := 0
	for _, f := range files {
		for range f.start - idx {
			fmt.Print(".")
			idx++
		}
		for range f.size {
			fmt.Print(idmap[f.start])
			fmt.Print(",")
			idx++
		}
	}
	fmt.Println()
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
