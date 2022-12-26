package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
)

const decryptionKey = 811589153

type DLLNode struct {
	value int
	next  *DLLNode
	prev  *DLLNode
}

func parseInput20(inputFile string) []*DLLNode {
	content := utils.LoadFileToArray(inputFile)
	numberList := []*DLLNode{}
	var prev *DLLNode = nil
	for _, line := range content {
		var number int
		fmt.Sscanf(line, "%d", &number)
		node := DLLNode{
			value: number,
			prev:  prev,
		}
		numberList = append(numberList, &node)
		prev = &node
	}

	for i, node := range numberList {
		node.next = numberList[(i+1)%len(numberList)]
	}

	numberList[0].prev = numberList[len(numberList)-1]

	return numberList
}

func (n *DLLNode) shiftForward(stopIn int) {
	if stopIn == 0 {
		return
	}

	prev := n.prev
	curr := n
	next := n.next

	curr.next = next.next
	curr.prev = next
	prev.next = next
	next.prev = prev
	next.next.prev = curr
	next.next = curr

	n.shiftForward(stopIn - 1)
}

func (n *DLLNode) shiftBackward(stopIn int) {
	if stopIn == 0 {
		return
	}

	prev := n.prev
	curr := n
	next := n.next

	curr.next = prev
	curr.prev = prev.prev
	prev.next = next
	next.prev = prev
	prev.prev.next = curr
	prev.prev = curr

	n.shiftBackward(stopIn - 1)
}

func (n *DLLNode) mix() {
	value := n.value
	if value >= 0 {
		n.shiftForward(value)
	} else {
		n.shiftBackward(-value)
	}
}

func (n *DLLNode) mixPart2(length int) {
	value := n.value % (length - 1)
	if value >= 0 {
		n.shiftForward(value)
	} else {
		n.shiftBackward(-value)
	}
}

func getSolution(zeroNode *DLLNode, size int) int {
	idx := []int{1000 % size, 2000 % size, 3000 % size}
	solution := 0
	curr := zeroNode
	for i := 0; i < size; i++ {
		if i == idx[0] || i == idx[1] || i == idx[2] {
			solution += curr.value
		}
		curr = curr.next
	}
	return solution
}

func Challenge20Part1(inputFile string) int {
	numberList := parseInput20(inputFile)
	var zeroNode *DLLNode
	for _, node := range numberList {
		if node.value == 0 {
			zeroNode = node
		}
	}

	for _, node := range numberList {
		node.mix()
	}

	solution := getSolution(zeroNode, len(numberList))
	return solution
}

func Challenge20Part2(inputFile string) int {
	numberList := parseInput20(inputFile)

	var zeroNode *DLLNode
	for _, node := range numberList {
		if node.value == 0 {
			zeroNode = node
		}
		node.value *= decryptionKey
	}

	for i := 0; i < 10; i++ {
		for _, node := range numberList {
			node.mixPart2(len(numberList))
		}
	}

	solution := getSolution(zeroNode, len(numberList))
	return solution
}
