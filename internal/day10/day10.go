package day10

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day10input.txt"

type CPU struct {
	cycle    int
	register int
}

type CRT struct {
	cycle          int
	screen         [][]string
	spritePosition int
	spriteLength   int
}

type Command string

const (
	noop Command = "noop"
	addx Command = "addx"
)

type Instruction struct {
	name  Command
	value int
}

func PrintResult() {
	fmt.Println("--- Day 10 ---")

	result1, result2 := getResult()

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ")
	result2.print()
}

func getResult() (int, *CRT) {
	result := 0

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	file2, _ := utils.GetFile(input)
	fileScanner2 := bufio.NewScanner(file2)
	fileScanner2.Split(bufio.ScanLines)

	defer file.Close()
	defer file2.Close()

	cpu := &CPU{
		register: 1,
		cycle:    0,
	}
	crt := &CRT{
		cycle:          0,
		screen:         make([][]string, 6),
		spritePosition: 1,
		spriteLength:   3,
	}

	for i := 0; i < 6; i++ {
		crt.screen[i] = make([]string, 40)
	}
	for fileScanner.Scan() {
		instruction := parseLine(fileScanner.Text())
		result += cpu.execInstruction(instruction, 40, 20)
	}
	for fileScanner2.Scan() {
		instruction := parseLine(fileScanner2.Text())
		crt.execInstruction(instruction)
	}

	return result, crt
}

func parseLine(line string) Instruction {
	commandName := line[:4]
	if commandName == string(noop) {
		return Instruction{name: noop}
	}
	value, _ := strconv.Atoi(line[5:])
	return Instruction{
		name:  addx,
		value: value,
	}
}

func (cpu *CPU) execInstruction(instruction Instruction, monitorCyle, offset int) int {
	result := 0
	switch instruction.name {
	case noop:
		cpu.cycle++
		result += cpu.monitorValue(monitorCyle, offset)
	case addx:
		for i := 0; i < 2; i++ {
			cpu.cycle++
			result += cpu.monitorValue(monitorCyle, offset)
			if i == 1 {
				cpu.register += instruction.value
			}
		}
	}
	return result
}
func (cpu *CPU) monitorValue(monitorCyle, offset int) int {
	result := 0
	if cpu.cycle%(monitorCyle) == offset {
		result = cpu.register * cpu.cycle
	}
	return result
}

func (crt *CRT) execInstruction(instruction Instruction) {
	switch instruction.name {
	case noop:
		crt.flickPixel()
		crt.cycle++
	case addx:
		for i := 0; i < 2; i++ {
			crt.flickPixel()
			crt.cycle++
			if i == 1 {
				crt.spritePosition += instruction.value
			}
		}
	}
}
func (crt *CRT) flickPixel() {
	row := crt.cycle / 40
	column := crt.cycle - 40*row
	if column >= (crt.spritePosition-1) && column < (crt.spritePosition-1)+crt.spriteLength {
		crt.screen[row][column] = "#"
	} else {
		crt.screen[row][column] = "."
	}

}

func (crt *CRT) print() {
	for _, line := range crt.screen {
		for _, pixel := range line {
			fmt.Printf(pixel)
		}
		fmt.Printf("\n")
	}
}
