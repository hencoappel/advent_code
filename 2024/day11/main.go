package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
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

type pair [2]int

func trimZeros(stone string) string {
	if stone[0] != '0' {
		return stone
	}
	for i, b := range stone {
		if b != '0' {
			return stone[i:]
		}
	}
	return "0"
}
func blink(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}
	if len(stone)%2 == 0 {
		mid := len(stone) / 2
		l, r := trimZeros(stone[0:mid]), trimZeros(stone[mid:])
		return []string{l, r}
	} else {
		s, _ := strconv.Atoi(stone)
		return []string{strconv.Itoa(s * 2024)}
	}
}

func blinkN(stones []string, blinks int) int {
	for i := range blinks {
		fmt.Println(i)
		newStones := make([]string, 0, int(float64(len(stones))*1.5))
		for _, stone := range stones {
			newStones = append(newStones, blink(stone)...)
		}
		stones = newStones
		// fmt.Println(stones)
	}
	return len(stones)
}

func solve1(reader *bufio.Scanner) {
	reader.Scan()
	stones := strings.Split(reader.Text(), " ")
	fmt.Println(blinkN(stones, 25))
}

func numDigits(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func numDigitsFast(i int) int {
	if i >= 1e18 {
		return 19
	}
	//   15 = even, /10=1   /100=0
	//  150 = odd,  /10=15  /100=1
	// 1500 = even, /10=150 /100=15
	x, count := i, 0
	for x >= 10 {
		x /= 100
		count += 2
	}
	if x != 0 {
		count++
	}
	return count
}

func blinkInt(stone int) (int, int) {
	if stone == 0 {
		return 0, -1
	}
	digits := numDigits(stone)
	// digits := numDigitsFast(stone)
	if digits&1 == 0 {
		mid := digits / 2
		l := stone / int(math.Pow10(mid))
		r := stone - l*int(math.Pow10(mid))
		return l, r
	} else {
		return stone * 2024, -1
	}
}

func blinkNInts(stones []int, blinks int) int {
	for i := range blinks {
		fmt.Println(i)
		newStones := make([]int, 0, len(stones)*2)
		for _, stone := range stones {
			l, r := blinkInt(stone)
			newStones = append(newStones, l)
			if r != -1 {
				newStones = append(newStones, r)
			}
		}
		stones = newStones
		// fmt.Println(stones)
	}
	return len(stones)
}

func blinkNInline(stones []int, blinks int) int {
	for range blinks {
		// fmt.Println(i)
		newStones := make([]int, 0, len(stones)*2)
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else {
				digits := numDigits(stone)
				if digits&1 == 0 {
					mid := digits / 2
					l := stone / int(math.Pow10(mid))
					r := stone - l*int(math.Pow10(mid))
					newStones = append(newStones, l, r)
				} else {
					newStones = append(newStones, stone*2024)
				}
			}
		}
		stones = newStones
	}
	// fmt.Println(stones)
	return len(stones)
}

type stack[E any] struct {
	s   []E
	len int
}

func (s *stack[E]) peak() E {
	return s.s[s.len-1]
}
func (s *stack[E]) pushReplace(i E) {
	s.s[s.len-1] = i
}
func (s *stack[E]) pushReplace2(i, i2 E) {
	s.s[s.len-1] = i
	s.push(i2)
}
func (s *stack[E]) reserve1() {
	var e E
	// if s.len < len(s.s) {
	s.s[s.len] = e
	// } else {
	// 	s.s = append(s.s, e)
	// }
	s.len++
}
func (s *stack[E]) push(i E) {
	// if s.len < len(s.s) {
	s.s[s.len] = i
	// } else {
	// 	fmt.Println("Appending")
	// 	s.s = append(s.s, i)
	// }
	s.len++
}
func (s *stack[E]) pop() E {
	s.len--
	return s.s[s.len]
}

