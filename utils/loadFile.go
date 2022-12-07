package utils

import (
	"bufio"
	"os"
)

func LoadFileToArray(filename string) []string {
	f, _ := os.Open(filename)
	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
