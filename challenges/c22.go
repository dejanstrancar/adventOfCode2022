package challenges

import (
	"adventOfCode2022/utils"
	"strconv"
	"strings"
	"unicode"
)

var (
	maxCol  int
	maxLine int
)

type P22 struct {
	line int
	col  int
	dir  int
}

type Tile struct {
	isOpen bool
	isWall bool
}

type MonkeyMap map[P22]Tile

func parseInput22(inputFile string) (MonkeyMap, []string) {
	monkeyMap := MonkeyMap{}
	content := utils.ReadRaw(inputFile)
	parts := strings.Split(content, "\n\n")
	mapStrLines := strings.Split(parts[0], "\n")
	maxLine = len(mapStrLines)
	for line, lineStr := range mapStrLines {
		for col, c := range lineStr {
			if len(lineStr) > maxCol {
				maxCol = col
			}

			switch c {
			case '#':
				monkeyMap[P22{line: line, col: col}] = Tile{isWall: true}
			case '.':
				monkeyMap[P22{line: line, col: col}] = Tile{isOpen: true}
			}
		}
	}

	instructions := []string{}

	instructionStr := strings.TrimSpace(parts[1])
	for i := 0; i < len(instructionStr); i++ {
		c := rune(parts[1][i])
		if unicode.IsLetter(c) {
			instructions = append(instructions, string(c))
		} else {
			j := i
			for _, c2 := range instructionStr[i:] {
				if unicode.IsLetter(c2) {
					break
				}
				j++
			}
			instructions = append(instructions, parts[1][i:j])
			i = j - 1
		}
	}

	return monkeyMap, instructions
}

func findInitialPosition(monkeyMap MonkeyMap) P22 {
	for i := 0; ; i++ {
		if _, ok := monkeyMap[P22{line: 0, col: i}]; ok {
			return P22{line: 0, col: i, dir: 0}
		}
	}
}

