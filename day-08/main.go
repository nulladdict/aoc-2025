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
	vectors := vectors(input)
	dist := distances(vectors)
	fmt.Println("part1", part1(vectors, dist))
	fmt.Println("part2", part2(vectors, dist))
}

func part1(vectors []Vector, distances [][]int) int {
	connections := connect(vectors, distances, 1000)
	components := components(connections)
	slices.SortFunc(components, func(a, b []int) int {
		return len(b) - len(a)
	})
	return len(components[0]) * len(components[1]) * len(components[2])
}

func part2(vectors []Vector, distances [][]int) int {
	shortest := shortest(distances)
	seen := make(map[int]bool)
	for _, conn := range shortest {
		a, b := conn.a, conn.b
		seen[a] = true
		seen[b] = true
		if len(seen) == len(vectors) {
			return vectors[a].x * vectors[b].x
		}
	}
	panic("unreachable")
}

func components(connections [][]bool) [][]int {
	n := len(connections)
	visited := make([]bool, n)
	components := make([][]int, 0)

	var dfs func(int, *[]int)
	dfs = func(node int, component *[]int) {
		visited[node] = true
		*component = append(*component, node)
		for neighbor, connected := range connections[node] {
			if connected && !visited[neighbor] {
				dfs(neighbor, component)
			}
		}
	}

	for i := range n {
		if !visited[i] {
			component := make([]int, 0)
			dfs(i, &component)
			components = append(components, component)
		}
	}

	return components
}

func connect(vectors []Vector, distances [][]int, count int) [][]bool {
	shortest := shortest(distances)[:count]
	connections := connections(vectors)
	for _, conn := range shortest {
		a, b := conn.a, conn.b
		connections[a][b] = true
		connections[b][a] = true
	}
	return connections
}

func shortest(distances [][]int) []Connection {
	n := len(distances)
	sorted := make([]Connection, 0)
	for i := range distances {
		for j := i + 1; j < n; j++ {
			sorted = append(sorted, Connection{i, j, distances[i][j]})
		}
	}
	slices.SortFunc(sorted, func(a, b Connection) int {
		switch {
		case a.d > b.d:
			return 1
		case a.d < b.d:
			return -1
		default:
			return 0
		}
	})
	return sorted
}

type Connection struct {
	a, b int
	d    int
}

type Vector struct {
	x, y, z int
}

func vectors(input string) []Vector {
	vectors := make([]Vector, 0)
	for line := range strings.SplitSeq(input, "\n") {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		vectors = append(vectors, Vector{x, y, z})
	}
	return vectors
}

func connections(vectors []Vector) [][]bool {
	n := len(vectors)
	connections := make([][]bool, n)
	for i := range connections {
		connections[i] = make([]bool, n)
	}
	return connections
}

func distances(vectors []Vector) [][]int {
	n := len(vectors)
	distances := make([][]int, n)
	for i := range distances {
		distances[i] = make([]int, n)
	}
	for i := range n {
		for j := range n {
			distances[i][j] = squared(vectors[i], vectors[j])
		}
	}
	return distances
}

func squared(a, b Vector) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}
