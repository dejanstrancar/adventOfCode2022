package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
	"math"
)

func SNAUFToDecimal(SNAUF string) int {
	number := 0
	for i := 0; i < len(SNAUF); i++ {
		var digit int
		switch string(SNAUF[len(SNAUF)-1-i]) {
		case "2":
			digit += 2
		case "1":
			digit += 1
		case "0":
			digit += 0
		case "-":
			digit -= 1
		case "=":
			digit -= 2
		}
		number += int(math.Pow(5, float64(i))) * digit
	}
	return number
}

func decimalToSNAUF(num int) string {
	solution := ""

	for num > 0 {
		switch num % 5 {
		case 4:
			solution = fmt.Sprintf("-%s", solution)
			num += 2
		case 3:
			solution = fmt.Sprintf("=%s", solution)
			num += 3
		case 2:
			solution = fmt.Sprintf("2%s", solution)
		case 1:
			solution = fmt.Sprintf("1%s", solution)
		case 0:
			solution = fmt.Sprintf("0%s", solution)
		}
		num /= 5
	}

	return solution
}

func Challenge25Part1(inputFile string) string {
	content := utils.LoadFileToArray(inputFile)

	sum := 0
	for _, SNAUF := range content {
		x := SNAUFToDecimal(SNAUF)
		sum += x
	}

	sol := decimalToSNAUF(sum)
	return sol
}
