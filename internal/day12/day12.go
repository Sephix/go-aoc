package day12

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/sephix/go-aoc/internal/utils"
)

var input string = "day12input.txt"

// var input string = "day12test.txt"

type Vertex struct {
	x, y  int
	value rune
}

func PrintResult() {
	fmt.Println("--- Day 12 ---")

	result1, result2 := getResult()

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
}

func getResult() (int, int) {
	result := 0

	file, _ := utils.GetFile(input)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	start := Vertex{0, 0, 'a'}
	end := Vertex{0, 0, 'z'}

	graph := make([]Vertex, 0)

	defer file.Close()

	i := 0
	var startIdx, endIdx int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		currentMatrixLine := strings.Split(line, "")
		for k, v := range currentMatrixLine {
			if v == "S" {
				start = Vertex{i, k, 'a'}
				graph = append(graph, start)
				startIdx = len(graph) - 1
			} else if v == "E" {
				end = Vertex{i, k, 'z'}
				graph = append(graph, end)
				endIdx = len(graph) - 1
			} else {
				vertex := Vertex{i, k, []rune(v)[0]}
				graph = append(graph, vertex)
			}
		}
		i++
	}

	adjList := make(map[int][]*Item)
	for i, v := range graph {
		adjList[i] = v.getNeighbors(graph)
	}

	dResult, prev := dijkstra(len(graph), adjList, startIdx)
	result = dResult[endIdx]

	result2 := 0

	currentPos := prev[endIdx]

	for {
		result2 += 1
		if currentPos.value == "a" {
			break
		}
		currentPos = prev[currentPos.idx]
	}

	return result, result2
}

func (vertex Vertex) getNeighbors(graph []Vertex) []*Item {
	result := make([]*Item, 0)

	for i, v := range graph {
		if isNeighbor(vertex, v) {
			newItem := &Item{
				value:    string(v.value),
				priority: 1,
				index:    i,
				idx:      i,
			}
			result = append(result, newItem)
		}
	}

	return result
}

func isNeighbor(v1, v2 Vertex) bool {
	value := v2.value - v1.value
	x := v1.x - v2.x
	y := v1.y - v2.y

	if value < 2 {
		if ((x == 1 || x == -1) && y == 0) ||
			((y == 1 || y == -1) && x == 0) {
			return true
		}
	}

	// result := math.Sqrt(deltaX + deltaY)
	return false
}

func dijkstra(V int, adjList map[int][]*Item, src int) ([]int, map[int]Item) {
	distance := make([]int, V)
	prev := make(map[int]Item)
	for i := 0; i < V; i++ {
		distance[i] = 1_000_000
	}
	distance[src] = 0

	pq := make(PriorityQueue, 0)

	pq.Push(&Item{
		value:    "a",
		index:    src,
		priority: 0,
		idx:      src,
	})

	for pq.Len() > 0 {
		x := pq.Pop()
		current := x.(*Item)

		for _, n := range adjList[current.idx] {
			if distance[current.idx]+n.priority < distance[n.idx] {
				distance[n.idx] = n.priority + distance[current.idx]
				prev[n.idx] = *current
				// fmt.Printf("From %v to %v\n", current.value, n.value)
				pq.Push(n)
			}
		}
	}
	return distance, prev
}
