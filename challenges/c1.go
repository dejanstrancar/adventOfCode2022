package challenges

import (
	"bufio"
	"os"
	"strconv"
)

func getMax(output []int) ([]int, int, int) {
	sum := 0
	max := 0
	maxIndex := -10
	for index, value := range output {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	sum += max

	output = append(output[:maxIndex], output[maxIndex+1:]...)
	return output, max, sum
}

func Challenge1Part1(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()

	sc := bufio.NewScanner(input)

	var output []int
	counter := 0

	for sc.Scan() {
		calories, _ := strconv.Atoi(sc.Text())
		if calories == 0 {
			counter++
		} else {
			if counter < len(output) {
				output[counter] = output[counter] + calories
			} else {
				output = append(output, calories)
			}

		}
	}

	_, max, _ := getMax(output)
	return max

}

func Challenge1Part2(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)
	var output []int
	counter := 0
	final := 0

	for sc.Scan() {
		calories, _ := strconv.Atoi(sc.Text())
		if calories == 0 {
			counter++
		} else {
			if len(output) > counter {
				output[counter] = output[counter] + calories
			} else {
				output = append(output, calories)
			}

		}
	}

	output, max, _ := getMax(output)
	final += max
	output, max, _ = getMax(output)
	final += max
	_, max, _ = getMax(output)
	final += max

	return final
}
