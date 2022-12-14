package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

const filename = "./day14/input.txt"

func main() {
	//Read input file
	matrix := readInput(filename)
	// set origin to +
	matrix[0][500] = '+'
	printMatrix(matrix)
	var curX, curY, count = 500, 0, 0
	bottom := findBottom(matrix)
	var end bool
	for !end {
		curY++
		if curY > bottom {
			end = true
			continue
		}
		if matrix[curY][curX] == '#' || matrix[curY][curX] == 'o' {
			if matrix[curY][curX-1] != '#' && matrix[curY][curX-1] != 'o' {
				curX--
				continue
			}
			if matrix[curY][curX+1] != '#' && matrix[curY][curX+1] != 'o' {
				curX++
				continue
			}
			matrix[curY-1][curX] = 'o'
			printMatrix(matrix)
			curX, curY = 500, 0
			count++
			continue
		}
	}
	fmt.Println(count)

	//Read input file
	matrix = readInput(filename)
	// set origin to +
	matrix[0][500] = '+'
	printMatrix(matrix)
	curX, curY, count = 500, 0, 0
	bottom = findBottom(matrix)
	for x := 0; x < 1000; x++ {
		matrix[bottom+2][x] = '#'
	}
	end = false
	for !end {
		curY++
		if matrix[curY][curX] == '#' || matrix[curY][curX] == 'o' {
			if matrix[curY][curX-1] != '#' && matrix[curY][curX-1] != 'o' {
				curX--
				continue
			}
			if matrix[curY][curX+1] != '#' && matrix[curY][curX+1] != 'o' {
				curX++
				continue
			}
			matrix[curY-1][curX] = 'o'
			printMatrix(matrix)
			count++
			if curY-1 == 0 && curX == 500 {
				end = true
				continue
			}
			curX, curY = 500, 0

			continue
		}
	}
	fmt.Println(count)
}

func readInput(filename string) [][]byte {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	sc := bufio.NewScanner(readFile)
	defer readFile.Close()
	matrix := make([][]byte, 500)
	for i := range matrix {
		matrix[i] = make([]byte, 1000)
	}
	// set all to .
	for y := 0; y < 500; y++ {
		for x := 0; x < 1000; x++ {
			matrix[y][x] = '.'
		}
	}
	for sc.Scan() {
		coords := strings.Split(sc.Text(), " -> ")
		for i := 0; i < len(coords)-1; i++ {
			setStones(coords[i], coords[i+1], matrix)
		}
	}
	return matrix
}

func setStones(coord1 string, coord2 string, matrix [][]byte) {
	var x1, y1, x2, y2 int
	x1 = cast.StringToInt(strings.Split(coord1, ",")[0])
	y1 = cast.StringToInt(strings.Split(coord1, ",")[1])
	x2 = cast.StringToInt(strings.Split(coord2, ",")[0])
	y2 = cast.StringToInt(strings.Split(coord2, ",")[1])

	if x1 == x2 && y1 > y2 {
		for y := y2; y <= y1; y++ {
			matrix[y][x1] = '#'
		}
	} else if x1 == x2 && y1 < y2 {
		for y := y1; y <= y2; y++ {
			matrix[y][x1] = '#'
		}
	} else if y1 == y2 && x1 > x2 {
		for x := x2; x <= x1; x++ {
			matrix[y1][x] = '#'
		}
	} else if y1 == y2 && x1 < x2 {
		for x := x1; x <= x2; x++ {
			matrix[y1][x] = '#'
		}
	}
}

func findBottom(matrix [][]byte) int {
	for y := 499; y > 0; y-- {
		for x := 0; x < 1000; x++ {
			if matrix[y][x] == '#' {
				return y
			}
		}
	}
	return 0
}

func printMatrix(matrix [][]byte) {
	// print matrix
	for y := 0; y < 12; y++ {
		ff := string(matrix[y][493:504])
		fmt.Println(ff)
	}
}
