package prio_queues_test

import (
	"data_structures/prio_queues"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	q := prio_queues.Empty().IsEmpty()
	if q != true {
		t.Errorf("q = %v; want true", q)
	}
}

func TestEnqueue(t *testing.T) {
	q := prio_queues.Empty()
	q.Enqueue("murmel", 2)

	assert.Equal(t, false, q.IsEmpty(), "expected nonempty queue")
}

func TestDequeue(t *testing.T) {
	q := prio_queues.Empty()
	q.Enqueue("murmel", 2)
	q.Enqueue("urmel", 3)
	q.Enqueue("bein", 10)
	q.Enqueue("maus", 2)

	assert.EqualValues(t, "bein", q.Dequeue(), "expected bein")
	assert.EqualValues(t, "urmel", q.Dequeue(), "expected bein")
}
