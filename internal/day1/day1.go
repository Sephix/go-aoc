package day1

import (
	"bufio"
	"os"
	"strconv"
)

var input string = "/assets/day1input.txt"

func Result() int {

	return parseInput()
}

func parseInput() int {
	result := 0
	array := make([]int, 1)

	path, _ := os.Getwd()
	file, err := os.Open(path + input)
	if err != nil {
		println(err.Error())
		return 0
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if currentLine == "" {
			if array[len(array)-1] > result {
				result = array[len(array)-1]
			}
			array = append(array, 0)
		} else {
			lineVale, _ := strconv.Atoi(fileScanner.Text())
			array[len(array)-1] = array[len(array)-1] + lineVale
		}
	}
	if array[len(array)-1] > result {
		result = array[len(array)-1]
	}

	return result
}
