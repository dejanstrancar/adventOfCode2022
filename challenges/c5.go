package challenges

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	elements []string
}

func (s *stack) append(c string) {
	s.elements = append([]string{c}, s.elements...)
}

func (s *stack) push(c string) {
	s.elements = append(s.elements, c)
}

func (s *stack) pop() (c string) {
	c = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

func (s *stack) pushStack(c []string) {
	s.elements = append(s.elements, c...)
}

func (s *stack) popStack(n int) (c []string) {
	c = s.elements[len(s.elements)-n : len(s.elements)]
	s.elements = s.elements[:len(s.elements)-n]
	return
}

func processInitial(input string, positions []stack) {
	for index, char := range input {
		char := string(char)
		if char != "[" && char != "]" && char != " " {
			positions[index/4].append(char)
		}
	}
}

func readPositions(sc *bufio.Scanner, positions []stack) {
	sc.Scan()
	for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
		processInitial(sc.Text(), positions)
		sc.Scan()
	}
}

func readInstructions(sc *bufio.Scanner) (element int, from int, to int) {
	fmt.Sscanf(sc.Text(), "move %d from %d to %d", &element, &from, &to)
	return
}

func Challenge5Part1(inputFile string) (result string) {
	input, _ := os.Open(inputFile)
	defer input.Close()
	var positions = make([]stack, 9)

	sc := bufio.NewScanner(input)

	readPositions(sc, positions)

	sc.Scan() //empty line

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
	var positions = make([]stack, 9)

	sc := bufio.NewScanner(input)

	readPositions(sc, positions)

	sc.Scan() //empty line

	for sc.Scan() {
		element, from, to := readInstructions(sc)
		crates := positions[from-1].popStack(element)
		positions[to-1].pushStack(crates)
	}

	for _, c := range positions {
		result += c.pop()
	}

	return
}