func blinkNLowMem(stones []int, blinks int) int {
	stoneCounts := make([]stone, len(stones))
	for i, s := range stones {
		stoneCounts[i] = stone{0, s}
	}
	s := &stack[stone]{stoneCounts, len(stoneCounts)}
	count := 0
	for s.len > 0 {
		currentStone := s.peak()
		// currentStone := &s.s[s.len-1]
		if currentStone.blink == blinks {
			s.pop()
			count++
		} else if currentStone.val == 0 {
			currentStone.val = 1
			currentStone.blink++
			s.pushReplace(currentStone)
		} else {
			digits := numDigits(currentStone.val)
			if digits&1 == 0 {
				mid := digits / 2
				l := currentStone.val / int(math.Pow10(mid))
				r := currentStone.val - l*int(math.Pow10(mid))
				s.pushReplace2(stone{currentStone.blink + 1, l}, stone{currentStone.blink + 1, r})
			} else {
				currentStone.val = currentStone.val * 2024
				currentStone.blink++
				s.pushReplace(currentStone)
			}
		}
	}
	// fmt.Println(stones)
	return count
}

var blinkmultimap = [][][]int{
	{ // 0
		{1},
		{2024},
		{20, 24},
		{2, 0, 2, 4},
	},
	{ // 1
		{2024},
		{20, 24},
		{2, 0, 2, 4},
		{4048, 1, 4048, 8096},
	},
	{ //2
		{4048},
		{40, 48},
		{4, 0, 4, 8},
		{8096, 1, 8096, 16192},
	},
	{ //3
		{6072},
		{60, 72},
		{6, 0, 7, 2},
		{12144, 1, 14168, 4048},
	},
	{ //4
		{8096},
		{80, 96},
		{8, 0, 9, 6},
		{16192, 1, 18216, 12144},
	},
}

