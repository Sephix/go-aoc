package day6

import "testing"

func TestFirst(t *testing.T) {
	testInput := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	runTest(testInput, 7, t)
}

func TestSecond(t *testing.T) {
	testInput := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	runTest(testInput, 5, t)
}

func TestThird(t *testing.T) {
	testInput := "nppdvjthqldpwncqszvftbrmjlhg"
	runTest(testInput, 6, t)
}

func TestFourth(t *testing.T) {
	testInput := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	runTest(testInput, 10, t)
}

func TestFifth(t *testing.T) {
	testInput := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	runTest(testInput, 11, t)
}

func runTest(testInput string, expected int, t *testing.T) {
	result := parseString(testInput)
	if result != expected {
		t.Fatalf(`Test = %d, want %d`, result, expected)
	}
}
