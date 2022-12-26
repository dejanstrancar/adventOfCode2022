package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
)

const (
	Ore      int = 0
	Clay     int = 1
	Obsidian int = 2
	Geode    int = 3
)

type Blueprint [4][4]int

type State struct {
	maxMinutes int
	blueprint  Blueprint
	cache      map[string]int
	maxRate    [4]int
}
type Resources struct {
	ores     int
	clay     int
	obsidian int
	geodes   int
}

func parseInput19(inputFile string) []Blueprint {
	blueprints := []Blueprint{}
	content := utils.LoadFileToArray(inputFile)
	for _, line := range content {
		var id, oreOre, clayOre, obsidianOre, obsidianClay, geodeOre, geodeObsidian int
		fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &id, &oreOre, &clayOre, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)

		bprint := Blueprint{
			{oreOre, 0, 0, 0},
			{clayOre, 0, 0, 0},
			{obsidianOre, obsidianClay, 0, 0},
			{geodeOre, 0, geodeObsidian, 0},
		}
		blueprints = append(blueprints, bprint)
	}
	return blueprints
}

func max(a ...int) int {
	max := a[0]
	for _, x := range a {
		if x > max {
			max = x
		}
	}
	return max
}

func setupState(b Blueprint) State {
	return State{
		maxMinutes: 24,
		blueprint:  b,
		cache:      make(map[string]int),
		maxRate: [4]int{
			max(b[Ore][Ore], b[Clay][Ore], b[Obsidian][Ore], b[Geode][Ore]),
			max(b[Ore][Clay], b[Clay][Clay], b[Obsidian][Clay], b[Geode][Clay]),
			max(b[Ore][Obsidian], b[Clay][Obsidian], b[Obsidian][Obsidian], b[Geode][Obsidian]),
			max(b[Ore][Geode], b[Clay][Geode], b[Obsidian][Geode], b[Geode][Geode]),
		},
	}
}

func (r Resources) duplicate() Resources {
	return Resources{
		ores:     r.ores,
		clay:     r.clay,
		obsidian: r.obsidian,
		geodes:   r.geodes,
	}
}

func miningResources(resources, robots Resources) Resources {
	return Resources{
		ores:     resources.ores + robots.ores,
		clay:     resources.clay + robots.clay,
		obsidian: resources.obsidian + robots.obsidian,
		geodes:   resources.geodes + robots.geodes,
	}
}

func (state *State) dfs(minute int, resources, robots, doNotBuild Resources) int {
	if minute >= state.maxMinutes {
		return resources.geodes
	}

	nextDoNotBuild := Resources{}
	if resources.ores >= state.blueprint[Geode][Ore] &&
		resources.obsidian >= state.blueprint[Geode][Obsidian] {
		newResources := miningResources(resources, robots)
		newRobots := robots.duplicate()
		newRobots.geodes++
		newResources.ores -= state.blueprint[Geode][Ore]
		newResources.obsidian -= state.blueprint[Geode][Obsidian]
		return state.dfs(minute+1, newResources, newRobots, Resources{})
	}

	maxGeodes := 0

	if doNotBuild.ores == 0 &&
		resources.ores >= state.blueprint[Ore][Ore] && robots.ores < state.maxRate[Ore] {

		newResources := miningResources(resources, robots)
		newRobots := robots.duplicate()
		newRobots.ores++
		newResources.ores -= state.blueprint[Ore][Ore]
		newMaxGeodes := state.dfs(minute+1, newResources, newRobots, Resources{})
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
		nextDoNotBuild.ores++
	}

	if doNotBuild.clay == 0 &&
		resources.ores >= state.blueprint[Clay][Ore] && robots.clay < state.maxRate[Clay] {

		newResources := miningResources(resources, robots)
		newRobots := robots.duplicate()
		newRobots.clay++
		newResources.ores -= state.blueprint[Clay][Ore]
		newMaxGeodes := state.dfs(minute+1, newResources, newRobots, Resources{})
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
		nextDoNotBuild.clay++
	}

	if doNotBuild.obsidian == 0 &&
		resources.ores >= state.blueprint[Obsidian][Ore] &&

		resources.clay >= state.blueprint[Obsidian][Clay] &&
		robots.obsidian < state.maxRate[Obsidian] {
		newResources := miningResources(resources, robots)
		newRobots := robots.duplicate()
		newRobots.obsidian++
		newResources.ores -= state.blueprint[Obsidian][Ore]
		newResources.clay -= state.blueprint[Obsidian][Clay]
		newMaxGeodes := state.dfs(minute+1, newResources, newRobots, Resources{})
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
		nextDoNotBuild.obsidian++
	}

	resources = miningResources(resources, robots)
	newMaxGeodes := state.dfs(minute+1, resources, robots, nextDoNotBuild)
	if newMaxGeodes > maxGeodes {
		maxGeodes = newMaxGeodes
	}

	return maxGeodes
}

func Challenge19Part1(inputFile string) int {
	bprints := parseInput19(inputFile)

	result := 0
	for id, bprint := range bprints {
		st := setupState(bprint)
		geodes := st.dfs(0, Resources{}, Resources{ores: 1}, Resources{})

		qLevel := (id + 1) * geodes
		result += qLevel
	}
	return result
}

func Challenge19Part2(inputFile string) int {
	bprints := parseInput19(inputFile)
	if len(bprints) > 3 {
		bprints = bprints[0:3]
	}

	result := 1
	for _, bprint := range bprints {
		st := setupState(bprint)
		st.maxMinutes = 32
		geodes := st.dfs(0, Resources{}, Resources{ores: 1}, Resources{})

		result *= geodes
	}
	return result
}
