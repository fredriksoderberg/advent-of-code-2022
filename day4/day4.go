package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

func main() {
	sections := make([]string, 0)

	readFile, err := os.Open("./day4/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		sections = append(sections, fileScanner.Text())
	}

	readFile.Close()

	var total int
	for _, sec := range sections {
		sec1 := strings.Split(sec, ",")[0]
		sec2 := strings.Split(sec, ",")[1]
		start1 := cast.StringToInt(strings.Split(sec1, "-")[0])
		end1 := cast.StringToInt(strings.Split(sec1, "-")[1])
		start2 := cast.StringToInt(strings.Split(sec2, "-")[0])
		end2 := cast.StringToInt(strings.Split(sec2, "-")[1])
		if fullyOverlap(start1, end1, start2, end2) {
			total++
		}
	}
	fmt.Println(cast.ToString(total))

	total = 0
	for _, sec := range sections {
		sec1 := strings.Split(sec, ",")[0]
		sec2 := strings.Split(sec, ",")[1]
		start1 := cast.StringToInt(strings.Split(sec1, "-")[0])
		end1 := cast.StringToInt(strings.Split(sec1, "-")[1])
		start2 := cast.StringToInt(strings.Split(sec2, "-")[0])
		end2 := cast.StringToInt(strings.Split(sec2, "-")[1])
		if anyOverlap(start1, end1, start2, end2) {
			total++
		}
	}
	fmt.Println(cast.ToString(total))

}

func fullyOverlap(start1 int, end1 int, start2 int, end2 int) bool {
	if start1 >= start2 && end1 <= end2 {
		return true
	}
	if start2 >= start1 && end2 <= end1 {
		return true
	}
	return false
}

func anyOverlap(start1 int, end1 int, start2 int, end2 int) bool {
	if start1 <= end2 && end1 >= start2 {
		return true
	}
	return false
}