var blinksplitmap = map[int][]int{
	20:   {2, 0},
	24:   {2, 4},
	40:   {4, 0},
	48:   {4, 8},
	60:   {6, 0},
	72:   {7, 2},
	2024: {20, 24},
	4048: {40, 48},
	6072: {60, 72},
	8096: {80, 96},
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func blinkNLowMem2(initialStone int, blinks int) int {
	stoneCounts := make([]stone, blinks*3)
	stoneCounts[0] = stone{0, initialStone}
	s := &stack[stone]{stoneCounts, 1}
	count := 0
	for s.len > 0 {
		currentStone := s.peak()
		// currentStone := &s.s[s.len-1]
		if currentStone.blink == blinks {
			s.pop()
			count++
			continue
		} else if currentStone.val <= 4 {
			b := min(blinks-currentStone.blink, 4) // blinks we're skipping
			// fmt.Println(currentStone.val, b-1)
			vals := blinkmultimap[currentStone.val][b-1]
			nb := currentStone.blink + b
			s.s[s.len-1].val = vals[0]
			s.s[s.len-1].blink = nb
			for i := 1; i < len(vals); i++ {
				s.reserve1()
				s.s[s.len-1].val = vals[i]
				s.s[s.len-1].blink = nb
			}
		} else {
			digits := numDigits(currentStone.val)
			nb := currentStone.blink + 1
			if digits&1 == 0 {
				midScale := int(math.Pow10(digits / 2))
				l := currentStone.val / midScale
				r := currentStone.val - l*midScale
				s.s[s.len-1].val = l
				s.s[s.len-1].blink = nb
				s.reserve1()
				s.s[s.len-1].val = r
				s.s[s.len-1].blink = nb
			} else {
				s.s[s.len-1].val *= 2024
				s.s[s.len-1].blink = nb
			}
		}
	}
	// fmt.Println(stones)
	return count
}

func blinkNLowMem3(initialStone int, blinks int) int {
	stoneCounts := make([]stone, 1)
	stoneCounts[0] = stone{0, initialStone}
	s := &stack[stone]{stoneCounts, len(stoneCounts)}
	count := 0
	for s.len > 0 {
		currentStone := s.peak()
		// currentStone := &s.s[s.len-1]
		if currentStone.blink == blinks {
			s.pop()
			count++
			continue
		} else if currentStone.val <= 4 {
			b := min(blinks-currentStone.blink, 4) // blinks we're skipping
			// fmt.Println(currentStone.val, b-1)
			vals := blinkmultimap[currentStone.val][b-1]
			nb := currentStone.blink + b
			s.s[s.len-1].val = vals[0]
			s.s[s.len-1].blink = nb
			for i := 1; i < len(vals); i++ {
				s.push(stone{nb, vals[i]})
			}
			// } else if vals, ok := blinksplitmap[currentStone.val]; ok {
			// 	nb := currentStone.blink + 1
			// 	s.s[s.len-1].val = vals[0]
			// 	s.s[s.len-1].blink = nb
			// 	s.reserve1()
			// 	s.s[s.len-1].val = vals[1]
			// 	s.s[s.len-1].blink = nb
		} else {
			digits := numDigits(currentStone.val)
			if digits&1 == 0 {
				midScale := int(math.Pow10(digits / 2))
				l := currentStone.val / midScale
				r := currentStone.val - l*midScale
				// s.pushReplace2(stone{currentStone.blink + 1, l}, stone{currentStone.blink + 1, r})
				nb := currentStone.blink + 1
				s.s[s.len-1].val = l
				s.s[s.len-1].blink = nb
				// s.push(stone{currentStone.blink + 1, r})
				s.reserve1()
				s.s[s.len-1].val = r
				s.s[s.len-1].blink = nb
			} else {
				s.s[s.len-1].val *= 2024
				s.s[s.len-1].blink++
				// currentStone.val = currentStone.val * 2024
				// currentStone.blink++
				// s.pushReplace(currentStone)
			}
		}
	}
	// fmt.Println(stones)
	return count
}

type stone struct {
	blink int
	val   int
}

func blinkNParallelChan(stones []int, blinks int) int {
	res := make([]int, len(stones))
	var wg sync.WaitGroup
	toblink := make(chan stone)
	// done := make(chan struct{})
	go func() {
		for _, s := range stones {
			toblink <- stone{0, s}
		}
	}()
	for range 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			timeout := 2 * time.Second
			t := time.NewTimer(timeout)
			for {
				select {
				case s := <-toblink:
					fmt.Println(s)
					// res[j] = blinkNInline([]int{stones[j]}, blinks)
					t.Reset(timeout)
				case <-t.C:
					break
				}
			}
		}()
	}
	wg.Wait()
	// fmt.Println(res)
	total := 0
	for _, i := range res {
		total += i
	}
	return total
}

type blinkFunc func(stone int, blinks int) int

func blinkNParallel(stones []int, blinks int, f blinkFunc) int {
	res := make([]int, len(stones))
	var wg sync.WaitGroup
	for i := range len(stones) {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			res[j] = f(stones[j], blinks)
		}()
	}
	wg.Wait()
	// fmt.Println(res)
	total := 0
	for _, i := range res {
		total += i
	}
	return total
}

func applyBlinkNInline(stones []int, blinks int) []int {
	for range blinks {
		// fmt.Println(i)
		newStones := make([]int, 0, len(stones)*2)
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else {
				digits := numDigits(stone)
				if digits&1 == 0 {
					mid := digits / 2
					l := stone / int(math.Pow10(mid))
					r := stone - l*int(math.Pow10(mid))
					newStones = append(newStones, l, r)
				} else {
					newStones = append(newStones, stone*2024)
				}
			}
		}
		stones = newStones
	}
	// fmt.Println(stones)
	return stones
}

