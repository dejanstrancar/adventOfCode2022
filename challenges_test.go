package main

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestC1P1(t *testing.T) {
	result := challenges.Challenge1Part1("./inputs/1.txt")
	solution := 70509

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C1P1:", result)
}

func TestC1P2(t *testing.T) {
	result := challenges.Challenge1Part2("./inputs/1.txt")
	solution := 208567

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C1P2:", result)
	fmt.Println("")
}

func TestC2P1(t *testing.T) {
	result := challenges.Challenge2Part1("./inputs/2.txt")
	solution := 10718

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C2P1:", result)
}

func TestC2P2(t *testing.T) {
	result := challenges.Challenge2Part2("./inputs/2.txt")
	solution := 14652

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C2P2:", result)
	fmt.Println("")

}

func TestC3P1(t *testing.T) {
	result := challenges.Challenge3Part1("./inputs/3.txt")
	solution := 8233

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C3P1:", result)
}

func TestC3P2(t *testing.T) {
	result := challenges.Challenge3Part2("./inputs/3.txt")
	solution := 2821

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C3P2:", result)
	fmt.Println("")
}
