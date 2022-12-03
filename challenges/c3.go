package challenges

import (
	"bufio"
	"os"
	"unicode"
)

func getCharValue(letter rune) int {
	//97 a -> a 1 // 96
	//65 A -> A 27 // 38
	val := int(letter)
	if unicode.IsUpper(letter) {
		return val - 38
	}
	return val - 96

}

func checkCompartments(compartment1 string, compartment2 string) int {
	for _, letter1 := range compartment1 {
		for _, letter2 := range compartment2 {
			if letter1 == letter2 {
				return getCharValue(letter1)
			}
		}
	}

	return 0
}

func checkGroups(input []string) int {
	for _, letter1 := range input[0] {
		for _, letter2 := range input[1] {
			for _, letter3 := range input[2] {
				if letter1 == letter2 && letter1 == letter3 {
					return getCharValue(letter1)
				}
			}
		}
	}

	return 0
}

func Challenge3Part1(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()

	sc := bufio.NewScanner(input)
	result := 0

	for sc.Scan() {
		input := sc.Text()
		length := len(sc.Text()) / 2

		result += checkCompartments(input[:length], input[length:])
	}

	return result
}

func Challenge3Part2(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()

	sc := bufio.NewScanner(input)
	result := 0
	var groups []string

	for sc.Scan() {
		input := sc.Text()
		groups = append(groups, input)
	}

	position := 3

	for position <= len(groups) {
		result += checkGroups(groups[position-3 : position])
		position += 3
	}

	return result
}
