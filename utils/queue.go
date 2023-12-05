package utils

import "container/heap"

type PriorityQueueItem[T any] struct {
	v T
	p int
}

type PriorityQueue[T any] []PriorityQueueItem[T]

func (q PriorityQueue[_]) Len() int           { return len(q) }
func (q PriorityQueue[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PriorityQueue[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PriorityQueue[T]) Push(x any)        { *q = append(*q, x.(PriorityQueueItem[T])) }
func (q *PriorityQueue[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PriorityQueue[T]) GPush(v T, p int)  { heap.Push(q, PriorityQueueItem[T]{v, p}) }
func (q *PriorityQueue[T]) GPop() (T, int)    { x := heap.Pop(q).(PriorityQueueItem[T]); return x.v, x.p }
