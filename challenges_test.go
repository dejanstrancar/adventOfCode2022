package main

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestC1P1(t *testing.T) {
	fmt.Println("C1P1:", challenges.Challenge1Part1("./inputs/1.txt"))
}

func TestC1P2(t *testing.T) {
	fmt.Println("C1P2: ", challenges.Challenge1Part2("./inputs/1.txt"))
}

func TestC2P1(t *testing.T) {
	fmt.Println("C2P1:", challenges.Challenge2Part1("./inputs/2.txt"))
}

func TestC2P2(t *testing.T) {
	fmt.Println("C2P2: ", challenges.Challenge2Part2("./inputs/2.txt"))
}

func TestC3P1(t *testing.T) {
	fmt.Println("C3P1:", challenges.Challenge3Part1("./inputs/3.txt"))
}

func TestC3P2(t *testing.T) {
	fmt.Println("C3P2: ", challenges.Challenge3Part2("./inputs/3.txt"))
}
