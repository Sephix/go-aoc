package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var input string = "/assets/day1input.txt"

func PrintResult() {
	best3 := parseInput()
	fmt.Printf("Day 1 part1: %d\n", best3[0])
	fmt.Printf("Day 1 part2: %d\n", sum(best3))
}

func parseInput() []int {
	result := make([]int, 3)
	array := make([]int, 1)

	path, _ := os.Getwd()
	file, err := os.Open(path + input)
	if err != nil {
		println(err.Error())
		return nil
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if currentLine == "" {
			updateResult(result, array)
			array = append(array, 0)
		} else {
			lineVale, _ := strconv.Atoi(fileScanner.Text())
			array[len(array)-1] = array[len(array)-1] + lineVale
		}
	}
	updateResult(result, array)

	fmt.Println(result)
	return result
}

func updateResult(currentVal []int, array []int) {
	for i, val := range currentVal {
		if array[len(array)-1] > val {
			for j := 2; j > i; j-- {
				currentVal[j] = currentVal[j-1]
			}
			currentVal[i] = array[len(array)-1]
			break
		}
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
