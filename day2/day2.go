package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

func main() {
	plays := make([]string, 0)

	readFile, err := os.Open("./day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		plays = append(plays, fileScanner.Text())
	}

	readFile.Close()

	total := 0
	for _, play := range plays {
		gc, err := playScore(play)
		if err != nil {
			panic("play score")
		}
		total += gc
	}
	fmt.Println(cast.ToString(total))

	total = 0
	for _, play := range plays {
		p, err := myPlay(play)
		if err != nil {
			panic("my play")
		}
		gs, err := playScore(p)
		if err != nil {
			panic("play score")
		}
		total += gs
	}
	fmt.Println(cast.ToString(total))
	readFile.Close()
}

func playScore(plays string) (int, error) {
	switch plays {
	case "A X":
		return 4, nil
	case "A Y":
		return 8, nil
	case "A Z":
		return 3, nil
	case "B X":
		return 1, nil
	case "B Y":
		return 5, nil
	case "B Z":
		return 9, nil
	case "C X":
		return 7, nil
	case "C Y":
		return 2, nil
	case "C Z":
		return 6, nil
	default:
		return 0, fmt.Errorf("unknown plays %v", plays)
	}
}

func myPlay(opplay string) (string, error) {
	switch opplay {
	case "A X":
		return "A Z", nil
	case "A Y":
		return "A X", nil
	case "A Z":
		return "A Y", nil
	case "B X":
		return "B X", nil
	case "B Y":
		return "B Y", nil
	case "B Z":
		return "B Z", nil
	case "C X":
		return "C Y", nil
	case "C Y":
		return "C Z", nil
	case "C Z":
		return "C X", nil
	default:
		return "", fmt.Errorf("unknown opplay %v", opplay)
	}
}
