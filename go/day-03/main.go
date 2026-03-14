package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	in, _ := os.ReadFile("input.txt")
	str := strings.TrimSpace(string(in))
	banks := strings.Split(str, "\n")
	fmt.Println("part1", part1(banks))
	fmt.Println("part2", part2(banks))
}

func part1(banks []string) int {
	sum := 0
	cache := make(map[SolvedFor]int)
	for _, bank := range banks {
		sum += maximum(bank, 2, cache)
	}
	return sum
}

func part2(banks []string) int {
	sum := 0
	cache := make(map[SolvedFor]int)
	for _, bank := range banks {
		sum += maximum(bank, 12, cache)
	}
	return sum
}

type SolvedFor struct {
	bank  string
	count int
}

func maximum(bank string, count int, cache map[SolvedFor]int) int {
	if count == 0 {
		return 0
	}
	if val, ok := cache[SolvedFor{bank, count}]; ok {
		return val
	}
	max := 0
	for i := 0; i < len(bank)-count+1; i += 1 {
		left := int(bank[i] - '0')
		right := maximum(bank[i+1:], count-1, cache)
		joltage := left*int(math.Pow10(count-1)) + right
		if joltage >= max {
			max = joltage
		}
	}
	cache[SolvedFor{bank, count}] = max
	return max
}
