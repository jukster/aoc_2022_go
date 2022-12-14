package main

import (
	"fmt"

	"github.com/jukster/aoc_2022_go/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	rounds := []round{}

	for _, line := range inputRaw {
		r := newRound(string(line[0]), string(line[2]))
		rounds = append(rounds, r)

	}

	totalScore := 0

	for _, round := range rounds {
		totalScore += round.eval()
	}

	fmt.Println(totalScore)

	fmt.Println("Day 2!")

	rounds2 := []round{}

	for _, line := range inputRaw {
		r := newRoundWithOutcome(string(line[0]), string(line[2]))
		rounds2 = append(rounds2, r)
	}

	totalScore2 := 0

	for _, round := range rounds2 {
		round.findMove()
		totalScore2 += round.eval()
	}

	fmt.Println(totalScore2)

}

func play(oppoMove string, myMove string) int {
	score := 0
	if oppoMove == "Rock" {
		if myMove == "Paper" {
			score = 6
		} else if myMove == "Rock" {
			score = 3
		}
	} else if oppoMove == "Paper" {
		if myMove == "Scissors" {
			score = 6
		} else if myMove == "Paper" {
			score = 3
		}
	} else {
		if myMove == "Rock" {
			score = 6
		} else if myMove == "Scissors" {
			score = 3
		}
	}

	return score
}

func newRound(oppoMove string, myMove string) round {
	translate := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	return round{translate[oppoMove], translate[myMove], ""}
}

func newRoundWithOutcome(oppoMove string, outcome string) round {
	translate := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	translateOutcome := map[string]string{
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}

	return round{translate[oppoMove], "", translateOutcome[outcome]}
}
