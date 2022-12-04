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

	s1, _ := strconv.Atoi(interval1[0])
	f1, _ := strconv.Atoi(interval1[1])

	s2, _ := strconv.Atoi(interval2[0])
	f2, _ := strconv.Atoi(interval2[1])

	return s1, f1, s2, f2
}

func fullyCover(input string) bool {
	s1, f1, s2, f2 := parseInput(input)
	if s1 >= s2 && f1 <= f2 {
		return true
	}

	if s2 >= s1 && f2 <= f1 {
		return true
	}

	return false
}

func overlapping(input string) bool {
	s1, f1, s2, f2 := parseInput(input)
	return !(f1 < s2 || f2 < s1)
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
