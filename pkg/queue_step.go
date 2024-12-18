package pkg

import "slices"

func NewStepQueue() StepQueue {
	return StepQueue{q: make([]Step, 0, 50)}
}

type StepQueue struct {
	q []Step
}

func (q *StepQueue) Push(s Step) *StepQueue {
	q.q = append(q.q, s)
	return q
}

func (q *StepQueue) SortAsc() {
	slices.SortFunc(q.q, func(a, b Step) int {
		return a.score - b.score
	})
}

func (q *StepQueue) Len() int {
	return len(q.q)
}

func (q *StepQueue) Get() Step {
	val := q.q[0]
	q.q = q.q[1:]
	return val
}
