package s_list_test

import (
	"data_structures/s_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedInsert(t *testing.T) {
	l := s_list.Empty()
	l.SortedInsert("baum", 8)
	l.SortedInsert("maus", 10)
	l.SortedInsert("laus", 8)
	l.SortedInsert("haus", 6)
	l.SortedInsert("klaus", 1)
	r := l.ToArray()

	assert.EqualValues(t, 5, len(r), "they should be equal")
	assert.EqualValues(t, "maus", r[0], "they should be equal")
	assert.EqualValues(t, "baum", r[1], "they should be equal")
	assert.EqualValues(t, "laus", r[2], "they should be equal")
	assert.EqualValues(t, "haus", r[3], "they should be equal")
	assert.EqualValues(t, "klaus", r[4], "they should be equal")

}
