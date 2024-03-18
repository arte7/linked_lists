package prio_queues

import "data_structures/s_list"

type PrioQueue struct {
	queue *s_list.List
}

func Empty() PrioQueue {
	return PrioQueue{queue: s_list.Empty()}
}

func (q PrioQueue) IsEmpty() bool {
	return q.queue.IsEmpty()
}

func (q PrioQueue) Enqueue(item interface{}, prio int32) {
	q.queue.SortedInsert(item, prio)
}

func (q PrioQueue) Dequeue() interface{} {
	return q.queue.PopFront()
}
