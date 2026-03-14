package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, _ := os.ReadFile("input.txt")
	str := string(lines)
	str = strings.TrimSpace(str)
	rotations := strings.Split(str, "\n")
	fmt.Println("part1", part1(rotations))
	fmt.Println("part2", part2(rotations))
}

func part1(rotations []string) int {
	dial := 50
	sum := 0
	for _, rotation := range rotations {
		num, _ := strconv.Atoi(rotation[1:])
		if rotation[0] == 'L' {
			dial = (dial - num + 100) % 100
		}
		if rotation[0] == 'R' {
			dial = (dial + num) % 100
		}
		if dial == 0 {
			sum += 1
		}
	}
	return sum
}

func part2(rotations []string) int {
	dial := 50
	sum := 0
	for _, rotation := range rotations {
		num, _ := strconv.Atoi(rotation[1:])
		sum += num / 100
		num = num % 100
		if rotation[0] == 'L' {
			next := (dial - num + 100) % 100
			if next > dial && dial != 0 {
				sum += 1
			}
			dial = next
		}
		if rotation[0] == 'R' {
			next := (dial + num) % 100
			if next < dial && next != 0 {
				sum += 1
			}
			dial = next
		}
		if dial == 0 {
			sum += 1
		}
	}
	return sum
}
