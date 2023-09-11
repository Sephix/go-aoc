package day11

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day11input.txt"

type OperationHandler func(int) int

type Monkey struct {
	items        []int
	operation    OperationHandler
	test         OperationHandler
	itemsHandled int
}

func PrintResult() {
	fmt.Println("--- Day 11 ---")

	result1 := getResult1()
	result2 := getResult2()

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
}

func getResult1() int {
	result := 0

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	defer file.Close()

	monkeys := make([]*Monkey, 0)

	mod := 1

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			fileScanner.Scan()
			items := parseItems(fileScanner.Text())
			fileScanner.Scan()
			operation := parseOperation(fileScanner.Text())
			fileScanner.Scan()
			test := fileScanner.Text()
			fileScanner.Scan()
			ifTrue := fileScanner.Text()
			fileScanner.Scan()
			ifFalse := fileScanner.Text()
			testFunc, op := parseTest(test, ifTrue, ifFalse)
			mod *= op
			newMonkey := &Monkey{
				items:        items,
				operation:    operation,
				test:         testFunc,
				itemsHandled: 0,
			}
			monkeys = append(monkeys, newMonkey)
		}
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.playRound(monkeys, 0)
		}
	}
	highestTwoMonkey := make([]int, 2, 2)
	for _, monkey := range monkeys {
		currentMonkeyBiz := monkey.itemsHandled
		if currentMonkeyBiz > highestTwoMonkey[1] {
			highestTwoMonkey[1] = currentMonkeyBiz
			if highestTwoMonkey[1] > highestTwoMonkey[0] {
				temp := highestTwoMonkey[1]
				highestTwoMonkey[1] = highestTwoMonkey[0]
				highestTwoMonkey[0] = temp

			}
		}
	}

	result = highestTwoMonkey[0] * highestTwoMonkey[1]

	return result
}
func getResult2() int {
	result := 0

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	defer file.Close()

	monkeys := make([]*Monkey, 0)

	mod := 1

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			fileScanner.Scan()
			items := parseItems(fileScanner.Text())
			fileScanner.Scan()
			operation := parseOperation(fileScanner.Text())
			fileScanner.Scan()
			test := fileScanner.Text()
			fileScanner.Scan()
			ifTrue := fileScanner.Text()
			fileScanner.Scan()
			ifFalse := fileScanner.Text()
			testFunc, op := parseTest(test, ifTrue, ifFalse)
			mod *= op
			newMonkey := &Monkey{
				items:        items,
				operation:    operation,
				test:         testFunc,
				itemsHandled: 0,
			}
			monkeys = append(monkeys, newMonkey)
		}
	}

	for i := 0; i < 10_000; i++ {
		for _, monkey := range monkeys {
			monkey.playRound(monkeys, mod)
		}
	}
	highestTwoMonkey := make([]int, 2, 2)
	for _, monkey := range monkeys {
		currentMonkeyBiz := monkey.itemsHandled
		if currentMonkeyBiz > highestTwoMonkey[1] {
			highestTwoMonkey[1] = currentMonkeyBiz
			if highestTwoMonkey[1] > highestTwoMonkey[0] {
				temp := highestTwoMonkey[1]
				highestTwoMonkey[1] = highestTwoMonkey[0]
				highestTwoMonkey[0] = temp
			}
		}
	}

	result = highestTwoMonkey[0] * highestTwoMonkey[1]

	return result
}

func parseItems(items string) []int {
	result := make([]int, 0)
	itemsString := strings.Split(items[17:], ",")
	for _, v := range itemsString {
		intValue, _ := strconv.Atoi(strings.Trim(v, " "))
		result = append(result, intValue)
	}
	return result
}

func parseOperation(items string) OperationHandler {
	var result OperationHandler
	operationString := strings.Split(items[19:], " ")

	rightString := operationString[0]
	operandString := operationString[1]
	leftString := operationString[2]

	result = func(old int) int {
		right := parseOperator(rightString, old)
		left := parseOperator(leftString, old)
		if operandString == "*" {
			return left * right
		}
		return left + right
	}

	return result
}

func parseOperator(operator string, old int) int {
	result, err := strconv.Atoi(operator)
	if err != nil {
		return old
	}
	return result
}

func parseTest(testString, ifTrueString, ifFalseString string) (OperationHandler, int) {
	var result OperationHandler

	test, _ := strconv.Atoi(testString[21:])
	ifTrue, _ := strconv.Atoi(ifTrueString[29:])
	ifFalse, _ := strconv.Atoi(ifFalseString[30:])

	result = func(old int) int {
		if old%test == 0 {
			return ifTrue
		}
		return ifFalse
	}

	return result, test
}

func (monkey *Monkey) playRound(monkeys []*Monkey, divise int) {
	for len(monkey.items) > 0 {
		monkey.itemsHandled++
		item := monkey.items[0]
		monkey.items = monkey.items[1:]
		item = monkey.operation(item)
		if divise == 0 {
			item /= 3
		} else {
			item %= divise
		}
		toMonkey := monkey.test(item)
		monkeys[toMonkey].items = append(monkeys[toMonkey].items, item)
	}
}
