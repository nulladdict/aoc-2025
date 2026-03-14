package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	str := strings.TrimSpace(string(input))
	ranges := make([]Range, 0, 100)
	for r := range strings.SplitSeq(str, ",") {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})
	}
	fmt.Println("part1", part1(ranges))
	fmt.Println("part2", part2(ranges))
}

type Range struct {
	start int
	end   int
}

func part1(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			str := strconv.Itoa(id)
			if len(str)%2 == 1 {
				continue
			}
			mid := len(str) / 2
			if str[:mid] == (str[mid:]) {
				sum += id
			}
		}
	}
	return sum
}

func part2(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			str := strconv.Itoa(id)
			for factor := 1; factor <= len(str)/2; factor++ {
				if check(str, factor) {
					sum += id
					break
				}
			}
		}
	}
	return sum
}

func check(str string, factor int) bool {
	if len(str)%factor != 0 {
		return false
	}
	pattern := str[:factor]
	for k := 0; k < len(str); k += factor {
		if str[k:k+factor] != pattern {
			return false
		}
	}
	return true
}
