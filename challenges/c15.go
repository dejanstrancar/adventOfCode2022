package challenges

import (
	"bufio"
	"fmt"
	"os"
)

type point5 struct {
	x, y int
}

func manhattan(sensorX, sensorY, beaconX, beaconY int) int {
	x := sensorX - beaconX
	y := sensorY - beaconY
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func Challenge15Part1(inputFile string) int {

	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	line2000000 := make(map[int]bool)

	for sc.Scan() {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		distanceFromBeacon := manhattan(sensorX, sensorY, beaconX, beaconY)

		distanceFromLine := (sensorY - 2000000)
		if distanceFromLine < 0 {
			distanceFromLine *= -1
		}

		for i := 0; i <= distanceFromBeacon-distanceFromLine; i++ {
			line2000000[sensorX+i] = true
			line2000000[sensorX-i] = true
		}

		if beaconY == 2000000 {
			delete(line2000000, beaconX)
		}
	}
	return len(line2000000)
}

func Challenge15Part2(inputFile string) int {

	input, _ := os.Open(inputFile)
	defer input.Close()
	sc := bufio.NewScanner(input)

	const limit int = 4000000

	toTry := make(map[point5]bool)
	nearestBeacon := make(map[point5]int)

	for sc.Scan() {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		distanceFromBeacon := manhattan(sensorX, sensorY, beaconX, beaconY)
		nearestBeacon[point5{sensorX, sensorY}] = distanceFromBeacon
		distanceFromBeacon++
		for i := 0; i < distanceFromBeacon; i++ {
			if sensorX+i > 0 && sensorX+i < limit {
				if sensorY-distanceFromBeacon+1+i > 0 && sensorY-distanceFromBeacon+1+i < limit {
					toTry[point5{sensorX + i, sensorY - distanceFromBeacon + 1 + i}] = true
				}
				if sensorY+distanceFromBeacon-1-i > 0 && sensorY+distanceFromBeacon-1-i < limit {
					toTry[point5{sensorX + i, sensorY + distanceFromBeacon - i}] = true
				}
			}
			if sensorX-i > 0 && sensorX-i < limit {
				if sensorY-distanceFromBeacon+1+i > 0 && sensorY-distanceFromBeacon+1+i < limit {
					toTry[point5{sensorX - i, sensorY - distanceFromBeacon + 1 + i}] = true
				}
				if sensorY+distanceFromBeacon-1-i > 0 && sensorY+distanceFromBeacon-1-i < limit {
					toTry[point5{sensorX - i, sensorY + distanceFromBeacon - 1 - i}] = true
				}
			}
		}
	}

	for beacon := range toTry {
		newBeacon := true
		for sensor, nearestBeaconDistance := range nearestBeacon {
			if manhattan(sensor.x, sensor.y, beacon.x, beacon.y) <= nearestBeaconDistance {
				newBeacon = false
				break
			}
		}
		if newBeacon {
			return beacon.x*limit + beacon.y
		}
	}

	return 0
}
