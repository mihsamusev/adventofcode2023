package main


type PositionQueue []Pos

func (q *PositionQueue) Enqueue(value Pos) {
	*q = append(*q, value)
}

func (q *PositionQueue) Dequeue() Pos {
	if len(*q) == 0 {
		return Pos{-1, -1}
	}

	p := (*q)[0]
	*q = (*q)[1:]
	return p
}