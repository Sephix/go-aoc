package day4

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day4input.txt"

func PrintResult() {
	fmt.Println("--- Day 4 ---")
	result1, result2 := parseInput()
	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)

}

func parseInput() (int, int) {
	result := 0
	result2 := 0

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	defer file.Close()

	for fileScanner.Scan() {
		var elvesPair = strings.Split(fileScanner.Text(), ",")
		if isPairContained(elvesPair[0], elvesPair[1]) {
			result += 1
		}
		if isPairOverlap(elvesPair[0], elvesPair[1]) {
			result2 += 1
		}
	}

	return result, result2
}

func isPairContained(firstPair string, secondPair string) bool {
	strLowestPair := strings.Split(firstPair, "-")
	strHighestPair := strings.Split(secondPair, "-")

	lowestPair := make([]int, 2)
	lowestPair[0], _ = strconv.Atoi(strLowestPair[0])
	lowestPair[1], _ = strconv.Atoi(strLowestPair[1])

	highestPair := make([]int, 2)
	highestPair[0], _ = strconv.Atoi(strHighestPair[0])
	highestPair[1], _ = strconv.Atoi(strHighestPair[1])

	if lowestPair[0] >= highestPair[0] && highestPair[1] >= lowestPair[1] {
		temp := lowestPair
		lowestPair = highestPair
		highestPair = temp
	}
	if highestPair[0] >= lowestPair[0] && lowestPair[1] >= highestPair[1] {
		temp := lowestPair
		lowestPair = highestPair
		highestPair = temp
	}

	if lowestPair[0] >= highestPair[0] && lowestPair[1] <= highestPair[1] {
		return true
	}
	return false
}

func isPairOverlap(firstPair string, secondPair string) bool {
	strLowestPair := strings.Split(firstPair, "-")
	strHighestPair := strings.Split(secondPair, "-")

	lowestPair := make([]int, 2)
	lowestPair[0], _ = strconv.Atoi(strLowestPair[0])
	lowestPair[1], _ = strconv.Atoi(strLowestPair[1])

	highestPair := make([]int, 2)
	highestPair[0], _ = strconv.Atoi(strHighestPair[0])
	highestPair[1], _ = strconv.Atoi(strHighestPair[1])

	if lowestPair[0] >= highestPair[0] {
		temp := lowestPair
		lowestPair = highestPair
		highestPair = temp
	}

	if lowestPair[1] <= highestPair[0] && lowestPair[1] < highestPair[0] {
		return false
	}

	fmt.Println(lowestPair, highestPair)
	return true

}
