package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(bytes))
	lines := strings.Split(input, "\n")
	fmt.Println("part1", part1(lines))
	fmt.Println("part2", part2(lines))
}

type Point struct {
	x, y int
}

func grid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, c := range line {
			grid[i][j] = c
		}
	}
	return grid
}

func part1(lines []string) int {
	grid := grid(lines)
	count := 0
	idx := slices.Index(grid[0], 'S')
	start := Point{idx, 0}
	check := make(map[Point]bool, 0)
	check[start] = true
	for i := 0; i < len(grid)-1; i++ {
		next := make(map[Point]bool, 0)
		for p := range check {
			switch grid[p.y+1][p.x] {
			case '.':
				next[Point{p.x, p.y + 1}] = true
			case '^':
				count++
				next[Point{p.x - 1, p.y + 1}] = true
				next[Point{p.x + 1, p.y + 1}] = true
			}
		}
		for p := range next {
			grid[p.y][p.x] = '*'
		}
		check = next
	}
	return count
}

func part2(lines []string) int {
	grid := grid(lines)
	idx := slices.Index(grid[0], 'S')
	start := Point{idx, 0}
	return simulate(grid, start, make(map[Point]int))
}

func simulate(grid [][]rune, point Point, cache map[Point]int) int {
	if val, ok := cache[point]; ok {
		return val
	}
	if point.y == len(grid)-1 {
		return 1
	}
	switch grid[point.y+1][point.x] {
	case '.':
		next := Point{point.x, point.y + 1}
		result := simulate(grid, next, cache)
		cache[next] = result
		cache[point] = result
		return result
	case '^':
		left := Point{point.x - 1, point.y + 1}
		right := Point{point.x + 1, point.y + 1}
		rleft := simulate(grid, left, cache)
		rright := simulate(grid, right, cache)
		sum := rleft + rright
		cache[left] = rleft
		cache[right] = rright
		cache[point] = sum
		return sum
	}
	panic("unreachable")
}
