package challenges

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type point4 struct {
	x, y int
}

func Challenge14Part1(inputFile string) int {

	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	cave := make(map[point4]rune)

	var maxY, maxX int

	//Generate cave
	for sc.Scan() {
		points := strings.Split(sc.Text(), " -> ")
		for i := range points[:len(points)-1] {
			from := strings.Split(points[i], ",")
			to := strings.Split(points[i+1], ",")
			fromX, _ := strconv.Atoi(from[0])
			fromY, _ := strconv.Atoi(from[1])
			toX, _ := strconv.Atoi(to[0])
			toY, _ := strconv.Atoi(to[1])

			cave[point4{toX, toY}] = '#'
			cave[point4{fromX, fromY}] = '#'
			if toY > maxY {
				maxY = toY
			}
			if toX > maxX {
				maxX = toX
			}

			for fromX != toX || fromY != toY {
				cave[point4{fromX, fromY}] = '#'
				switch {
				case fromX < toX:
					fromX++
				case fromY < toY:
					fromY++
				case fromX > toX:
					fromX--
				case fromY > toY:
					fromY--
				}
			}
			if fromY > maxY {
				maxY = fromY
			}
			if fromX > maxX {
				maxX = fromX
			}
		}
	}

	intoVoid := false
	var sand int = 0
	for !intoVoid {
		newSand := point4{500, 0}
		for {
			cave[newSand] = '"'
			if newSand.y+1 > maxY {
				intoVoid = true
				break
			}
			if cave[point4{newSand.x, newSand.y + 1}] < '#' {
				newSand.y++
			} else if cave[point4{newSand.x - 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x--
			} else if cave[point4{newSand.x + 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x++
			} else {
				cave[newSand] = 'o'
				sand++
				break
			}
		}
	}
	return sand
}

func Challenge14Part2(inputFile string) int {

	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	cave := make(map[point4]rune)

	var maxY, maxX, minX int = 0, 0, 500

	//Generate cave
	for sc.Scan() {
		points := strings.Split(sc.Text(), " -> ")
		for i := range points[:len(points)-1] {
			from := strings.Split(points[i], ",")
			to := strings.Split(points[i+1], ",")
			fromX, _ := strconv.Atoi(from[0])
			fromY, _ := strconv.Atoi(from[1])
			toX, _ := strconv.Atoi(to[0])
			toY, _ := strconv.Atoi(to[1])

			cave[point4{toX, toY}] = '#'
			cave[point4{fromX, fromY}] = '#'
			if toY > maxY {
				maxY = toY
			}
			if toX > maxX {
				maxX = toX
			}
			if toX < minX {
				minX = toX
			}

			for fromX != toX || fromY != toY {
				cave[point4{fromX, fromY}] = '#'
				switch {
				case fromX < toX:
					fromX++
				case fromY < toY:
					fromY++
				case fromX > toX:
					fromX--
				case fromY > toY:
					fromY--
				}
			}
			if fromY > maxY {
				maxY = fromY
			}
			if fromX > maxX {
				maxX = fromX
			}
		}
	}
	for i := minX - 500; i < maxX+500; i++ {
		cave[point4{i, maxY + 2}] = '#'
	}

	var sand int = 0
	for {
		newSand := point4{500, 0}
		if cave[newSand] == 'o' {
			break
		}
		for {
			cave[newSand] = '"'
			if cave[point4{newSand.x, newSand.y + 1}] < '#' {
				newSand.y++
			} else if cave[point4{newSand.x - 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x--
			} else if cave[point4{newSand.x + 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x++
			} else {
				cave[newSand] = 'o'
				sand++
				break
			}
		}
	}
	return sand
}
