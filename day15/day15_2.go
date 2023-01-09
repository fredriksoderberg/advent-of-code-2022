package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Sensor struct {
	X                int
	Y                int
	DistanceToBeacon int
	BeaconX          int
	BeaconY          int
}

func AbsDistance(v1 int, v2 int) int {
	if v1 > v2 {
		return v1 - v2
	}
	return v2 - v1
}

type GridSize struct {
	minX int
	maxX int
	minY int
	maxY int
}

func CountBlockedPositions(sensors []Sensor, row int, grid GridSize) (int, []int) {

	var free []int
	counter := 0
	for x := grid.minX; x <= grid.maxX; x++ {

		isBlocked := false
		isBeacon := false
		isSensor := false
		for _, sensor := range sensors {
			distance := AbsDistance(sensor.X, x) + AbsDistance(sensor.Y, row)
			if sensor.BeaconX == x && sensor.BeaconY == row {
				isBeacon = true
			}
			if sensor.X == x && sensor.Y == row {
				isSensor = true
			}
			if distance <= sensor.DistanceToBeacon {
				isBlocked = true
				x = sensor.X + (sensor.DistanceToBeacon - AbsDistance(row, sensor.Y))
			}
		}
		if isSensor {
			//fmt.Print("S")
			counter++
		} else if isBeacon {
			//fmt.Print("B")
			counter++
		} else if isBlocked {
			//fmt.Print("#")
			counter++
		} else {
			//fmt.Print(".")
			free = append(free, x)
		}
	}
	//fmt.Println("")
	return counter, free
}

func main2() {
	inputFile, err := os.Open("./day15/input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	sensorPattern, err := regexp.Compile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")
	if err != nil {
		panic(err)
	}

	var sensors []Sensor
	scanner := bufio.NewScanner(inputFile)
	maxX := 0
	maxY := 0
	minX := 0
	minY := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := sensorPattern.FindAllStringSubmatch(line, len(line))
		firstMatch := matches[0]
		sensorX, _ := strconv.Atoi(firstMatch[1])
		sensorY, _ := strconv.Atoi(firstMatch[2])
		beaconX, _ := strconv.Atoi(firstMatch[3])
		beaconY, _ := strconv.Atoi(firstMatch[4])

		if maxX < sensorX {
			maxX = sensorX
		}
		if maxY < sensorY {
			maxY = sensorY
		}
		if minX > sensorX {
			minX = sensorX
		}
		if minY > sensorY {
			minY = sensorY
		}

		if maxX < beaconX {
			maxX = beaconX
		}
		if maxY < beaconY {
			maxY = beaconY
		}
		if minX > beaconX {
			minX = beaconX
		}
		if minY > beaconY {
			minY = beaconY
		}

		distanceToBeacon := AbsDistance(sensorX, beaconX) + AbsDistance(sensorY, beaconY)

		//fmt.Printf("Found sensor at %d,%d with beacon %d,%d in distance %d\n", sensorX, sensorY, beaconX, beaconY, distanceToBeacon)
		sensor := Sensor{
			X:                sensorX,
			Y:                sensorY,
			DistanceToBeacon: distanceToBeacon,
			BeaconX:          beaconX,
			BeaconY:          beaconY,
		}
		sensors = append(sensors, sensor)
	}

	// Part 1
	//grid := GridSize{minX: minX, maxX: maxX, minY: minY, maxY: maxY}
	//row := 2000000
	//blocked := CountBlockedPositions(sensors, row, grid)
	//fmt.Printf("Found %d blocked positions in row %d\n", blocked, row)

	grid := GridSize{minX: 0, maxX: 4000000, minY: 0, maxY: 4000000}
	for row := grid.minY; row <= grid.maxY; row++ {
		_, free := CountBlockedPositions(sensors, row, grid)
		if len(free) == 1 {
			tuningFrequency := free[0]*4000000 + row
			fmt.Printf("Found free position %d,%d (%d)\n", free[0], row, tuningFrequency)
		} else if len(free) > 1 {
			fmt.Printf("Found multiple free positions %s\n", free)
		}
		if row%100000 == 0 {
			fmt.Printf("Checked row %d\n", row)
		}
	}
}
