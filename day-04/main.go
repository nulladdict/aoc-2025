package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(bytes), "\n")
	points := make(map[complex128]bool)
	for y, row := range rows {
		for x, char := range row {
			if char == '@' {
				points[complex(float64(x), float64(y))] = true
			}
		}
	}
	fmt.Println("part1", part1(points))
	fmt.Println("part2", part2(points))
}

func part1(points map[complex128]bool) int {
	count := 0
	for point := range points {
		neighbors := neighbors8(point, points)
		if neighbors < 4 {
			count++
		}
	}
	return count
}

func neighbors8(point complex128, points map[complex128]bool) int {
	directions := []complex128{1, -1, 1i, -1i, 1 + 1i, 1 - 1i, -1 + 1i, -1 - 1i}
	count := 0
	for _, direction := range directions {
		if points[point+direction] {
			count++
		}
	}
	return count
}

func part2(points map[complex128]bool) int {
	count := 0
	for {
		removed := false
		for point := range points {
			neighbors := neighbors8(point, points)
			if neighbors < 4 {
				removed = true
				count++
				delete(points, point)
			}
		}
		if !removed {
			return count
		}
	}
}
