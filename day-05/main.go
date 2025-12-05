package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	parts := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
	ranges := make([]Range, 0)
	for r := range strings.SplitSeq(parts[0], "\n") {
		ranges = append(ranges, parse(r))
	}
	ids := make([]int, 0)
	for id := range strings.SplitSeq(parts[1], "\n") {
		x, _ := strconv.Atoi(id)
		ids = append(ids, x)
	}
	fmt.Println("part1", part1(ranges, ids))
	fmt.Println("part2", part2(ranges))
}

type Range struct {
	low  int
	high int
}

func parse(s string) Range {
	var r Range
	fmt.Sscanf(s, "%d-%d", &r.low, &r.high)
	return r
}

func isValid(ranges []Range, id int) bool {
	for _, r := range ranges {
		if id >= r.low && id <= r.high {
			return true
		}
	}
	return false
}

func part1(ranges []Range, ids []int) int {
	count := 0
	for _, id := range ids {
		if isValid(ranges, id) {
			count++
		}
	}
	return count
}

func (a Range) compare(b Range) int {
	switch {
	case a.low < b.low:
		return -1
	case a.low > b.low:
		return 1
	default:
		switch {
		case a.high < b.high:
			return -1
		case a.high > b.high:
			return 1
		default:
			return 0
		}
	}
}

func part2(ranges []Range) int {
	slices.SortFunc(ranges, func(a, b Range) int { return a.compare(b) })
	sum := 0
	last := 0
	for _, r := range ranges {
		start := max(r.low, last+1)
		end := max(r.high, last)
		sum += end - start + 1
		last = max(last, r.high)
	}
	return sum
}
