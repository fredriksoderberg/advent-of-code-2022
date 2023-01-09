package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

const filename = "./day16/input.txt"

type Valve struct {
	name     string
	flowRate int
	conns    []string
}

func main() {
	valves := readInput(filename)
	fmt.Println(valves)
}

func readInput(filename string) []Valve {
	inputFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	sensorPattern, err := regexp.Compile("Valve ([A-Z]+) has flow rate=(\\d+); tunnels lead to valves (\\d+)(,\\s*\\d+)*")
	if err != nil {
		panic(err)
	}

	var valves []Valve
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		matches := sensorPattern.FindAllStringSubmatch(line, len(line))
		firstMatch := matches[0]
		name := firstMatch[1]
		flow := cast.StringToInt(firstMatch[2])
		conns := strings.Split(firstMatch[3], ",")

		valve := Valve{
			name:     name,
			flowRate: flow,
			conns:    conns,
		}
		valves = append(valves, valve)
	}
	return valves
}
