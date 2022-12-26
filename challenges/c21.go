package challenges

import (
	"adventOfCode2022/utils"
	"strconv"
	"strings"
)

var operators map[string]*Oper

type Oper func(a, b int) int

type Monkey struct {
	name           string
	rawMonkeyValue string
	operator       string
	value          *int
	left           *Monkey
	right          *Monkey
}

func init() {
	var oSum Oper = func(a, b int) int { return a + b }
	var oSub Oper = func(a, b int) int { return a - b }
	var oMul Oper = func(a, b int) int { return a * b }
	var oDiv Oper = func(a, b int) int { return a / b }
	operators = map[string]*Oper{
		"+": &oSum,
		"-": &oSub,
		"*": &oMul,
		"/": &oDiv,
	}

}

func parseInput21(inputFile string) map[string]*Monkey {
	monkeys := make(map[string]*Monkey, 0)
	content := utils.LoadFileToArray(inputFile)
	for _, line := range content {
		parts := strings.Split(line, ": ")
		monkey := Monkey{
			name:           parts[0],
			rawMonkeyValue: parts[1],
		}
		monkeys[monkey.name] = &monkey
	}
	return monkeys
}

func buildTree(monkeys map[string]*Monkey, m *Monkey) {
	if m.value != nil || m.operator != "" {
		return
	}

	if x, err := strconv.Atoi(m.rawMonkeyValue); err == nil {
		m.value = &x
		return
	}

	parts := strings.Split(m.rawMonkeyValue, " ")
	m.left = monkeys[parts[0]]
	m.right = monkeys[parts[2]]
	m.operator = parts[1]

	buildTree(monkeys, m.left)
	buildTree(monkeys, m.right)
}

func (m *Monkey) getValue() int {
	if m.value != nil {
		return *m.value
	}

	x := (*operators[m.operator])(m.left.getValue(), m.right.getValue())
	m.value = &x
	return *m.value
}

func (m *Monkey) traverseToDestination(destination string, wantedValue int) (int, bool) {
	var newWanted int
	if m.name == destination {
		return wantedValue, true
	}

	if m.left != nil {
		switch m.operator {
		case "+":
			// R = a+b | a=R-b
			newWanted = wantedValue - *m.right.value
		case "-":
			// R = a-b | a=R+b
			newWanted = wantedValue + *m.right.value
		case "*":
			// R = a*b | a=R/b
			newWanted = wantedValue / *m.right.value
		case "/":
			// R = a/b | a=R*b
			newWanted = wantedValue * *m.right.value
		}

		if v, ok := m.left.traverseToDestination(destination, newWanted); ok {
			return v, true
		}
	}

	if m.right != nil {
		switch m.operator {
		case "+":
			// R = a+b | b=R-a
			newWanted = wantedValue - *m.left.value
		case "-":
			// R = a-b | b=a-R
			newWanted = *m.left.value - wantedValue
		case "*":
			// R = a*b | b=R/a
			newWanted = wantedValue / *m.left.value
		case "/":
			// R = a/b | b=a/R
			newWanted = *m.left.value / wantedValue
		}
		if v, ok := m.right.traverseToDestination(destination, newWanted); ok {
			return v, true
		}
	}

	return 0, false
}

func Challenge21Part1(inputFile string) int {
	monkeys := parseInput21(inputFile)
	buildTree(monkeys, monkeys["root"])
	return monkeys["root"].getValue()
}

func Challenge21Part2(inputFile string) int {
	monkeys := parseInput21(inputFile)
	buildTree(monkeys, monkeys["root"])
	monkeys["root"].getValue()

	if v, ok := monkeys["root"].left.traverseToDestination("humn", *monkeys["root"].right.value); ok {
		return v
	}

	if v, ok := monkeys["root"].right.traverseToDestination("humn", *monkeys["root"].left.value); ok {
		return v
	}

	return 0
}
