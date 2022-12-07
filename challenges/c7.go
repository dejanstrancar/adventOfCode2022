package challenges

import (
	"adventOfCode2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const MIN_FILE_SIZE = 100_000
const TOTAL_DISK_SPACE = 70_000_000
const REQUIRED_SPACE = 30_000_000

type file struct {
	name   string
	size   int
	isDir  bool
	parent *file
	files  []*file
}

func findCandidates(directory *file, minSize int) []*file {
	dirs := make([]*file, 0)

	if directory.size >= minSize {
		dirs = append(dirs, directory)
	}

	for _, subdir := range directory.files {
		dirs = append(dirs, findCandidates(subdir, minSize)...)
	}

	return dirs
}

func (f *file) addFile(name string, size int) {
	f.files = append(f.files, &file{name: name, parent: f})
	f.increaseSize(size)
}

func (f *file) addDir(name string) {
	f.files = append(f.files, &file{name: name, parent: f, files: []*file{}, isDir: true})
}

func (f *file) increaseSize(size int) {
	f.size += size
	if f.parent != nil {
		f.parent.increaseSize(size)
	}
}

func (f *file) findDir(name string) *file {
	for _, d := range f.files {
		if d.name == name && d.isDir {
			return d
		}
	}
	return nil
}

func findCandidatesFileSize(dir *file) int {
	total := 0

	if dir.size <= MIN_FILE_SIZE {
		total += dir.size
	}

	for _, file := range dir.files {
		if file.isDir {
			total += findCandidatesFileSize(file)
		}
	}

	return total
}

func parseFolderContent(lines []string, i int, current *file) *file {
	line := lines[i]
	pieces := strings.Split(line, " ")

	if pieces[0] == "dir" {
		current.addDir(pieces[1])
	} else {
		size, _ := strconv.Atoi(pieces[0])
		current.addFile(pieces[1], size)
	}

	return current
}

func generateFilemap(lines []string) *file {
	root := &file{name: "/", files: []*file{}}
	current := root

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		action := ""
		fmt.Sscanf(line, "$ cd %s", &action)

		switch {
		case action == "/":
			current = root
		case action == "..":
			current = current.parent
		case action != "":
			current = current.findDir(action)
		case line == "$ ls":
			for {
				if i == len(lines)-1 || lines[i+1][0] == '$' {
					break
				}
				i++
				current = parseFolderContent(lines, i, current)
			}

		}
	}

	return root
}

func Challenge7Part1(inputFile string) int {
	content := utils.LoadFileToArray(inputFile)
	tree := generateFilemap(content)
	return findCandidatesFileSize(tree)
}

func Challenge7Part2(inputFile string) int {
	content := utils.LoadFileToArray(inputFile)
	tree := generateFilemap(content)

	minDirSpace := REQUIRED_SPACE - TOTAL_DISK_SPACE - tree.size
	candidates := findCandidates(tree, minDirSpace)

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].size < candidates[j].size
	})

	return candidates[0].size
}
