package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const filename = "./day13/input.txt"

type tree struct {
	value    int
	elements []*tree
	parent   *tree
}

func main() {
	//Read input file
	input, _ := os.Open(filename)
	sc := bufio.NewScanner(input)

	index := 1
	var indexSum int
	for sc.Scan() {
		package1 := readTree(sc.Text())
		sc.Scan()
		package2 := readTree(sc.Text())

		if areOrdered(package1, package2) == 1 {
			indexSum += index
		}

		index++
		sc.Scan()
	}
	input.Close()
	fmt.Println(indexSum)

	//Read input file
	input, _ = os.Open(filename)
	sc = bufio.NewScanner(input)

	var packages []tree
	for sc.Scan() {
		packages = append(packages, readTree(sc.Text()))
		sc.Scan()
		packages = append(packages, readTree(sc.Text()))
		sc.Scan()
	}
	input.Close()

	packages = append(packages, readTree("[[2]]"))
	packages = append(packages, readTree("[[6]]"))

	sort.Slice(packages, func(i, j int) bool {
		return areOrdered(packages[i], packages[j]) == 1
	})

	decoderKey := 1
	for i, p := range packages {
		if areOrdered(p, readTree("[[2]]")) == 0 || areOrdered(p, readTree("[[6]]")) == 0 {
			decoderKey *= i + 1
		}
	}

	fmt.Println(decoderKey)
}

func readTree(input string) tree {
	root := tree{-1, []*tree{}, nil}
	temp := &root

	var currentNumber string
	for _, r := range input {
		switch r {
		case '[':
			newTree := tree{-1, []*tree{}, temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		case ']':
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.value = number
				currentNumber = ""
			}
			temp = temp.parent
		case ',':
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.value = number
				currentNumber = ""
			}
			temp = temp.parent
			newTree := tree{-1, []*tree{}, temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		default:
			currentNumber += string(r)
		}
	}
	return root
}

func areOrdered(first, second tree) int {
	switch {
	case len(first.elements) == 0 && len(second.elements) == 0:
		if first.value > second.value {
			return -1
		} else if first.value == second.value {
			return 0
		}
		return 1

	case first.value >= 0:
		return areOrdered(tree{-1, []*tree{&first}, nil}, second)

	case second.value >= 0:
		return areOrdered(first, tree{-1, []*tree{&second}, nil})
	default:
		var i int
		for i = 0; i < len(first.elements) && i < len(second.elements); i++ {
			ordered := areOrdered(*first.elements[i], *second.elements[i])
			if ordered != 0 {
				return ordered
			}
		}
		if i < len(first.elements) {
			return -1
		} else if i < len(second.elements) {
			return 1
		}
	}
	return 0
}
