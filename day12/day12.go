package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "./day12/input.txt"

// Point in matrix with distance from start.
type Point struct {
	row      int
	col      int
	distance int
}

func main() {
	matrix, startRow, startCol, endRow, endCol := readInput(filename)
	// do breadth first search in matrix
	// keep track of visited nodes
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}
	// keep track of nodes to visit
	queue := make([]Point, 0)
	queue = append(queue, Point{startRow, startCol, 0})
	// keep track of max distance
	maxDistance := 0
	for len(queue) > 0 {
		// pop first element
		current := queue[0]
		queue = queue[1:]
		// destination found
		if current.row == endRow && current.col == endCol {
			maxDistance = current.distance
			break
		}
		// update max distance
		if current.distance > maxDistance {
			maxDistance = current.distance
		}
		// add neighbours to queue
		neighbours := []Point{
			{current.row + 1, current.col, current.distance + 1},
			{current.row - 1, current.col, current.distance + 1},
			{current.row, current.col + 1, current.distance + 1},
			{current.row, current.col - 1, current.distance + 1},
		}
		for _, neighbour := range neighbours {
			// check if out of bounds
			if neighbour.row < 0 || neighbour.row >= len(matrix) || neighbour.col < 0 || neighbour.col >= len(matrix[0]) {
				continue
			}
			// check if higher than current
			if matrix[neighbour.row][neighbour.col]-matrix[current.row][current.col] > 1 {
				continue
			}
			// check if visited
			if visited[neighbour.row][neighbour.col] {
				continue
			}
			//mark as visited
			visited[neighbour.row][neighbour.col] = true
			// add to queue
			queue = append(queue, neighbour)
		}
	}
	// print result
	fmt.Println(maxDistance)

	matrix, _, _, endRow, endCol = readInput(filename)
	// switch start and end coordinates
	startRow = endRow
	startCol = endCol
	// do breadth first search in matrix
	// keep track of visited nodes
	visited = make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}
	// keep track of nodes to visit
	queue = make([]Point, 0)
	queue = append(queue, Point{startRow, startCol, 0})
	// keep track of max distance
	maxDistance = 0
	for len(queue) > 0 {
		// pop first element
		current := queue[0]
		queue = queue[1:]
		// destination 'a' found
		if matrix[current.row][current.col] == 97 {
			maxDistance = current.distance
			break
		}
		// update max distance
		if current.distance > maxDistance {
			maxDistance = current.distance
		}
		// add neighbours to queue
		neighbours := []Point{
			{current.row + 1, current.col, current.distance + 1},
			{current.row - 1, current.col, current.distance + 1},
			{current.row, current.col + 1, current.distance + 1},
			{current.row, current.col - 1, current.distance + 1},
		}
		for _, neighbour := range neighbours {
			// check if out of bounds
			if neighbour.row < 0 || neighbour.row >= len(matrix) || neighbour.col < 0 || neighbour.col >= len(matrix[0]) {
				continue
			}
			// check if lower than current
			if matrix[current.row][current.col]-matrix[neighbour.row][neighbour.col] > 1 {
				continue
			}
			// check if visited
			if visited[neighbour.row][neighbour.col] {
				continue
			}
			//mark as visited
			visited[neighbour.row][neighbour.col] = true
			// add to queue
			queue = append(queue, neighbour)
		}
	}
	// print result
	fmt.Println(maxDistance)
}

// import text file into a matrix of runes
func readInput(filename string) ([][]rune, int, int, int, int) {
	matrix := make([][]rune, 0)
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		matrix = append(matrix, []rune(fileScanner.Text()))
	}
	readFile.Close()
	// find start position
	startRow := 0
	startCol := 0
	endRow := 0
	endCol := 0
	for r, row := range matrix {
		for c, col := range row {
			if col == 'S' {
				startRow = r
				startCol = c
				matrix[startRow][startCol] = 'a'
			}
			if col == 'E' {
				endRow = r
				endCol = c
				matrix[endRow][endCol] = 'z'
			}
		}
	}
	return matrix, startRow, startCol, endRow, endCol
}
