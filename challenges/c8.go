package challenges

import (
	"bufio"
	"os"
)

type point struct {
	x, y int
}

func seeRight(forest [][]rune, i, j int, house rune, isFirst bool) int {
	if j == len(forest[0])-1 || !isFirst && forest[i][j] >= house {
		return 0
	}
	return 1 + seeRight(forest, i, j+1, house, false)
}

func seeLeft(forest [][]rune, i, j int, house rune, isFirst bool) int {
	if j == 0 || !isFirst && forest[i][j] >= house {
		return 0
	}
	return 1 + seeLeft(forest, i, j-1, house, false)
}

func seeDown(forest [][]rune, i, j int, house rune, isFirst bool) int {
	if i == len(forest)-1 || !isFirst && forest[i][j] >= house {
		return 0
	}
	return 1 + seeDown(forest, i+1, j, house, false)
}

func seeTop(forest [][]rune, i, j int, house rune, isFirst bool) int {
	if i == 0 || !isFirst && forest[i][j] >= house {
		return 0
	}
	return 1 + seeTop(forest, i-1, j, house, false)
}

func Challenge8Part1(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var forest [][]rune
	for sc.Scan() {
		row := []rune{}
		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	maxLeft := make([]rune, len(forest))
	maxRigth := make([]rune, len(forest))

	// Horizontal
	isVisible := make(map[point]bool)
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if j == 0 {
				maxLeft[i] = forest[i][j]
				maxRigth[i] = forest[i][len(forest[0])-1]
				isVisible[point{i, j}] = true
				isVisible[point{i, len(forest[0]) - 1}] = true
				continue
			}
			if forest[i][j] > maxLeft[i] {
				isVisible[point{i, j}] = true
				maxLeft[i] = forest[i][j]
			}
			if forest[i][len(forest[0])-1-j] > maxRigth[i] {
				isVisible[point{i, len(forest[0]) - 1 - j}] = true
				maxRigth[i] = forest[i][len(forest[0])-1-j]
			}
		}
	}

	// Vertical

	maxTop := make([]rune, len(forest))
	maxDown := make([]rune, len(forest))

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if i == 0 {
				maxTop[j] = forest[i][j]
				maxDown[j] = forest[len(forest)-1][j]
				isVisible[point{i, j}] = true
				isVisible[point{len(forest) - 1, j}] = true
				continue
			}
			if forest[i][j] > maxTop[j] {
				isVisible[point{i, j}] = true
				maxTop[j] = forest[i][j]
			}
			if forest[len(forest)-1-i][j] > maxDown[j] {
				isVisible[point{len(forest) - 1 - i, j}] = true
				maxDown[j] = forest[len(forest)-1-i][j]
			}
		}
	}

	return len(isVisible)
}

func Challenge8Part2(inputFile string) int {

	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var forest [][]rune
	for sc.Scan() {
		row := []rune{}
		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	var HighestScore int
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			score := seeDown(forest, i, j, forest[i][j], true) *
				seeLeft(forest, i, j, forest[i][j], true) *
				seeRight(forest, i, j, forest[i][j], true) *
				seeTop(forest, i, j, forest[i][j], true)
			if score > HighestScore {
				HighestScore = score
			}
		}
	}
	return HighestScore
}
