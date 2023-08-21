package day5

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day5input.txt"

func PrintResult() {
	fmt.Println("--- Day 5 ---")
	result1, result2 := getResult()
	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
}

func getResult() (string, string) {
	result := ""
	result2 := ""

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	defer file.Close()

	stack1, moves := parseInput(fileScanner)
	stack2 := deepCopy(stack1)

	for _, move := range moves {
		runMove9000(stack1, move)
	}
	for _, move := range moves {
		runMove9001(stack2, move)
	}
	for i := range stack1 {
		result += stack1[i][len(stack1[i])-1]
		result2 += stack2[i][len(stack2[i])-1]
	}

	return result, result2
}

func parseInput(scanner *bufio.Scanner) ([][]string, [][]int) {
	moves := make([][]int, 0)
	var stack [][]string

	stackWidth := 0
	parsingStack := true
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || string(line[1]) == "1" {
			parsingStack = false
			continue
		}

		if stackWidth == 0 {
			stackWidth = getStackWidth(line)
			stack = make([][]string, stackWidth, stackWidth)
			for i := 0; i < stackWidth; i++ {
				stack[i] = make([]string, 0)
			}
		}

		if parsingStack {
			for i := 0; i < stackWidth; i++ {
				itemPosition := 1 + i*4
				newItem := string(line[itemPosition])
				if newItem != " " {
					stack[i] = append([]string{newItem}, stack[i]...)
				}
			}
		} else {
			r := regexp.MustCompile("[0-9]+")
			movesString := r.FindAllString(line, -1)
			currentMove := make([]int, 3, 3)
			currentMove[0], _ = strconv.Atoi(movesString[0])
			currentMove[1], _ = strconv.Atoi(movesString[1])
			currentMove[2], _ = strconv.Atoi(movesString[2])
			moves = append(moves, currentMove)
		}
	}
	return stack, moves
}

func getStackWidth(stack string) int {
	return int(math.Ceil(float64(len(stack)) / 4.0))
}

func runMove9000(stack [][]string, move []int) {
	numberOfItems, fromStack, toStack := move[0], move[1]-1, move[2]-1
	targetStack := stack[fromStack]
	itemsToMove, targetStack := targetStack[len(targetStack)-numberOfItems:], targetStack[:len(targetStack)-numberOfItems]
	stack[fromStack] = targetStack
	reverseSlice(itemsToMove)
	stack[toStack] = append(stack[toStack], itemsToMove...)

}
func runMove9001(stack [][]string, move []int) {
	numberOfItems, fromStack, toStack := move[0], move[1]-1, move[2]-1
	targetStack := stack[fromStack]
	itemsToMove, targetStack := targetStack[len(targetStack)-numberOfItems:], targetStack[:len(targetStack)-numberOfItems]
	stack[fromStack] = targetStack
	stack[toStack] = append(stack[toStack], itemsToMove...)
}

func reverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func deepCopy[K interface{}](matrix [][]K) [][]K {
	result := make([][]K, len(matrix))
	for i, v := range matrix {
		innerCopy := make([]K, len(v))
		copy(innerCopy, v)
		result[i] = innerCopy
	}
	return result
}
