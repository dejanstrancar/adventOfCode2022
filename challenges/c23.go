package challenges

import (
	"adventOfCode2022/utils"
)

type P struct {
	x int
	y int
}

type Elfs map[P]*P

func parseInput25(inputFile string) Elfs {
	elfsMap := make(Elfs)
	content := utils.LoadFileToArray(inputFile)
	for y, line := range content {
		for x, c := range line {
			if c == '#' {
				elfsMap[P{x, y}] = nil
			}
		}
	}
	return elfsMap
}

func (me *Elfs) plan(round int) map[P]int {
	directions := []string{"N", "S", "W", "E"}
	planCounter := make(map[P]int)

	for elf := range *me {

		adjFound := false
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}
				if _, ok := (*me)[P{elf.x + x, elf.y + y}]; ok {
					adjFound = true
					break
				}
			}
		}
		if !adjFound {
			continue
		}

		for dirIdx := 0; dirIdx < 4; dirIdx++ {
			dir := directions[(round+dirIdx)%4]

			checkPos := [3]P{}
			newP := P{}
			switch dir {
			case "N":
				checkPos[0] = P{elf.x - 1, elf.y - 1}
				checkPos[1] = P{elf.x, elf.y - 1}
				checkPos[2] = P{elf.x + 1, elf.y - 1}
				newP = P{elf.x, elf.y - 1}
			case "S":
				checkPos[0] = P{elf.x - 1, elf.y + 1}
				checkPos[1] = P{elf.x, elf.y + 1}
				checkPos[2] = P{elf.x + 1, elf.y + 1}
				newP = P{elf.x, elf.y + 1}
			case "W":
				checkPos[0] = P{elf.x - 1, elf.y - 1}
				checkPos[1] = P{elf.x - 1, elf.y}
				checkPos[2] = P{elf.x - 1, elf.y + 1}
				newP = P{elf.x - 1, elf.y}
			case "E":
				checkPos[0] = P{elf.x + 1, elf.y - 1}
				checkPos[1] = P{elf.x + 1, elf.y}
				checkPos[2] = P{elf.x + 1, elf.y + 1}
				newP = P{elf.x + 1, elf.y}
			}

			_, ok0 := (*me)[checkPos[0]]
			_, ok1 := (*me)[checkPos[1]]
			_, ok2 := (*me)[checkPos[2]]
			if !ok0 && !ok1 && !ok2 {
				(*me)[elf] = &newP
				planCounter[newP]++
				break
			}
		}
	}
	return planCounter
}

func (me *Elfs) boundary() (x, X, y, Y int) {
	for elf := range *me {
		if elf.x < x {
			x = elf.x
		}
		if elf.x > X {
			X = elf.x
		}
		if elf.y < y {
			y = elf.y
		}
		if elf.y > Y {
			Y = elf.y
		}
	}
	return
}

func (me *Elfs) countEmtpyTiles() int {
	x, X, y, Y := (*me).boundary()

	elfCount := len(*me)
	fullRect := (X - x + 1) * (Y - y + 1)

	return fullRect - elfCount
}

func (me *Elfs) exec(planCounter map[P]int) Elfs {
	newElfsMap := make(Elfs)
	for elf, plan := range *me {
		if plan != nil && planCounter[*plan] == 1 {
			newElfsMap[*plan] = nil
		} else {
			newElfsMap[elf] = nil
		}
	}
	return newElfsMap
}

func (me *Elfs) exec2(planCounter map[P]int) (Elfs, int) {
	newElfsMap := make(Elfs)
	movedElfs := 0
	for elf, plan := range *me {
		if plan != nil && planCounter[*plan] == 1 {
			movedElfs++
			newElfsMap[*plan] = nil
		} else {
			newElfsMap[elf] = nil
		}
	}
	return newElfsMap, movedElfs
}

func Challenge23Part1(inputFile string) int {
	elfs := parseInput25(inputFile)

	for i := 0; i < 10; i++ {
		planCounter := elfs.plan(i)
		elfs = elfs.exec(planCounter)

	}

	return elfs.countEmtpyTiles()
}

func Challenge23Part2(inputFile string) int {
	elfs := parseInput25(inputFile)
	var moved, solution int

	for i := 0; ; i++ {
		planCounter := elfs.plan(i)
		elfs, moved = elfs.exec2(planCounter)
		if moved == 0 {
			solution = i + 1
			break
		}
	}
	return solution
}
