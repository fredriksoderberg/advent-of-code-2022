package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

const filename = "./day7/input.txt"

var folderSizes = make(map[string]int)
var commands = readInput(filename)

func main() {
	folderSize("/")

	sum := 0
	for _, v := range folderSizes {
		if v <= 100000 {
			sum += v
		}
	}
	fmt.Println(cast.ToString(sum))

	space := 70000000 - folderSizes["/"]
	need := 30000000 - space
	deletable := 30000000
	for _, v := range folderSizes {
		if v >= need {
			if v < deletable {
				deletable = v
			}
		}
	}
	fmt.Println(cast.ToString(deletable))
}

var dirs = func() map[string][]string {
	dirs := make(map[string][]string)
	parent := ""
	for _, cmd := range commands {
		if strings.HasPrefix(cmd, "$ cd ") {
			d := cmd[5:]
			if d == ".." {
				if parent == "/" {
					continue
				}
				li := strings.LastIndex(parent, "/")
				if li == 0 {
					parent = "/"
				} else {
					parent = parent[:li]
				}
			} else {
				if d == "/" {
					parent = "/"
				} else {
					if parent == "/" {
						parent = parent + d
					} else {
						parent = parent + "/" + d
					}
				}
			}
		} else if cmd == "$ ls" {
			continue
		} else {
			dirs[parent] = append(dirs[parent], cmd)
		}
	}
	return dirs
}()

func folderSize(parent string) int {
	if size, ok := folderSizes[parent]; ok {
		return size
	}

	sum := 0
	for _, v := range dirs[parent] {
		if strings.HasPrefix(v, "dir ") {
			subfolder := strings.Split(v, " ")
			if parent == "/" {
				subfolder[1] = "/" + subfolder[1]
			} else {
				subfolder[1] = parent + "/" + subfolder[1]
			}
			sum += folderSize(subfolder[1])
		} else {
			file := strings.Split(v, " ")
			num, _ := strconv.Atoi(file[0])
			sum += num
		}
	}
	folderSizes[parent] = sum
	return sum
}

func readInput(filename string) []string {
	input := make([]string, 0)

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	readFile.Close()
	return input
}
