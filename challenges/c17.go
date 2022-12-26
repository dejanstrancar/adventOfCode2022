package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
	"strings"
)

type Piece struct {
	width int
	shape [4][4]bool
}

type Count struct {
	height     int
	pieceCount int
}

type Chamber struct {
	pieceCounter int
	moveIdx      int
	height       int
	moves        string
	pieces       []Piece
	state        [][7]bool
	cache        map[string]Count
}

func buildPieces() []Piece {
	return []Piece{
		{
			width: 4,
			shape: [4][4]bool{
				{true, true, true, true},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
		},
		{
			width: 3,
			shape: [4][4]bool{
				{false, true, false, false},
				{true, true, true, false},
				{false, true, false, false},
				{false, false, false, false},
			},
		},
		{
			width: 3,
			shape: [4][4]bool{
				{true, true, true, false},
				{false, false, true, false},
				{false, false, true, false},
				{false, false, false, false},
			},
		},
		{
			width: 1,
			shape: [4][4]bool{
				{true, false, false, false},
				{true, false, false, false},
				{true, false, false, false},
				{true, false, false, false},
			},
		},
		{
			width: 2,
			shape: [4][4]bool{
				{true, true, false, false},
				{true, true, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
		},
	}
}

func (c *Chamber) isColision(pos [2]int, piece Piece) bool {
	if pos[0] < 0 || pos[1] < 0 || pos[1]+piece.width > len(c.state[0]) {
		return true
	}
	for line := 0; line < len(piece.shape); line++ {
		for col := 0; col < len(piece.shape[line]); col++ {
			if piece.shape[line][col] && c.state[pos[0]+line][pos[1]+col] {
				return true
			}
		}
	}
	return false
}
func (c *Chamber) savePiece(pos [2]int, piece Piece) {
	for line := 0; line < len(piece.shape); line++ {
		for col := 0; col < len(piece.shape[line]); col++ {
			if piece.shape[line][col] {
				c.state[pos[0]+line][pos[1]+col] = true

				if pos[0]+line+1 > c.height {
					c.height = pos[0] + line + 1
				}
			}
		}
	}
}

func (c *Chamber) nextPiece() {
	c.pieceCounter++
	pos := [2]int{c.height + 3, 2}
	piece := c.pieces[(c.pieceCounter-1)%len(c.pieces)]

	newLines := (c.height + 7) - len(c.state)
	for i := 0; i < newLines; i++ {
		c.state = append(c.state, [7]bool{})
	}

	for {

		newPos := [2]int{pos[0], pos[1]}
		switch c.moves[c.moveIdx] {
		case '>':
			newPos[1]++
		case '<':
			newPos[1]--
		}

		if !c.isColision(newPos, piece) {
			pos = newPos
		}

		c.moveIdx = (c.moveIdx + 1) % len(c.moves)

		if c.isColision([2]int{pos[0] - 1, pos[1]}, piece) {

			c.savePiece(pos, piece)
			return
		}
		pos[0]--

	}

}

func (c *Chamber) updateSeenState() Count {
	if len(c.state) < 30 {
		return Count{}
	}

	pieceIdx := c.pieceCounter % len(c.pieces)
	stateHash := fmt.Sprintf("%03d:%05d:", pieceIdx, c.moveIdx)
	for l := 0; l < 30; l++ {
		lineHash := 0
		for x := 0; x < 7; x++ {
			if c.state[l][6-x] {
				lineHash |= 1 << x
			}
		}
		stateHash += fmt.Sprintf("%03d", lineHash)
	}

	if x, ok := c.cache[stateHash]; ok {
		return x
	}
	c.cache[stateHash] = Count{height: c.height, pieceCount: c.pieceCounter}
	return Count{}
}

func Challenge17Part1(inputFile string) int {
	content := utils.LoadFileToArray(inputFile)
	chamber := Chamber{
		moves:  strings.Join(content, ""),
		pieces: buildPieces(),
		state:  [][7]bool{},
	}

	for i := 0; i < 2022; i++ {
		chamber.nextPiece()
	}

	return chamber.height
}

func Challenge17Part2(inputFile string) int {
	content := utils.LoadFileToArray(inputFile)

	chamber := Chamber{
		moves:  strings.Join(content, ""),
		pieces: buildPieces(),
		state:  [][7]bool{},
		cache:  make(map[string]Count),
	}

	for i := 0; i < 1000000000000; i++ {
		chamber.nextPiece()

		cacheState := chamber.updateSeenState() // 1

		if cacheState.height > 0 {
			remaining := 1000000000000 - chamber.pieceCounter
			repeatedHeight := chamber.height - cacheState.height
			repeatedLen := chamber.pieceCounter - cacheState.pieceCount
			repeatedTotalHeight := repeatedHeight * (remaining / repeatedLen)

			for j := 0; j < remaining%repeatedLen; j++ {
				chamber.nextPiece()
			}
			chamber.height += repeatedTotalHeight
			break
		}
	}

	return chamber.height
}
