package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

type queue []rune

func (s queue) Push(v rune) queue {
	return append(s, v)
}

func (s queue) Pop() (queue, rune) {
	l := len(s)
	return s[1:l], s[0]
}

func main() {
	readFile, err := os.Open("./day6/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	buffer := ""
	for fileScanner.Scan() {
		buffer = fileScanner.Text()
	}
	readFile.Close()

	marker := make(queue, 0)
	count := 0
	for i, char := range buffer {
		if len(marker) >= 4 {
			if unique(marker) {
				count = i
				break
			}
			q, _ := marker.Pop()
			marker = q
			marker = marker.Push(char)
		} else {
			marker = marker.Push(char)
		}
	}
	fmt.Println(cast.ToString(count))

	marker = make(queue, 0)
	count = 0
	for i, char := range buffer {
		if len(marker) >= 14 {
			if unique(marker) {
				count = i
				break
			}
			q, _ := marker.Pop()
			marker = q
			marker = marker.Push(char)
		} else {
			marker = marker.Push(char)
		}
	}
	fmt.Println(cast.ToString(count))
}

func unique(q queue) bool {
	chars := make(map[rune]int)
	for _, c := range q {
		chars[c]++
	}
	if len(chars) == len(q) {
		return true
	}
	return false
}
