package challenges

import (
	"bufio"
	"os"
)

func checkDuplicates(queue []rune) bool {
	for index, c := range queue {
		for _, c2 := range queue[:index] {
			if c == c2 {
				return true
			}
		}
		for _, c2 := range queue[index+1:] {
			if c == c2 {
				return true
			}
		}
	}

	return false
}

func Challenge6Part1(inputFile string) (result int) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	sc := bufio.NewScanner(input)

	sc.Scan()
	transmition := sc.Text()
	queue := []rune{}

	for index, char := range transmition {

		queue = append(queue, char)
		if len(queue) == 4 {
			if !checkDuplicates(queue) {
				result = index + 1
				return
			}

			queue = queue[1:]

		}

	}

	return
}

func Challenge6Part2(inputFile string) (result int) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	sc := bufio.NewScanner(input)

	sc.Scan()
	transmition := sc.Text()
	queue := []rune{}

	for index, char := range transmition {

		queue = append(queue, char)
		if len(queue) == 14 {
			if !checkDuplicates(queue) {
				result = index + 1
				return
			}

			queue = queue[1:]

		}

	}

	return
}
