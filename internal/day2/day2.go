package day2

import (
	"bufio"
	"fmt"
	"os"
)

var (
	input   string         = "/assets/day2test.txt"
	player1 map[string]int = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	player2 map[string]int = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	outcome map[string]int = map[string]int{
		"AX": 0,
		"AY": -1,
		"AZ": 1,
		"BX": 1,
		"BY": 0,
		"BZ": -1,
		"CX": -1,
		"CY": 1,
		"CZ": 0,
	}
	strategy map[string]int = map[string]int{
		"Y": 0,
		"X": 1,
		"Z": -1,
	}
	score map[int]int = map[int]int{
		0:  3,
		1:  0,
		-1: 6,
	}
)

func PrintResult() {
	fmt.Println("--- Day 2 ---")
	result1, result2 := parseInput()
	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
}

func parseInput() (int, int) {
	path, _ := os.Getwd()
	file, err := os.Open(path + input)
	if err != nil {
		println(err.Error())
		return 0, 0
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	result1 := 0
	result2 := 0

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		p1 := string(currentLine[0])
		p2 := string(currentLine[2])
		result1 += playRound(p1, p2)
		result2 += playStrategy(p1, p2)
	}
	return result1, result2
}

func playRound(p1 string, p2 string) int {
	result := player2[p2]
	result += score[outcome[p1+p2]]
	return result
}

func playStrategy(p1 string, strat string) int {
	currentStrategy := strategy[strat]
	var p2 string
	for key := range player2 {
		if outcome[p1+key] == currentStrategy {
			p2 = key
		}
	}
	return playRound(p1, p2)
}
