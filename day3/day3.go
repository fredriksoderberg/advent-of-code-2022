package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

func main() {
	rucksacks := make([]string, 0)

	readFile, err := os.Open("./day3/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		rucksacks = append(rucksacks, fileScanner.Text())
	}

	readFile.Close()

	var total int
	startSecondComp := 0
	for _, rs := range rucksacks {
		startSecondComp = len(rs) / 2
		common, err := findCommon(rs[:startSecondComp], rs[startSecondComp:])
		if err != nil {
			panic("find common")
		}
		pri, err := getPriority(int32(common[0]))
		if err != nil {
			panic("get priority")
		}
		total += int(pri)
	}
	fmt.Println(cast.ToString(total))

	total = 0
	for i := 0; i < len(rucksacks); i += 3 {
		firstCommon, err := findCommon(rucksacks[i], rucksacks[i+1])
		if err != nil {
			panic("find common")
		}
		secondCommon, err := findCommon(firstCommon, rucksacks[i+2])
		if err != nil {
			panic("find common")
		}
		pri, err := getPriority(int32(secondCommon[0]))
		if err != nil {
			panic("get priority")
		}
		total += int(pri)
	}
	fmt.Println(cast.ToString(total))
}

func findCommon(rs1 string, rs2 string) (string, error) {
	items := make(map[rune]int)
	var common string
	for _, item := range rs1 {
		if val, ok := items[item]; ok {
			items[item] = val + 1
		} else {
			items[item] = 1
		}
	}
	for _, item := range rs2 {
		if _, ok := items[item]; ok {
			common = common + string(item)
		}
	}
	if len(common) > 0 {
		return common, nil
	}
	return "", fmt.Errorf("no common item %v", rs1+":"+rs2)
}

func getPriority(r rune) (int32, error) {
	if r >= 65 && r <= 90 {
		return r - 38, nil
	} else if r >= 97 && r <= 122 {
		return r - 96, nil
	}
	return 0, fmt.Errorf("illegal rune %v", r)
}
