package queues_test

import (
	"data_structures/queues"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	q := queues.Empty().IsEmpty()
	if q != true {
		t.Errorf("q = %v; want true", q)
	}
}

func TestEnqueue(t *testing.T) {
	q := queues.Empty()
	q.Enqueue("katze")
	if q.IsEmpty() == true {
		t.Errorf("expected q to be non empty")
	}
}

func TestDequeue(t *testing.T) {
	q := queues.Empty()
	q.Enqueue("katze")
	item := q.Dequeue()
	if item != "katze" {
		t.Errorf("expected item to be katze but got %v", item)
	}

	if !q.IsEmpty() {
		t.Errorf("expected q to be empty but got %v", q)
	}
}

func TestPeek(t *testing.T) {
	q := queues.Empty()
	q.Enqueue("katze")
	if q.Peek() != "katze" {
		t.Errorf("expected katze")
	}

	if q.IsEmpty() {
		t.Errorf("expected non empty queue")
	}
}
