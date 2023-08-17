package day3

import (
	"bufio"
	"fmt"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day3input.txt"

func PrintResult() {
	fmt.Println("--- Day 3 ---")
	result1, result2 := parseInput()
	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 1: ", result2)

}

func parseInput() (int, int) {
	result := 0
	result2 := 0

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	defer file.Close()

	group := make([]string, 3, 3)

	i := 0

	for fileScanner.Scan() {

		currentLine := fileScanner.Text()
		commomElem := findCommonElem(currentLine[:len(currentLine)/2], currentLine[len(currentLine)/2:])
		result += getValue(commomElem)

		group[i] = currentLine
		i += 1

		if i == 3 {
			groupBadge := findCommonGroupElem(group[0], group[1], group[2])
			result2 += getValue(groupBadge)
			group = make([]string, 3)
			i = 0
		}
	}

	return result, result2
}

func findCommonElem(compart1 string, compart2 string) rune {
	var result rune
	for _, i := range []rune(compart1) {
		for _, j := range []rune(compart2) {
			if i == j {
				result = i
				break
			}
		}
		if result != 0 {
			break
		}
	}
	return result
}

func findCommonGroupElem(
	compart1 string,
	compart2 string,
	compart3 string) rune {
	var result rune
	fmt.Println("Finding badge")
	fmt.Println(compart1, compart2, compart3)
	for _, i := range []rune(compart1) {
		for _, j := range []rune(compart2) {
			for _, k := range []rune(compart3) {
				if i == j && i == k {
					result = i
					break
				}
			}
		}
		if result != 0 {
			break
		}
	}
	return result
}

func getValue(c rune) int {
	result := int(c)
	if result < 97 {
		return result - 38
	}
	return int(c) - 96
}
