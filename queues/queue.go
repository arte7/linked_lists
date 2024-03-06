package queues

import "data_structures/d_list"

type Queue struct {
	queue *d_list.List
}

func Empty() Queue {
	return Queue{queue: d_list.Empty()}
}

func (q Queue) IsEmpty() bool {
	return q.queue.IsEmpty()
}

// ...interface{} ...bool ......<< variable anzahl an elementen
func (q Queue) Enqueue(item interface{}) {
	q.queue.PushBack(item)
}

func (q Queue) Dequeue() interface{} {
	return q.queue.PopFront()
}

func (q Queue) Peek() interface{} {
	return q.queue.Front()
}
