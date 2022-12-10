package aocutil

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ScanFile(f *os.File) *bufio.Scanner {

	scanner := bufio.NewScanner(f)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scanner
}

func ReadAndSplit(s *bufio.Scanner) []string {
	lines := []string{}

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func OpenFile(f string) *os.File {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func ReadFile(path string) []string {
	f := OpenFile(path)

	defer f.Close()

	s := ScanFile(f)

	c := ReadAndSplit(s)

	return c
}

func ConvStrToInt(in []string) (out []int) {
	for _, s := range in {
		num, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		out = append(out, num)
	}

	return out
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
