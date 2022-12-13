package main

type Queue struct {
	data []*Coord
}

func (q *Queue) Enqueue(item *Coord) {
	q.data = append(q.data, item)
}

func (q *Queue) Dequeue() *Coord {
	if len(q.data) == 0 {
		return nil
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}