func (mm *MonkeyMap) step(p *P22) bool {
	switch p.dir {
	// face right
	case 0:
		newCol := p.col + 1
		tile, ok := (*mm)[P22{line: p.line, col: newCol}]
		if !ok { //wrap around
			for i := 0; ; i++ {
				if tile, ok = (*mm)[P22{line: p.line, col: i}]; ok {
					newCol = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.col = newCol
		return true

	// face down
	case 1:
		newLine := p.line + 1
		tile, ok := (*mm)[P22{line: newLine, col: p.col}]
		if !ok { //wrap around
			for i := 0; ; i++ {
				if tile, ok = (*mm)[P22{line: i, col: p.col}]; ok {
					newLine = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.line = newLine
		return true

	// face left
	case 2:
		newCol := p.col - 1
		tile, ok := (*mm)[P22{line: p.line, col: newCol}]
		if !ok { //wrap around
			for i := maxCol; ; i-- {
				if tile, ok = (*mm)[P22{line: p.line, col: i}]; ok {
					newCol = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.col = newCol
		return true

	// face up
	case 3:
		newLine := p.line - 1
		tile, ok := (*mm)[P22{line: newLine, col: p.col}]
		if !ok { //wrap around
			for i := maxLine; ; i-- {
				if tile, ok = (*mm)[P22{line: i, col: p.col}]; ok {
					newLine = i
					break
				}
			}
		}
		if tile.isWall {
			return false
		}
		p.line = newLine
		return true
	}
	return false
}

func (mm *MonkeyMap) move(pos *P22, instr string) {
	if steps, err := strconv.Atoi(instr); err == nil {
		for i := 0; i < steps; i++ {
			if !mm.step(pos) {
				return
			}
		}
		return
	}

	switch instr {
	case "R":
		pos.dir = (pos.dir + 1) % 4
	case "L":
		if pos.dir == 0 {
			pos.dir = 3
		} else {
			pos.dir--
		}
	}

}

func (mm *MonkeyMap) stepOnCube(p *P22, cubeWrapper CubeWrapper) bool {
	if newP, ok := cubeWrapper[*p]; ok {
		if (*mm)[P22{line: newP.line, col: newP.col}].isWall {
			return false
		}
		p.line = newP.line
		p.col = newP.col
		p.dir = newP.dir
		return true
	}

	switch p.dir {
	// face right
	case 0:
		newCol := p.col + 1
		tile := (*mm)[P22{line: p.line, col: newCol}]
		if tile.isWall {
			return false
		}
		p.col = newCol
		return true
	// face down
	case 1:
		newLine := p.line + 1
		tile := (*mm)[P22{line: newLine, col: p.col}]

		if tile.isWall {
			return false
		}
		p.line = newLine
		return true
	// face left
	case 2:
		newCol := p.col - 1
		tile := (*mm)[P22{line: p.line, col: newCol}]

		if tile.isWall {
			return false
		}
		p.col = newCol
		return true

	// face up
	case 3:
		newLine := p.line - 1
		tile := (*mm)[P22{line: newLine, col: p.col}]

		if tile.isWall {
			return false
		}
		p.line = newLine
		return true
	}

	return false
}

func (mm *MonkeyMap) moveOnCube(pos *P22, instr string, cubeWrapper CubeWrapper) {
	if steps, err := strconv.Atoi(instr); err == nil {
		for i := 0; i < steps; i++ {
			if !mm.stepOnCube(pos, cubeWrapper) {
				return
			}
		}
		return
	}

	switch instr {
	case "R":
		pos.dir = (pos.dir + 1) % 4
	case "L":
		if pos.dir == 0 {
			pos.dir = 3
		} else {
			pos.dir--
		}
	}

}

func shrinkTo1x1Cube(mm MonkeyMap) string {
	cube1x1 := ""

	F := byte('A')
	for l := 0; l < 4; l++ {
		for c := 0; c < 4; c++ {
			if _, ok := mm[P22{line: l * 50, col: c * 50}]; ok {
				cube1x1 += string(F)
				F++
			} else {
				cube1x1 += "."
			}
		}
		cube1x1 += "\n"
	}
	return cube1x1
}

type CubeWrapper map[P22]P22

func buildCubeWrapper(mm MonkeyMap) CubeWrapper {
	cubeWrapper := CubeWrapper{}

	cube1x1 := shrinkTo1x1Cube(mm)

	for x := 0; x < 50; x++ {
		oldP := P22{line: 0, col: 50 + x, dir: 3}
		newP := P22{line: 3*50 + x, col: 0, dir: 0}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 0, col: 2*50 + x, dir: 3}
		newP := P22{line: 4*50 - 1, col: x, dir: 3}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: x, col: 3*50 - 1, dir: 0}
		newP := P22{line: 3*50 - 1 - x, col: 2*50 - 1, dir: 2}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 50 - 1, col: 2*50 + x, dir: 1}
		newP := P22{line: 50 + x, col: 2*50 - 1, dir: 2}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 50 + x, col: 2*50 - 1, dir: 0}
		newP := P22{line: 50 - 1, col: 2*50 + x, dir: 3}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 2*50 + x, col: 2*50 - 1, dir: 0}
		newP := P22{line: 50 - 1 - x, col: 3*50 - 1, dir: 2}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 3*50 - 1, col: 50 + x, dir: 1}
		newP := P22{line: 3*50 + x, col: 50 - 1, dir: 2}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 3*50 + x, col: 50 - 1, dir: 0}
		newP := P22{line: 3*50 - 1, col: 50 + x, dir: 3}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 4*50 - 1, col: x, dir: 1}
		newP := P22{line: 0, col: 2*50 + x, dir: 1}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 3*50 + x, col: 0, dir: 2}
		newP := P22{line: 0, col: 50 + x, dir: 1}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 2*50 + x, col: 0, dir: 2}
		newP := P22{line: 50 - 1 - x, col: 50, dir: 0}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 2 * 50, col: x, dir: 3}
		newP := P22{line: 50 + x, col: 50, dir: 0}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: 50 + x, col: 50, dir: 2}
		newP := P22{line: 2 * 50, col: x, dir: 1}
		cubeWrapper[oldP] = newP
	}

	for x := 0; x < 50; x++ {
		oldP := P22{line: x, col: 50, dir: 2}
		newP := P22{line: 3*50 - 1 - x, col: 0, dir: 0}
		cubeWrapper[oldP] = newP
	}

	_ = cube1x1
	return cubeWrapper
}

func Challenge22Part1(inputFile string) int {
	monkeyMap, instructions := parseInput22(inputFile)
	position := findInitialPosition(monkeyMap)

	for _, instruction := range instructions {
		monkeyMap.move(&position, instruction)

	}
	return (position.line+1)*1000 + (position.col+1)*4 + position.dir
}

func Challenge22Part2(inputFile string) int {
	monkeyMap, instructions := parseInput22(inputFile)
	position := findInitialPosition(monkeyMap)

	cubeWrapper := buildCubeWrapper(monkeyMap)

	for _, instruction := range instructions {
		monkeyMap.moveOnCube(&position, instruction, cubeWrapper)
	}
	return (position.line+1)*1000 + (position.col+1)*4 + position.dir
}
