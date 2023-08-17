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
	set1 := toSortedSet(compart1)
	set2 := toSortedSet(compart2)

	for key := range set1 {
		if set2[key] {
			result = key
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

	set1 := toSortedSet(compart1)
	set2 := toSortedSet(compart2)
	set3 := toSortedSet(compart3)

	for key := range set1 {
		if set2[key] && set3[key] {
			result = key
			break
		}
	}

	return result
}

func toSortedSet(items string) map[rune]bool {
	result := make(map[rune]bool)
	for _, val := range items {
		result[val] = true
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
