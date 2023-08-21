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
	startSecond, endSecond := convertStringPair(firstPair)
	startFirst, endFirst := convertStringPair(secondPair)

	if startSecond >= startFirst && endSecond <= endFirst || startFirst >= startSecond && endFirst <= endSecond {
		return true
	}
	return false
}

func isPairOverlap(firstPair string, secondPair string) bool {
	startSecond, endSecond := convertStringPair(firstPair)
	startFirst, endFirst := convertStringPair(secondPair)

	if startSecond <= endFirst && endSecond >= startFirst || startFirst <= endSecond && endFirst >= startSecond {
		return true
	}
	return false

}

func convertStringPair(pair string) (int, int) {
	slice := strings.Split(pair, "-")
	first, _ := strconv.Atoi(slice[0])
	second, _ := strconv.Atoi(slice[1])
	return first, second
}
