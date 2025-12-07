package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]
	fmt.Println("part1", part1(lines))
	fmt.Println("part2", part2(lines))
}

func part1(lines []string) int {
	rows := make([][]int, 0)
	for _, row := range lines[:len(lines)-1] {
		nums := make([]int, 0)
		for _, w := range strings.Fields(row) {
			num, _ := strconv.Atoi(w)
			nums = append(nums, num)
		}
		rows = append(rows, nums)
	}
	ops := strings.Fields(lines[len(lines)-1])
	sum := 0
	for i := 0; i < len(ops); i++ {
		op, neurtal := operation(ops[i])
		res := neurtal
		for _, row := range rows {
			res = op(res, row[i])
		}
		sum += res
	}
	return sum
}

func operation(op string) (func(int, int) int, int) {
	switch op {
	case "+":
		return func(a, b int) int { return a + b }, 0
	case "*":
		return func(a, b int) int { return a * b }, 1
	default:
		panic("unknown operation: " + op)
	}
}

func part2(lines []string) int {
	nums := lines[:len(lines)-1]
	ops := lines[len(lines)-1]
	sum := 0
	idx := 0
	for {
		if idx >= len(ops) {
			return sum
		}
		char := ops[idx]
		if char == '+' || char == '*' {
			length := 0
			for {
				next := idx + length + 1
				if next >= len(ops) {
					length++
					break
				}
				if ops[next] == '+' || ops[next] == '*' {
					break
				}
				length++
			}
			op, neurtal := operation(string(char))
			res := neurtal
			for d := 0; d < length; d++ {
				num := numAt(nums, idx+d)
				res = op(res, num)
			}
			sum += res
			idx += length + 1
			continue
		}
		panic("unreachable")
	}
}

func numAt(nums []string, idx int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		char := nums[i][idx]
		var digit int
		fmt.Sscanf(string(char), "%d", &digit)
		if num != 0 && digit == 0 {
			break
		}
		num = num*10 + digit
	}
	return num
}
