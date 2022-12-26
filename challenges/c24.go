package challenges

import (
	"adventOfCode2022/utils"
)

const (
	Right = '>'
	Left  = '<'
	Up    = '^'
	Down  = 'v'
)

var (
	dx, dX, dy, dY int
)

type Blizzard struct {
	x, y int
	dir  rune
}
type Blizards []Blizzard

type Point struct{ X, Y int }
type Point3 struct{ X, Y, t int }

type BMap [][]int

func parseInput24(inputFile string) Blizards {
	blizs := Blizards{}

	content := utils.LoadFileToArray(inputFile)
	dY = len(content) - 1
	for y, line := range content {
		dX = len(line) - 1
		for x, c := range line {
			switch c {
			case Right, Left, Up, Down:
				blizs = append(blizs, Blizzard{x, y, c})
			}
		}
	}
	return blizs
}

func wrap(n, nmin, nmax int) int {
	mod := (nmax + 1) - nmin
	if (n - nmin) < 0 {
		return nmax + (n+1-nmin)%mod
	}
	return (n-nmin)%mod + nmin
}

func (bs *Blizards) Step(n int) BMap {
	bMap := make(BMap, dY+1)
	for y := range bMap {
		bMap[y] = make([]int, dX+1)
	}

	for _, b := range *bs {
		switch b.dir {
		case Right:
			bMap[b.y][wrap(b.x+n, dx+1, dX-1)]++
		case Left:
			bMap[b.y][wrap(b.x-n, dx+1, dX-1)]++
		case Up:
			bMap[wrap(b.y-n, dy+1, dY-1)][b.x]++
		case Down:
			bMap[wrap(b.y+n, dy+1, dY-1)][b.x]++
		}
	}
	return bMap
}

func (bmap *BMap) isClear(p Point3) bool {
	if (p.X == dx+1 && p.Y == dy) || (p.Y == dY && p.X == dX-1) {
		return true
	}
	if p.X <= dx || p.X >= dX || p.Y <= dy || p.Y >= dY {
		return false
	}
	if (*bmap)[p.Y][p.X] > 0 { // has blizzard
		return false
	}
	return true
}

func (p Point3) nextStepPoints(t int) []Point3 {
	return []Point3{
		{p.X, p.Y, t},
		{p.X + 1, p.Y, t},
		{p.X - 1, p.Y, t},
		{p.X, p.Y + 1, t},
		{p.X, p.Y - 1, t},
	}
}

func Challenge24Part1(inputFile string) int {
	blizs := parseInput24(inputFile)
	// setup State
	startingPoint, finishingPoint := Point3{dx + 1, dy, 0}, Point{dX - 1, dY}

	// main cycle
	queue := map[Point3]bool{startingPoint: true}
	for t := 1; ; t++ {
		bmap := blizs.Step(t)
		nextQueue := map[Point3]bool{}

		for p := range queue {
			if p.X == finishingPoint.X && p.Y == finishingPoint.Y {
				return t - 1
			}

			for _, nextP := range p.nextStepPoints(t) {
				if _, ok := nextQueue[nextP]; ok {
					continue
				}
				if !bmap.isClear(nextP) {
					continue
				}
				nextQueue[nextP] = true
			}
			queue = nextQueue
		}
	}
}

func Challenge24Part2(inputFile string) int {
	blizs := parseInput24(inputFile)

	startingPoint, finishingPoint := Point3{dx + 1, dy, 0}, Point{dX - 1, dY}
	destination := finishingPoint
	phase := 1

	queue := map[Point3]bool{startingPoint: true}
	for t := 1; ; t++ {
		bmap := blizs.Step(t)
		nextQueue := map[Point3]bool{}

	queueLabel:
		for p := range queue {
			if p.X == destination.X && p.Y == destination.Y {
				switch phase {
				case 1:
					destination = Point{startingPoint.X, startingPoint.Y}
					queue = map[Point3]bool{{finishingPoint.X, finishingPoint.Y, t + 1}: true}
					phase = 2
					break queueLabel
				case 2:
					destination = finishingPoint
					queue = map[Point3]bool{{startingPoint.X, startingPoint.Y, t + 1}: true}
					phase = 3
					break queueLabel
				case 3:
					return t - 1
				}
			}

			for _, nextP := range p.nextStepPoints(t) {
				if _, ok := nextQueue[nextP]; ok {
					continue
				}
				if !bmap.isClear(nextP) {
					continue
				}
				nextQueue[nextP] = true
			}
			queue = nextQueue
		}
	}
}
