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

func TestC10P1(t *testing.T) {
	result := challenges.Challenge10Part1("./inputs/10.txt")
	solution := 12840

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C10P1:", result)
}

func TestC10P2(t *testing.T) {
	result := challenges.Challenge10Part2("./inputs/10.txt")
	solution := 12840

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C10P2:", result)
}

func TestC11P1(t *testing.T) {
	result := challenges.Challenge11Part1("./inputs/11.txt")
	solution := 58065

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C11P1:", result)
}

func TestC11P2(t *testing.T) {
	result := challenges.Challenge11Part2("./inputs/11.txt")
	solution := 16161909698

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C11P2:", result)
}

func TestC12P1(t *testing.T) {
	result := challenges.Challenge12Part1("./inputs/12.txt")
	solution := 425

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C12P1:", result)
}

func TestC12P2(t *testing.T) {
	result := challenges.Challenge12Part2("./inputs/12.txt")
	solution := 418

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C12P2:", result)
}

func TestC13P1(t *testing.T) {
	result := challenges.Challenge13Part1("./inputs/13.txt")
	solution := 5717

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C13P1:", result)
}

func TestC13P2(t *testing.T) {
	result := challenges.Challenge13Part2("./inputs/13.txt")
	solution := 25935

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C13P2:", result)
}

func TestC14P1(t *testing.T) {
	result := challenges.Challenge14Part1("./inputs/14.txt")
	solution := 862

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C14P1:", result)
}

func TestC14P2(t *testing.T) {
	result := challenges.Challenge14Part2("./inputs/14.txt")
	solution := 28744

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C14P2:", result)
}

func TestC15P1(t *testing.T) {
	result := challenges.Challenge15Part1("./inputs/15.txt")
	solution := 5299855

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C15P1:", result)
}

func TestC15P2(t *testing.T) {
	result := challenges.Challenge15Part2("./inputs/15.txt")
	solution := 13615843289729

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C15P2:", result)
}

func TestC16P1(t *testing.T) {
	result := challenges.Challenge16Part1("./inputs/16.txt")
	solution := 1741

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C16P1:", result)
}

func TestC16P2(t *testing.T) {
	result := challenges.Challenge16Part2("./inputs/16.txt")
	solution := 2316

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C16P2:", result)
}

func TestC17P1(t *testing.T) {
	result := challenges.Challenge17Part1("./inputs/17.txt")
	solution := 3130

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C17P1:", result)
}

func TestC17P2(t *testing.T) {
	result := challenges.Challenge17Part2("./inputs/17.txt")
	solution := 1568604651130

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C17P2:", result)
}

func TestC18P1(t *testing.T) {
	result := challenges.Challenge18Part1("./inputs/18.txt")
	solution := 4320

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C18P1:", result)
}

func TestC18P2(t *testing.T) {
	result := challenges.Challenge18Part2("./inputs/18.txt")
	solution := 2456

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C18P2:", result)
}

func TestC19P1(t *testing.T) {
	result := challenges.Challenge19Part1("./inputs/19.txt")
	solution := 1550

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C19P1:", result)
}

func TestC19P2(t *testing.T) {
	result := challenges.Challenge19Part2("./inputs/19.txt")
	solution := 18630

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C19P2:", result)
}

func TestC20P1(t *testing.T) {
	result := challenges.Challenge20Part1("./inputs/20.txt")
	solution := 2275

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C20P1:", result)
}

func TestC20P2(t *testing.T) {
	result := challenges.Challenge20Part2("./inputs/20.txt")
	solution := 4090409331120

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C20P2:", result)
}

func TestC21P1(t *testing.T) {
	result := challenges.Challenge21Part1("./inputs/21.txt")
	solution := 54703080378102

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C21P1:", result)
}

func TestC21P2(t *testing.T) {
	result := challenges.Challenge21Part2("./inputs/21.txt")
	solution := 3952673930912

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C21P2:", result)
}

func TestC22P1(t *testing.T) {
	result := challenges.Challenge22Part1("./inputs/22.txt")
	solution := 186128

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C22P1:", result)
}

func TestC22P2(t *testing.T) {
	result := challenges.Challenge22Part2("./inputs/22.txt")
	solution := 34426

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C22P2:", result)
}

func TestC23P1(t *testing.T) {
	result := challenges.Challenge23Part1("./inputs/23.txt")
	solution := 4109

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C23P1:", result)
}

func TestC23P2(t *testing.T) {
	result := challenges.Challenge23Part2("./inputs/23.txt")
	solution := 1055

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C23P2:", result)
}

func TestC24P1(t *testing.T) {
	result := challenges.Challenge24Part1("./inputs/24.txt")
	solution := 279

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C24P1:", result)
}

func TestC24P2(t *testing.T) {
	result := challenges.Challenge24Part2("./inputs/24.txt")
	solution := 762

	if result != solution {
		t.Fatalf("Not expected. %d != %d", result, solution)
	}

	fmt.Println("C24P2:", result)
}

func TestC25P1(t *testing.T) {
	result := challenges.Challenge25Part1("./inputs/25.txt")
	solution := "2-1-110-=01-1-0-0==2"

	if result != solution {
		t.Fatalf("Not expected. %s != %s", result, solution)
	}

	fmt.Println("C25P1:", result)
}
