package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jukster/aoc_2022_go/aocutil"
)

type folder struct {
	size     int
	name     string
	parent   *folder
	children []*folder
}

func (f *folder) getSize() int {
	totalSize := f.size

	for _, child := range f.children {
		totalSize += child.getSize()
	}

	return totalSize
}

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	root := folder{0, "/", nil, nil}

	folders := []*folder{&root}

	currentFolder := folders[0]

	for _, line := range inputRaw[1:] {
		dirUp, enterDir, filesize := processLine(line)

		if dirUp {
			currentFolder = currentFolder.parent
			continue

		} else if enterDir != "" {
			newFolder := folder{0, enterDir, currentFolder, nil}
			currentFolder.children = append(currentFolder.children, &newFolder)
			currentFolder = &newFolder
			folders = append(folders, &newFolder)
			continue

		} else if filesize != 0 {
			currentFolder.size += filesize
		}
	}

	totalForSmallDirs := 0

	for _, folder := range folders {

		if size := folder.getSize(); size < 100000 {
			totalForSmallDirs += size
		}
	}

	fmt.Println(totalForSmallDirs)

	fmt.Println("Starting Day 2")

	used := folders[0].getSize()

	maxOccupied := 40000000

	biggestAfterDeletion := 0

	winnerSize := 0

	for _, folder := range folders {
		usedIfDeleted := used - folder.getSize()

		if usedIfDeleted < maxOccupied {
			if usedIfDeleted > biggestAfterDeletion {

				biggestAfterDeletion = usedIfDeleted

				winnerSize = folder.getSize()
			}

		}

	}

	fmt.Println(winnerSize)

}

func processLine(line string) (dirUp bool, enterDir string, filesize int) {

	patCmdUp := regexp.MustCompile(`^\$\scd\s\.\.`)
	patEnterDir := regexp.MustCompile(`^\$\scd\s(.+)`)
	patFil := regexp.MustCompile(`^(\d+)\s(.+)`)

	if m := patCmdUp.FindAllString(line, -1); len(m) > 0 {

		return true, "", 0

	} else if m := patEnterDir.FindAllStringSubmatch(line, -1); len(m) > 0 {

		return false, m[0][1], 0

	} else if m := patFil.FindAllStringSubmatch(line, -1); len(m) > 0 {
		newFileSizeS := m[0][1]

		newFileSize, err := strconv.Atoi(newFileSizeS)

		if err != nil {
			fmt.Println(fmt.Errorf("conversion failed"))
		}

		return false, "", newFileSize
	}

	return false, "", 0
}
