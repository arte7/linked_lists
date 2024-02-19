package main

import "fmt"

type d_node struct {
	data interface{} // interface is some reference type
	next *d_node
	prev *d_node
}

type d_list struct {
	head *d_node
	last *d_node
}

func Empty() *d_list {
	return &d_list{head: nil, last: nil}
}

func New(data ...interface{}) *d_list {
	e := Empty()
	for _, d := range data {
		e.PushBack(d) // methodenaufruf auf pointer auf structuren oder auf strukturen
	}
	return e
}

func (ls *d_list) IsEmpty() bool {
	return ls.head == nil && ls.last == nil
}

func (ls *d_list) PushBack(d interface{}) {
	node := &d_node{data: d, next: nil, prev: nil}

	if ls.IsEmpty() {
		ls.head = node
		ls.last = node
	} else {
		ls.last.next = node
		node.prev = ls.last
		ls.last = node
	}
}

func (ls *d_list) Draw() {
	if ls.IsEmpty() {
		fmt.Print("■")
	}
	cur := ls.head

	for cur != nil {
		fmt.Printf("%v", cur.data)
		if cur.next != nil {
			fmt.Print(" -> ")
		}
		cur = cur.next
	}
	fmt.Print("\n")
}

func (ls *d_list) PushFront(d interface{}) {
	node := &d_node{data: d, next: nil, prev: nil}

	if ls.IsEmpty() {
		ls.head = node
		ls.last = node
	} else {
		ls.head.prev = node
		node.next = ls.head
		ls.head = node
	}
}

func (ls1 *d_list) appendList(ls2 *d_list) *d_list {
	if ls2.IsEmpty() {
		ls2.head = ls1.head
		ls2.last = ls1.last
		return ls2
	} else if ls1.IsEmpty() {
		ls1.head = ls2.head
		ls1.last = ls2.last
		return ls1
	} else {
		ls2.head.prev = ls1.last
		ls1.last.next = ls2.head
		ls1.last = ls2.last
	}
	return ls1
}

// -push front-, -appendList-, dLength, reverse(inplace), queues & stacks (time complexity)

func main() {
	list := New("baum", "bert")
	ls := Empty()
	ls.PushBack("haus")
	ls.PushBack("maus")
	ls.PushFront("keks")
	fmt.Printf("ls: %v \n", ls)
	fmt.Printf("head.next.data: %v \n", ls.head.next.data)
	fmt.Printf("head.next.next: %v \n", ls.head.next.next)
	fmt.Printf("last.data: %v \n", ls.last.data)

	ls.Draw()
	list.Draw()
	newList := list.appendList(ls)
	newList.Draw()
}
