package main

type round struct {
	oppoMove string
	myMove   string
	outcome  string
}

func (r *round) eval() int {

	choiceScore := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	score := choiceScore[r.myMove]

	score = score + play(r.oppoMove, r.myMove)

	return score

}

func (r *round) findMove() {

	moveMap := map[string]string{
		"Rock-Lose":     "Scissors",
		"Rock-Draw":     "Rock",
		"Rock-Win":      "Paper",
		"Paper-Lose":    "Rock",
		"Paper-Draw":    "Paper",
		"Paper-Win":     "Scissors",
		"Scissors-Lose": "Paper",
		"Scissors-Draw": "Scissors",
		"Scissors-Win":  "Rock",
	}

	key := r.oppoMove + "-" + r.outcome

	r.myMove = moveMap[key]
}
