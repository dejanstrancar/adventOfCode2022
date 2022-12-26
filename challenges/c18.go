package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
)

type Cube struct {
	x int
	y int
	z int
}

type Node struct {
	visited bool
	isCube  bool
}

type CubesMap [25][25][25]Node

func parseInput18(inputFile string) []Cube {
	cubes := []Cube{}
	content := utils.LoadFileToArray(inputFile)
	for _, line := range content {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		cubes = append(cubes, Cube{x, y, z})
	}
	return cubes
}

func buildCubesMap(cubes []Cube) (map[Cube]bool, int) {
	cubesMap := map[Cube]bool{}
	area := 0
	for _, cube := range cubes {
		cubesMap[cube] = true
		area += 6

		// check neighbors
		// up
		if cubesMap[Cube{cube.x, cube.y, cube.z + 1}] {
			area -= 2
		}
		// down
		if cubesMap[Cube{cube.x, cube.y, cube.z - 1}] {
			area -= 2
		}
		// front
		if cubesMap[Cube{cube.x + 1, cube.y, cube.z}] {
			area -= 2
		}
		// back
		if cubesMap[Cube{cube.x - 1, cube.y, cube.z}] {
			area -= 2
		}
		// left
		if cubesMap[Cube{cube.x, cube.y + 1, cube.z}] {
			area -= 2
		}
		// right
		if cubesMap[Cube{cube.x, cube.y - 1, cube.z}] {
			area -= 2
		}
	}
	return cubesMap, area
}

func buildMap3D(cubes []Cube) CubesMap {
	myMap := [25][25][25]Node{}

	for _, c := range cubes {
		myMap[c.x+1][c.y+1][c.z+1].isCube = true
	}
	return myMap
}

func (m *CubesMap) dfsSearch(c Cube) {
	m[c.x][c.y][c.z].visited = true
	L := len(m)

	for _, p := range []Cube{
		{c.x, c.y, c.z + 1},
		{c.x, c.y, c.z - 1},
		{c.x, c.y + 1, c.z},
		{c.x, c.y - 1, c.z},
		{c.x + 1, c.y, c.z},
		{c.x - 1, c.y, c.z},
	} {
		if p.x < 0 || p.x >= L || p.y < 0 || p.y >= L || p.z < 0 || p.z >= L {
			continue
		}
		if !m[p.x][p.y][p.z].isCube && !m[p.x][p.y][p.z].visited {
			m.dfsSearch(p)
		}
	}
}

func (m *CubesMap) area() int {
	area := 0
	L := len(m)

	for x := range m {
		for y := range m[x] {
			for z := range m[x][y] {
				if !m[x][y][z].isCube {
					continue
				}

				for _, p := range []Cube{
					{x, y, z + 1},
					{x, y, z - 1},
					{x, y + 1, z},
					{x, y - 1, z},
					{x + 1, y, z},
					{x - 1, y, z},
				} {
					if p.x < 0 || p.x >= L || p.y < 0 || p.y >= L || p.z < 0 || p.z >= L {
						continue
					}
					if m[p.x][p.y][p.z].visited {
						area += 1
					}
				}
			}
		}
	}
	return area
}

func Challenge18Part1(inputFile string) int {
	cubes := parseInput18(inputFile)
	_, area := buildCubesMap(cubes)
	return area
}

func Challenge18Part2(inputFile string) int {
	cubes := parseInput18(inputFile)
	cubesMap := buildMap3D(cubes)

	cubesMap.dfsSearch(Cube{0, 0, 0})
	area := cubesMap.area()

	return area
}
