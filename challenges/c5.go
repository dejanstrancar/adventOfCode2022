package challenges

import (
	"bufio"
	"fmt"
	"os"
)

type Stack []string

func (s *Stack) append(c string) {
	*s = append([]string{c}, *s...)
}

func (s *Stack) push(c string) {
	*s = append(*s, c)
}

func (s *Stack) pop() (c string) {
	c = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack) pushBulk(c []string) {
	*s = append(*s, c...)
}

func (s *Stack) popBulk(n int) (c []string) {
	c = (*s)[len(*s)-n : len(*s)]
	*s = (*s)[:len(*s)-n]
	return
}

func processInitial(input string, positions []Stack) {
	for index, char := range input {
		char := string(char)
		if char != "[" && char != "]" && char != " " {
			positions[index/4].append(char)
		}
	}
}

func readPositions(sc *bufio.Scanner, positions []Stack) {
	sc.Scan()
	for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
		processInitial(sc.Text(), positions)
		sc.Scan()
	}
	sc.Scan() //empty line
}

func readInstructions(sc *bufio.Scanner) (element int, from int, to int) {
	fmt.Sscanf(sc.Text(), "move %d from %d to %d", &element, &from, &to)
	return
}

func Challenge5Part1(inputFile string) (result string) {
	input, _ := os.Open(inputFile)
	defer input.Close()
	var positions = make([]Stack, 9)

	sc := bufio.NewScanner(input)

	readPositions(sc, positions)

	for sc.Scan() {
		element, from, to := readInstructions(sc)
		for i := 0; i < element; i++ {
			crate := positions[from-1].pop()
			positions[to-1].push(crate)
		}
	}

	for _, c := range positions {
		result += c.pop()
	}

	return
}

func Challenge5Part2(inputFile string) (result string) {
	input, _ := os.Open(inputFile)
	defer input.Close()
	var positions = make([]Stack, 9)

	sc := bufio.NewScanner(input)

	readPositions(sc, positions)

	for sc.Scan() {
		element, from, to := readInstructions(sc)
		crates := positions[from-1].popBulk(element)
		positions[to-1].pushBulk(crates)
	}

	for _, c := range positions {
		result += c.pop()
	}

	return
}
