package main

import (
	"sort"

	"golang.org/x/exp/slices"
)

type queue struct {
	unfinished []*node
}

func (q *queue) sortQueue() {
	sort.Slice(q.unfinished, func(i, j int) bool {
		return q.unfinished[i].distanceFromStart < q.unfinished[j].distanceFromStart
	})
}

func (q *queue) getNext() (next *node) {
	if len(q.unfinished) > 0 {
		next = q.unfinished[0]
		q.unfinished = q.unfinished[1:]
	}
	return next
}

func (q *queue) add(new *node) {
	if !slices.Contains(q.unfinished, new) {
		q.unfinished = append(q.unfinished, new)
	}
	q.sortQueue()
}
