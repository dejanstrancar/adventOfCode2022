package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
)

func numberAtCycle(content []string, eoc int) int {
	x := 1
	cycle := 0
	for _, instruction := range content {

		if cycle == eoc {
			break
		}
		switch instruction {
		case "noop":
			cycle++
		default:
			for i := 1; i <= 2; i++ {
				cycle++
				if cycle == eoc {
					break
				}
				if i == 2 {
					y := 0
					fmt.Sscanf(instruction, "addx %d", &y)
					x = x + y
				}
			}
		}
	}

	return x
}

func numberAtCycle2(content []string, start int, eoc int, startingX int) int {
	x := startingX
	cycle := start
	for _, instruction := range content {

		if cycle == eoc {
			break
		}
		switch instruction {
		case "noop":
			cycle++
			fmt.Print(".")
		default:
			for i := 1; i <= 2; i++ {
				cycle++
				if cycle == eoc {
					break
				}
				if i == 2 {
					y := 0
					fmt.Sscanf(instruction, "addx %d", &y)
					x = x + y
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
		}
	}

	return x
}

func Challenge10Part1(inputFile string) int {

	content := utils.LoadFileToArray(inputFile)
	fmt.Println(content)
	cycles := []int{20, 60, 100, 140, 180, 220}
	total := 0

	for _, i := range cycles {
		total += i * numberAtCycle(content, i)
	}

	return total
}

func Challenge10Part2(inputFile string) int {

	content := utils.LoadFileToArray(inputFile)
	cycles := []int{40, 80, 120, 160, 200, 240}
	index2 := 0
	x := 1

	for index, i := range cycles {
		if index-1 < 0 {
			index2 = 0
		} else {
			index2 = index - 1
		}

		x = numberAtCycle2(content, cycles[index2], i, x)
		fmt.Println()
	}

	return 0
}
