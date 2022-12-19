package challenges

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type tree struct {
	valueLeaf int
	elements  []*tree
	father    *tree
}

func readTree(input string) tree {
	root := tree{-1, []*tree{}, nil}
	temp := &root

	var current string
	for _, r := range input {
		switch r {
		case '[':
			newTree := tree{-1, []*tree{}, temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		case ']':
			if len(current) > 0 {
				number, _ := strconv.Atoi(current)
				temp.valueLeaf = number
				current = ""
			}
			temp = temp.father
		case ',':
			if len(current) > 0 {
				number, _ := strconv.Atoi(current)
				temp.valueLeaf = number
				current = ""
			}
			temp = temp.father
			newTree := tree{-1, []*tree{}, temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		default:
			current += string(r)
		}
	}
	return root
}

func areOrdered(first, second tree) int {
	switch {
	case len(first.elements) == 0 && len(second.elements) == 0:
		if first.valueLeaf > second.valueLeaf {
			return -1
		} else if first.valueLeaf == second.valueLeaf {
			return 0
		}
		return 1

	case first.valueLeaf >= 0:
		return areOrdered(tree{-1, []*tree{&first}, nil}, second)

	case second.valueLeaf >= 0:
		return areOrdered(first, tree{-1, []*tree{&second}, nil})
	default:
		var i int
		for i = 0; i < len(first.elements) && i < len(second.elements); i++ {
			ordered := areOrdered(*first.elements[i], *second.elements[i])
			if ordered != 0 {
				return ordered
			}
		}
		if i < len(first.elements) {
			return -1
		} else if i < len(second.elements) {
			return 1
		}
	}
	return 0
}

func Challenge13Part1(inputFile string) int {

	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	index := 1
	var indexSum int
	for sc.Scan() {
		package1 := readTree(sc.Text())
		sc.Scan()
		package2 := readTree(sc.Text())

		if areOrdered(package1, package2) == 1 {
			indexSum += index
		}

		index++
		sc.Scan()
	}
	return indexSum
}

func Challenge13Part2(inputFile string) int {

	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var packeges []tree
	for sc.Scan() {
		packeges = append(packeges, readTree(sc.Text()))
		sc.Scan()
		packeges = append(packeges, readTree(sc.Text()))
		sc.Scan()
	}
	packeges = append(packeges, readTree("[[2]]"))
	packeges = append(packeges, readTree("[[6]]"))

	sort.Slice(packeges, func(i, j int) bool {
		return areOrdered(packeges[i], packeges[j]) == 1
	})

	decoderKey := 1
	for i, p := range packeges {
		if areOrdered(p, readTree("[[2]]")) == 0 || areOrdered(p, readTree("[[6]]")) == 0 {
			decoderKey *= i + 1
		}
	}

	return decoderKey
}
