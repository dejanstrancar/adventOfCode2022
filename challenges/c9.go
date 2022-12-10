package challenges

import (
	"bufio"
	"os"
	"strconv"
)

type point2 struct {
	x, y int
}

func Challenge9Part1(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	visitedByTail := make(map[point2]bool)
	visited := make(map[point2]bool)
	head := point2{0, 0}
	tail := point2{0, 0}
	visited[tail] = true
	for sc.Scan() {
		direction := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])

		//I calculate moves one by one
		for moves > 0 {
			switch direction {
			case 'U':
				head.y++
			case 'R':
				head.x++
			case 'D':
				head.y--
			case 'L':
				head.x--
			}
			moves--
			tail = adjustTail(tail, head)
			visitedByTail[tail] = true
		}
	}

	return len(visitedByTail)
}

func Challenge9Part2(inputFile string) int {
	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	visitedByTail := make(map[point2]bool)
	knots := make([]point2, 10)

	visitedByTail[knots[9]] = true

	for sc.Scan() {
		direction := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])
		//I calculate moves one by one
		for moves > 0 {
			switch direction {
			case 'U':
				knots[0].y++
			case 'R':
				knots[0].x++
			case 'D':
				knots[0].y--
			case 'L':
				knots[0].x--
			}

			for i := range knots[:len(knots)-1] {
				knots[i+1] = adjustTail2(knots[i+1], knots[i])
			}
			moves--
			visitedByTail[knots[9]] = true
		}
	}

	return len(visitedByTail)
}

func adjustTail(tail point2, head point2) (newTail point2) {
	newTail = tail
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{-2, 1}, point2{-1, 2}, point2{0, 2}, point2{1, 2}, point2{2, 1}:
		newTail.y++
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{1, 2}, point2{2, 1}, point2{2, 0}, point2{2, -1}, point2{1, -2}:
		newTail.x++
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{2, -1}, point2{1, -2}, point2{0, -2}, point2{-1, -2}, point2{-2, -1}:
		newTail.y--
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{-1, -2}, point2{-2, -1}, point2{-2, -0}, point2{-2, 1}, point2{-1, 2}:
		newTail.x--
	}
	return
}

func adjustTail2(tail point2, head point2) (newTail point2) {
	newTail = tail
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{-2, 1}, point2{-1, 2}, point2{0, 2}, point2{1, 2}, point2{2, 1}, point2{2, 2}, point2{-2, 2}:
		newTail.y++
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{1, 2}, point2{2, 1}, point2{2, 0}, point2{2, -1}, point2{1, -2}, point2{2, 2}, point2{2, -2}:
		newTail.x++
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{-2, -2}, point2{2, -1}, point2{1, -2}, point2{0, -2}, point2{-1, -2}, point2{-2, -1}, point2{2, -2}:
		newTail.y--
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{-2, -2}, point2{-1, -2}, point2{-2, -1}, point2{-2, -0}, point2{-2, 1}, point2{-1, 2}, point2{-2, 2}:
		newTail.x--
	}
	return
}
