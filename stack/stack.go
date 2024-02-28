package stack

import (
	"data_structures/s_list"
)

type Stack struct {
	ls *s_list.List
}

func New() *Stack {
	return &Stack{ls: s_list.Empty()}
}

func (st *Stack) Top() interface{} {
	if st.ls.IsEmpty() {
		return nil
	}
	return st.ls.First().Data
}

func (st *Stack) Push(data interface{}) {
	st.ls.PushFront(data)
}

func (st *Stack) Pop() interface{} {
	return st.ls.PopFront()
}