func blinkNCountMap(initialStone int, blinks int) int {
	stoneCounts := make([]stone, blinks*3)
	stoneCounts[0] = stone{0, initialStone}
	s := &stack[stone]{stoneCounts, 1}
	count := 0
	for s.len > 0 {
		currentStone := s.peak()
		// currentStone := &s.s[s.len-1]
		if currentStone.blink == blinks {
			s.pop()
			count++
		} else if blinksToGo := blinks - currentStone.blink; currentStone.val < len(blinkcountmap) && blinksToGo < countMapLookAhead {
			s.pop()
			count += blinkcountmap[currentStone.val][blinksToGo-1]
		} else if currentStone.val <= 0 {
			b := min(blinks-currentStone.blink, 4) // blinks we're skipping
			// fmt.Println(currentStone.val, b-1)
			vals := blinkmultimap[currentStone.val][b-1]
			nb := currentStone.blink + b
			s.s[s.len-1].val = vals[0]
			s.s[s.len-1].blink = nb
			for i := 1; i < len(vals); i++ {
				s.reserve1()
				s.s[s.len-1].val = vals[i]
				s.s[s.len-1].blink = nb
			}
		} else {
			digits := numDigits(currentStone.val)
			nb := currentStone.blink + 1
			if digits&1 == 0 {
				midScale := int(math.Pow10(digits / 2))
				l := currentStone.val / midScale
				r := currentStone.val - l*midScale
				s.s[s.len-1].val = l
				s.s[s.len-1].blink = nb
				s.reserve1()
				s.s[s.len-1].val = r
				s.s[s.len-1].blink = nb
			} else {
				s.s[s.len-1].val *= 2024
				s.s[s.len-1].blink = nb
			}
		}
	}
	// fmt.Println(stones)
	return count
}

const countMapLookAhead = 37 // most optimal?
const countMapLen = 10

var blinkcountmap = make([][]int, countMapLen)

func initCountMap() {
	for i := range len(blinkcountmap) {
		blinkcountmap[i] = make([]int, countMapLookAhead-1)
		stones := []int{i}
		for j := range len(blinkcountmap[i]) {
			stones = applyBlinkNInline(stones, 1)
			blinkcountmap[i][j] = len(stones)
		}
	}
}

func solve2(reader *bufio.Scanner) {
	t := time.Now()
	initCountMap()
	reader.Scan()
	stones := readIntLine(reader.Text())
	n, _ := strconv.Atoi(os.Args[2])
	preapply := 3
	stones = applyBlinkNInline(stones, preapply)
	fmt.Println("init stones", len(stones))
	fmt.Println(blinkNParallel(stones, n-preapply, blinkNCountMap))
	t2 := time.Now()
	fmt.Println(t2.Sub(t))
}

func blinkNCache(cache map[stone]int, s stone) int {
	if s.blink == 0 {
		return 1
	}
	if c, ok := cache[s]; ok {
		return c
	}
	ns := stone{s.blink - 1, s.val}

	var count int
	if s.val == 0 {
		ns.val = 1
		count = blinkNCache(cache, ns)
	} else {
		digits := numDigits(s.val)
		if digits&1 == 0 {
			midScale := int(math.Pow10(digits / 2))
			l := s.val / midScale
			r := s.val - l*midScale
			ns.val = l
			count = blinkNCache(cache, ns)
			ns.val = r
			count += blinkNCache(cache, ns)
		} else {
			ns.val *= 2024
			count = blinkNCache(cache, ns)
		}
	}
	cache[s] = count
	return count
}

func solve2_after_looking_up_solutions(reader *bufio.Scanner) {
	t := time.Now()
	reader.Scan()
	stones := readIntLine(reader.Text())
	blinks, _ := strconv.Atoi(os.Args[2])
	total := 0
	cache := make(map[stone]int)
	for _, s := range stones {
		total += blinkNCache(cache, stone{blinks, s})
	}
	fmt.Println(total)
	t2 := time.Now()
	fmt.Println(t2.Sub(t))
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
	// solve2(scanner)
	solve2_after_looking_up_solutions(scanner)
}
