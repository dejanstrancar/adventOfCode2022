package challenges

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) (int, int, int, int) {
	values := strings.Split(input, ",")

	interval1 := strings.Split(values[0], "-")
	interval2 := strings.Split(values[1], "-")

	x1, _ := strconv.Atoi(interval1[0])
	y1, _ := strconv.Atoi(interval1[1])

	x2, _ := strconv.Atoi(interval2[0])
	y2, _ := strconv.Atoi(interval2[1])

	return x1, y1, x2, y2
}

func fullyCover(input string) bool {
	x1, y1, x2, y2 := parseInput(input)
	if x1 >= x2 && y1 <= y2 {
		return true
	}

	if x2 >= x1 && y2 <= y1 {
		return true
	}

	return false
}

func overlapping(input string) bool {
	x1, y1, x2, y2 := parseInput(input)
	return !(y1 < x2 || y2 < x1)
}

func Challenge4Part1(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()

	sc := bufio.NewScanner(input)

	result := 0

	for sc.Scan() {
		input := sc.Text()
		counts := fullyCover(input)
		if counts {
			result++
		}
	}

	return result
}

func Challenge4Part2(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()

	sc := bufio.NewScanner(input)

	result := 0

	for sc.Scan() {
		input := sc.Text()
		counts := overlapping(input)
		if counts {
			result++
		}
	}

	return result
}
