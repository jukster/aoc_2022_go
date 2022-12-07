package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jukster/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	path := []string{}

	fileList := []file{}

	for _, line := range inputRaw {
		dirUp, enterDir, filename, filesize := processLine(line)

		if dirUp {
			if len(path) > 0 {
				path = path[:len(path)-1]
			}
			continue
		} else if enterDir != "" {
			path = append(path, enterDir)

			continue
		} else if filename != "" {
			thisPath := make([]string, len(path))
			copy(thisPath, path)
			fileList = append(fileList, file{thisPath, filename, filesize})
		}

	}

	directories := map[string]int{}

	for _, file := range fileList {
		for i := range file.path {
			path := ""

			path = strings.Join(file.path[:i+1], "/")

			if _, ok := directories[path]; !ok {
				directories[path] = 0
			}

			directories[path] += file.size

		}
	}

	totalForSmallDirs := 0

	for _, value := range directories {
		if value < 100000 {
			totalForSmallDirs += value
		}
	}

	fmt.Println(totalForSmallDirs)

	fmt.Println("Starting Day 2")

	used := directories["/"]

	maxOccupied := 40000000

	biggestAfterDeletion := 0

	winnerSize := 0

	for _, size := range directories {
		usedIfDeleted := used - size

		if usedIfDeleted < maxOccupied {
			if usedIfDeleted > biggestAfterDeletion {
				biggestAfterDeletion = usedIfDeleted
				winnerSize = size
			}

		}

	}

	fmt.Println(winnerSize)

}

type file struct {
	path []string
	name string
	size int
}

func processLine(line string) (dirUp bool, enterDir string, filename string, filesize int) {

	patCmdUp := regexp.MustCompile(`^\$\scd\s\.\.`)
	patEnterDir := regexp.MustCompile(`^\$\scd\s(.+)`)
	patFil := regexp.MustCompile(`^(\d+)\s(.+)`)

	if m := patCmdUp.FindAllString(line, -1); len(m) > 0 {

		return true, "", "", 0

	} else if m := patEnterDir.FindAllStringSubmatch(line, -1); len(m) > 0 {

		return false, m[0][1], "", 0

	} else if m := patFil.FindAllStringSubmatch(line, -1); len(m) > 0 {
		newFileName := m[0][2]
		newFileSizeS := m[0][1]

		newFileSize, err := strconv.Atoi(newFileSizeS)

		if err != nil {
			fmt.Println(fmt.Errorf("conversion failed"))
		}

		return false, "", newFileName, newFileSize
	}

	return false, "", "", 0
}
