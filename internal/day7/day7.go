package day7

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day7input.txt"
var totalDiskSpace = 70_000_000
var freeSpaceNeeded = 30_000_000

type CmdType string

const (
	Cd CmdType = "cd"
	Ls CmdType = "ls"
)

type Cmd struct {
	name CmdType
	arg  string
}

type File struct {
	name string
	size int
}

type DirectoryTree struct {
	name    string
	content []File
	child   map[string]*DirectoryTree
	parent  *DirectoryTree
}

func PrintResult() {
	fmt.Println("--- Day 7 ---")

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

	// input := strings.Split(cmdList, "\n")

	fileSystem := &DirectoryTree{
		name:    "root",
		content: make([]File, 0),
		child:   make(map[string]*DirectoryTree, 1),
		parent:  nil,
	}

	for fileScanner.Scan() {
		v := fileScanner.Text()
		if string(v[0]) == "$" {
			cmd := parseCmd(v)
			nextFs := fileSystem.execCmd(cmd)
			fileSystem = nextFs
		} else {
			file, err := parseLine(v)
			if err == nil {
				fileSystem.content = append(fileSystem.content, file)
			}
		}
	}

	result, result2 = sumOf100kDir(fileSystem)

	return result, result2
}

func parseCmd(cmd string) Cmd {
	lineSplit := strings.Split(cmd, " ")
	name := CmdType(lineSplit[1])
	arg := ""
	if len(lineSplit) == 3 {
		arg = lineSplit[2]

	}
	return Cmd{
		name: name,
		arg:  arg,
	}
}

func (dir *DirectoryTree) execCmd(cmd Cmd) *DirectoryTree {
	switch name := cmd.name; name {
	case Cd:
		if cmd.arg == ".." {
			if dir.parent != nil {
				return dir.parent
			}
		} else if cmd.arg == "/" {
			return dir.findRoot()
		} else {
			if v, ok := dir.child[cmd.arg]; ok {
				return v
			} else {
				newNode := DirectoryTree{
					name:    cmd.arg,
					content: make([]File, 0),
					child:   make(map[string]*DirectoryTree, 0),
					parent:  dir,
				}
				dir.child[cmd.arg] = &newNode
				return &newNode
			}
		}
	}
	return dir
}

func (dir *DirectoryTree) findRoot() *DirectoryTree {
	if dir.parent == nil {
		return dir
	}
	return dir.parent.findRoot()
}

func parseLine(line string) (File, error) {
	if strings.HasPrefix(line, "dir") {
		return File{}, errors.New("Unhandled type dir")
	}
	lineSplit := strings.Split(line, " ")
	fileSize, err := strconv.Atoi(lineSplit[0])
	if err != nil {
		return File{}, errors.New("Unhandled file format")
	}
	return File{
		name: lineSplit[1],
		size: fileSize,
	}, nil
}

func sumOf100kDir(fileSystem *DirectoryTree) (int, int) {
	root := fileSystem.execCmd(Cmd{
		name: Cd,
		arg:  "/",
	})

	return bfs(*root)
}

func bfs(fileSystem DirectoryTree) (int, int) {
	sum := 0
	free := freeSpaceNeeded

	spaceUsed := dfs(fileSystem)
	spaceLeft := totalDiskSpace - spaceUsed
	spaceToFree := freeSpaceNeeded - spaceLeft

	fmt.Println(spaceUsed, spaceToFree)

	queue := make([]DirectoryTree, 0)
	queue = append(queue, fileSystem)

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		if elemSum := dfs(element); elemSum <= 100_000 {
			sum += elemSum
		} else if elemSum >= spaceToFree && elemSum <= free {
			free = elemSum
		}
		for _, child := range element.child {
			queue = append(queue, *child)
		}

	}

	return sum, free
}

func dfs(fileSystem DirectoryTree) int {
	currentDirSize := sumDirContent(fileSystem)
	for _, child := range fileSystem.child {
		currentDirSize += dfs(*child)
	}
	return currentDirSize
}

func sumDirContent(fileSystem DirectoryTree) int {
	result := 0
	for _, v := range fileSystem.content {
		result += v.size
	}
	return result
}
