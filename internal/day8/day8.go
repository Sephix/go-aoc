package day8

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day8test.txt"

type Color string

const (
	Green Color = "green"
	Red   Color = "red"
)

type Tree struct {
	height int
	color  Color
}

func PrintResult() {
	fmt.Println("--- Day 8 ---")

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

	treeMatrix := make([][]*Tree, 0)

	for fileScanner.Scan() {
		treeLine := make([]*Tree, 0)
		v := fileScanner.Text()
		treesToAdd := strings.Split(v, "")
		for _, tree := range treesToAdd {
			height, _ := strconv.Atoi(tree)
			newTree := &Tree{
				height: height,
				color:  Red,
			}
			treeLine = append(treeLine, newTree)
		}
		treeMatrix = append(treeMatrix, treeLine)
	}

	result, result2 = colorizeGrid(treeMatrix)
	printTreeGrid(treeMatrix)

	return result, result2
}

func colorizeGrid(treeMatrix [][]*Tree) (int, int) {
	result := 0
	result2 := 0

	currentPerimeter := 0

	for i, treeLine := range treeMatrix {
		for j := range treeLine {
			result += colorizeTree(treeMatrix, i, j)
			treeValue := computeTreeValue(treeMatrix, i, j)
			if treeValue > result2 {
				result2 = treeValue
			}
		}
		currentPerimeter++
	}

	return result, result2
}

func colorizeTree(treeMatrix [][]*Tree, top int, left int) int {
	tree := treeMatrix[top][left]
	numberOfGreen := 0
	if !findBlockingTree(treeMatrix, top, left) {
		tree.color = Green
		numberOfGreen++
	}
	return numberOfGreen
}

func findBlockingTree(treeMatrix [][]*Tree, top int, left int) bool {
	tree := treeMatrix[top][left]
	maxHeigth := len(treeMatrix)
	maxWidth := len(treeMatrix[0])
	if top == 0 || top == maxHeigth-1 || left == 0 || left == maxWidth-1 {
		return false
	}
	topLine := getTreeLine(treeMatrix, top, left, "top")
	if !checkLine(*tree, topLine) {
		return false
	}
	rightLine := getTreeLine(treeMatrix, top, left, "right")
	if !checkLine(*tree, rightLine) {
		return false
	}
	bottomLine := getTreeLine(treeMatrix, top, left, "bottom")
	if !checkLine(*tree, bottomLine) {
		return false
	}
	leftLine := getTreeLine(treeMatrix, top, left, "left")
	if !checkLine(*tree, leftLine) {
		return false
	}
	return true
}

func checkLine(tree Tree, treeLine []*Tree) bool {
	for _, currentTree := range treeLine {
		if tree.height <= currentTree.height {
			return true
		}
	}
	return false
}

func getTreeLine(treeMatrix [][]*Tree, top, left int, direction string) []*Tree {
	switch direction {
	case "top":
		result := make([]*Tree, 0)
		for i := top - 1; i >= 0; i-- {
			result = append(result, treeMatrix[i][left])
		}
		return result
	case "right":
		return treeMatrix[top][left+1:]
	case "bottom":
		result := make([]*Tree, 0)
		for i := top + 1; i < len(treeMatrix); i++ {
			result = append(result, treeMatrix[i][left])
		}
		return result
	case "left":
		result := make([]*Tree, 0)
		for i := left - 1; i >= 0; i-- {
			result = append(result, treeMatrix[top][i])
		}
		return result
	}
	return make([]*Tree, 0)
}

func printTreeGrid(treeMatrix [][]*Tree) {
	for _, treeLine := range treeMatrix {
		for j, tree := range treeLine {
			if tree.color == Red {
				fmt.Printf("\033[0;31m %v", tree.height)
			} else {
				fmt.Printf("\033[0;32m %v", tree.height)
			}
			if len(treeLine)-1 == j {
				fmt.Printf("\n")
			}
		}
	}
	fmt.Printf("\033[0;3255m")
}

func computeTreeValue(treeMatrix [][]*Tree, top int, left int) int {
	tree := *treeMatrix[top][left]

	topLine := getTreeLine(treeMatrix, top, left, "top")
	rightLine := getTreeLine(treeMatrix, top, left, "right")
	bottomLine := getTreeLine(treeMatrix, top, left, "bottom")
	leftLine := getTreeLine(treeMatrix, top, left, "left")

	fmt.Println(leftLine)

	topLineValue := treeLineValue(tree, topLine)
	rightLineValue := treeLineValue(tree, rightLine)
	bottomLineValue := treeLineValue(tree, bottomLine)
	leftLineValue := treeLineValue(tree, leftLine)

	return topLineValue * rightLineValue * bottomLineValue * leftLineValue
}

func treeLineValue(tree Tree, treeLine []*Tree) int {
	result := 0
	for _, v := range treeLine {
		result++
		if tree.height <= v.height {
			break
		}
	}
	return result
}
