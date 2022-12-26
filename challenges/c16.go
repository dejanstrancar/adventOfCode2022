package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
	"sort"
	"strings"
)

type Valves struct {
	flowRates   map[string]int
	connections map[string][]string
}

type Queue []string

type Cache struct {
	totalMinutes     int
	stateMaxPressure map[string]int
	distances        map[string]map[string]int
	valvesFlow       map[string]int
}

func parseInput16(inputFile string) Valves {
	content := utils.LoadFileToArray(inputFile)
	valves := Valves{
		flowRates:   make(map[string]int),
		connections: make(map[string][]string),
	}
	for _, line := range content {
		var label string
		var flowRate int
		var connections []string
		lineParts := strings.Split(line, "; tunnels lead to valves ")
		if len(lineParts) == 1 {
			lineParts = strings.Split(line, "; tunnel leads to valve ")
		}

		fmt.Sscanf(lineParts[0], "Valve %s has flow rate=%d", &label, &flowRate)
		connections = strings.Split(lineParts[1], ", ")

		valves.flowRates[label] = flowRate
		valves.connections[label] = connections

	}
	return valves
}

func contains(ss []string, match string) bool {
	for _, x := range ss {
		if x == match {
			return true
		}
	}
	return false
}

func findNonEmptyValves(valves Valves) map[string]int {
	nonEmpty := map[string]int{}
	for label, flow := range valves.flowRates {
		if flow > 0 {
			nonEmpty[label] = flow
		}
	}
	return nonEmpty
}

func (q *Queue) popLeft() string {
	elem := (*q)[0]
	*q = (*q)[1:]
	return elem
}

func computeDistanceToAll(connections map[string][]string, nonEmpty map[string]int, origin string) map[string]int {
	distances := map[string]int{}
	dist := 0
	q := append(Queue{}, connections[origin]...)
	for dist <= 26 {
		dist++
		tmpQueue := Queue{}
		for len(q) > 0 {
			next := q.popLeft()
			if _, ok := distances[next]; ok || next == origin {
				continue
			}
			if _, ok := nonEmpty[next]; ok {
				distances[next] = dist
			}
			for _, conn := range connections[next] {
				if !contains(tmpQueue, conn) {
					tmpQueue = append(tmpQueue, conn)
				}
			}
		}
		q = append(q, tmpQueue...)
	}
	return distances
}

func computeValvesDistances(connections map[string][]string, nonEmpty map[string]int) map[string]map[string]int {
	distances := map[string]map[string]int{}
	for origin := range nonEmpty {
		distances[origin] = computeDistanceToAll(connections, nonEmpty, origin)
	}

	distances["AA"] = computeDistanceToAll(connections, nonEmpty, "AA")

	return distances
}

func pathToState(s []string) string {
	newS := append([]string{}, s[1:]...)
	sort.Strings(newS)
	return strings.Join(newS, ",")
}

func (c *Cache) calculateTotalPressure(path []string) (voidMinutes int, totalFlow int) {
	totalPressure := 0
	dist := 0

	for i, v := range path[1:] {
		neededMinutes := c.distances[path[i]][v] + 1
		dist += neededMinutes
		totalPressure += (c.totalMinutes - dist) * c.valvesFlow[v]
	}

	return c.totalMinutes - dist, totalPressure
}

func (c *Cache) dfs(path []string, valveState string) int {
	currValve := path[len(path)-1]

	remaining, totalPressure := c.calculateTotalPressure(path)
	if n, ok := c.stateMaxPressure[valveState]; ok && n > totalPressure {
		return n
	}
	c.stateMaxPressure[valveState] = totalPressure

	maxPressure := totalPressure
	for nextValve := range c.distances[currValve] {
		if remaining < 0 {
			break
		}
		if strings.Contains(valveState, nextValve) {
			continue
		}

		newPath := append(path, nextValve)
		newState := pathToState(newPath)
		pressure := c.dfs(newPath, newState)
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}

	return maxPressure
}

func overlap(s1, s2 string) bool {
	for i, j := 0, 0; i <= len(s1)-2 && j <= len(s2)-2; {
		if s1[i:i+2] == s2[j:j+2] {
			return true
		}
		if s1[i:i+2] < s2[j:j+2] {
			i += 3
		} else {
			j += 3
		}
	}
	return false
}

func (c *Cache) findMaxExclusivePair() int {
	max := 0
	for state1, pressure1 := range c.stateMaxPressure {
		for state2, pressure2 := range c.stateMaxPressure {
			if overlap(state1, state2) {
				continue
			}
			if pressure1+pressure2 > max {
				max = pressure1 + pressure2
			}
		}
	}
	return max
}

func Challenge16Part1(inputFile string) int {
	valves := parseInput16(inputFile)
	nonEmptyValves := findNonEmptyValves(valves)
	distances := computeValvesDistances(valves.connections, nonEmptyValves)

	cache := Cache{
		totalMinutes:     30,
		stateMaxPressure: make(map[string]int),
		distances:        distances,
		valvesFlow:       valves.flowRates,
	}
	solution := cache.dfs([]string{"AA"}, "")

	return solution
}

func Challenge16Part2(input string) int {
	valves := parseInput16(input)
	nonEmptyValves := findNonEmptyValves(valves)
	distances := computeValvesDistances(valves.connections, nonEmptyValves)

	cache := Cache{
		totalMinutes:     26,
		stateMaxPressure: make(map[string]int),
		distances:        distances,
		valvesFlow:       valves.flowRates,
	}
	cache.dfs([]string{"AA"}, "")
	solution := cache.findMaxExclusivePair()

	return solution
}
