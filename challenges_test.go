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
	fmt.Println("==============")
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
	fmt.Println("==============")

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
	fmt.Println("==============")
}

func TestC4P1(t *testing.T) {
	result := challenges.Challenge4Part1("./inputs/4.txt")
	solution := 500

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C4P1:", result)
}

func TestC4P2(t *testing.T) {
	result := challenges.Challenge4Part2("./inputs/4.txt")
	solution := 815

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C4P2:", result)
	fmt.Println("==============")
}

func TestC5P1(t *testing.T) {
	result := challenges.Challenge5Part1("./inputs/5.txt")
	solution := "GRTSWNJHH"

	if result != solution {
		t.Fatalf("Not expected. %s != %s", result, solution)
	}

	fmt.Println("C5P1:", result)
}

func TestC5P2(t *testing.T) {
	result := challenges.Challenge5Part2("./inputs/5.txt")
	solution := "QLFQDBBHM"

	if result != solution {
		t.Fatalf("Not expected. %s != %s", result, solution)
	}

	fmt.Println("C5P2:", result)
	fmt.Println("==============")
}

func TestC6P1(t *testing.T) {
	result := challenges.Challenge6Part1("./inputs/6.txt")
	solution := 1623

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C6P1:", result)
}

func TestC6P2(t *testing.T) {
	result := challenges.Challenge6Part2("./inputs/6.txt")
	solution := 3774

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C6P2:", result)
	fmt.Println("==============")
}

func TestC7P1(t *testing.T) {
	result := challenges.Challenge7Part1("./inputs/7.txt")
	solution := 1743217

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C7P1:", result)
}

func TestC7P2(t *testing.T) {
	result := challenges.Challenge7Part2("./inputs/7.txt")
	solution := 8319096

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C7P2:", result)
	fmt.Println("==============")
}

func TestC8P1(t *testing.T) {
	result := challenges.Challenge8Part1("./inputs/8.txt")
	solution := 1789

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C8P1:", result)
}

func TestC8P2(t *testing.T) {
	result := challenges.Challenge8Part2("./inputs/8.txt")
	solution := 314820

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C8P2:", result)
	fmt.Println("==============")
}

func TestC9P1(t *testing.T) {
	result := challenges.Challenge9Part1("./inputs/9.txt")
	solution := 6037

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C9P1:", result)
}

func TestC9P2(t *testing.T) {
	result := challenges.Challenge9Part2("./inputs/9.txt")
	solution := 2485

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C9P2:", result)
	fmt.Println("==============")
}

// func TestC10P1(t *testing.T) {
// 	result := challenges.Challenge10Part1("./inputs/10.txt")
// 	solution := 1789

// 	if result != solution {
// 		t.Fatalf("Not expected. %d != %d", result, solution)
// 	}

// 	fmt.Println("C10P1:", result)
// }

// func TestC10P2(t *testing.T) {
// 	result := challenges.Challenge10Part2("./inputs/10.txt")
// 	solution := 314820

// 	if result != solution {
// 		t.Fatalf("Not expected. %d != %d", result, solution)
// 	}

// 	fmt.Println("C10P2:", result)
// 	fmt.Println("==============")
// }

// Test placeholders

// func TestC_P1(t *testing.T) {
// 	result := challenges.Challenge_Part1("./inputs/_.txt")
// 	solution := 0

// 	if result != solution {
// 		t.Fatalf("Not expected. %d != %d", result, solution)
// 	}

// 	fmt.Println("C_P1:", result)
// }

// func TestC_P2(t *testing.T) {
// 	result := challenges.Challenge_Part2("./inputs/_.txt")
// 	solution := 0

// 	if result != solution {
// 		t.Fatalf("Not expected. %d != %d", result, solution)
// 	}

// 	fmt.Println("C_P2:", result)
// }
