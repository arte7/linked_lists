package s_list

import "fmt"

type node struct {
	Data interface{}
	next *node
}

type List struct {
	head *node
}

func Empty() *List {
	return &List{head: nil}
}

func New(data ...interface{}) *List {
	e := Empty()
	for _, d := range data {
		e.PushFront(d) // methodenaufruf auf pointer auf structuren oder auf strukturen
	}
	e.Reverse()
	return e
}

func CreateNode(data interface{}) *node {
	return &node{Data: data, next: nil}
}

func (ls *List) IsEmpty() bool {
	return ls.head == nil
}

func (ls *List) PushFront(data interface{}) {
	node := CreateNode(data)
	node.next = ls.head
	ls.head = node
}

func (ls *List) First() *node {
	return ls.head
}

func (ls *List) Draw() {
	if ls.IsEmpty() {
		fmt.Print("■")
	}
	cur := ls.head

	for cur != nil {
		fmt.Printf("%v", cur.Data)
		if cur.next != nil {
			fmt.Print(" -> ")
		}
		cur = cur.next
	}
	fmt.Print("\n")
}

func (ls *List) Length() int {
	acc := 0
	cur := ls.head

	for {
		if cur == nil {
			return acc
		} else {
			acc += 1
			cur = cur.next
		}
	}
}

func (ls *List) PopFront() interface{} {
	h := ls.head
	ls.head = ls.head.next
	return h.Data
}

func (ls *List) Reverse() {
	cur := ls.head
	var prev *node = nil
	var next *node = nil

	for {
		if cur == nil {
			break
		} else {
			next = cur.next
			cur.next = prev
			prev = cur
			cur = next
		}
	}
	ls.head = prev
}

func main() {
	m := New("baum") // is a list
	m.PushFront("meise")
	m.PushFront("maus")
	m.PushFront("haus")
	fmt.Printf("len of o: %v \n", m.Length())
	fmt.Printf("list:%v \n", m)
	m.Draw()
	m.Reverse()
	m.Draw()
}
