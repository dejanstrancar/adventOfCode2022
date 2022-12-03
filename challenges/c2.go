package challenges

import (
	"bufio"
	"os"
	"strings"
)

// A - rock
// B - paper
// C - scissors

// X - rock 1
// Y - paper 2
// Z - scissors 3

func roundPoints(played string) int {
	switch played {
	case "win":
		return 6
	case "draw":
		return 3
	default:
		return 0
	}
}

func decideShape(played string, option string) string {
	// X - loose
	// Y - draw
	// Z - win

	switch option {
	case "Z":
		//win
		switch played {
		case "A":
			return "Y"
		case "B":
			return "Z"
		default:
			return "X"
		}
	case "Y":
		//draw
		switch played {
		case "A":
			return "X"
		case "B":
			return "Y"
		default:
			return "Z"
		}
	default:
		//loose
		switch played {
		case "A":
			return "Z"
		case "B":
			return "X"
		default:
			return "Y"
		}
	}
}

func calculateScore(a string, b string) int {
	score := 0

	switch b {

	// played rock
	case "X":
		score = 1
		switch a {
		case "A":
			return score + roundPoints("draw")
		case "B":
			return score + roundPoints("loose")
		default:
			return score + roundPoints("win")
		}

	// played paper
	case "Y":
		score = 2
		switch a {
		case "A":
			return score + roundPoints("win")
		case "B":
			return score + roundPoints("draw")
		default:
			return score + roundPoints("loose")
		}

	// played scissors
	default:
		score = 3
		switch a {
		case "A":
			return score + roundPoints("loose")
		case "B":
			return score + roundPoints("win")
		default:
			return score + roundPoints("draw")
		}
	}
}

func Challenge2Part1(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)
	score := 0

	for sc.Scan() {
		value := sc.Text()
		res1 := strings.Split(value, " ")

		a := res1[0]
		b := res1[1]

		score += calculateScore(a, b)

	}

	return score
}

func Challenge2Part2(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)
	score := 0

	for sc.Scan() {
		value := sc.Text()
		res1 := strings.Split(value, " ")

		a := res1[0]
		b := decideShape(res1[0], res1[1])

		score += calculateScore(a, b)

	}

	return score
}
