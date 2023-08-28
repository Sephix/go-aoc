package day9

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day9input.txt"

type Point struct {
	name string
	x, y int
}

func PrintResult() {
	fmt.Println("--- Day 9 ---")

	result1, result2 := getResult()

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
}

func getResult() (int, int) {
	result := 0
	result2 := 0

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	defer file.Close()

	head := &Point{
		name: "head",
		x:    0,
		y:    0,
	}
	tail := &Point{
		name: "tail",
		x:    0,
		y:    0,
	}
	first := make(map[string]bool)

	second := make(map[string]bool)
	numberOfKnot := 10
	knotList := make([]*Point, numberOfKnot, numberOfKnot)
	knotList[0] = head
	for i := 1; i < numberOfKnot; i++ {
		knotList[i] = &Point{
			name: "knot",
			x:    0,
			y:    0,
		}
	}

	for fileScanner.Scan() {
		move, length := parseLine(fileScanner.Text())
		for i := 0; i < length; i++ {
			head.move(move)
			tail.updateTail(*head)
			first[tail.toString()] = true
			for k, v := range knotList[1:] {
				v.updateTail(*knotList[k])
			}
			second[knotList[9].toString()] = true
		}
	}

	result = len(first)
	result2 = len(second)
	return result, result2
}

func (point *Point) move(move rune) {
	switch move {
	case 'U':
		point.y++
	case 'R':
		point.x++
	case 'D':
		point.y--
	case 'L':
		point.x--
	}
}

func parseLine(line string) (rune, int) {
	length, _ := strconv.Atoi(line[2:])
	return rune(line[0]), length
}

func (point *Point) updateTail(head Point) {
	tail := *point
	switch (Point{x: head.x - tail.x, y: head.y - tail.y}) {
	case Point{x: -2, y: 1},
		Point{x: -1, y: 2},
		Point{x: 0, y: 2},
		Point{x: 1, y: 2},
		Point{x: 2, y: 1},
		Point{x: -2, y: 2},
		Point{x: 2, y: 2}:
		point.y++
	}
	switch (Point{x: head.x - tail.x, y: head.y - tail.y}) {
	case Point{x: 1, y: 2},
		Point{x: 2, y: 1},
		Point{x: 2, y: 0},
		Point{x: 2, y: -1},
		Point{x: 1, y: -2},
		Point{x: 2, y: -2},
		Point{x: 2, y: 2}:
		point.x++
	}
	switch (Point{x: head.x - tail.x, y: head.y - tail.y}) {
	case Point{x: 2, y: -1},
		Point{x: 1, y: -2},
		Point{x: 0, y: -2},
		Point{x: -1, y: -2},
		Point{x: -2, y: -1},
		Point{x: -2, y: -2},
		Point{x: 2, y: -2}:
		point.y--
	}
	switch (Point{x: head.x - tail.x, y: head.y - tail.y}) {
	case Point{x: -1, y: -2},
		Point{x: -2, y: -1},
		Point{x: -2, y: -0},
		Point{x: -2, y: 1},
		Point{x: -1, y: 2},
		Point{x: -2, y: -2},
		Point{x: -2, y: 2}:
		point.x--
	}
}

func (point *Point) toString() string {
	result := point.name
	result += fmt.Sprintf("%v", point.x)
	result += fmt.Sprintf("-")
	result += fmt.Sprintf("%v", point.y)
	return result
}
