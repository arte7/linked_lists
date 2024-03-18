package s_list

import "fmt"

type node struct {
	Data interface{}
	next *node
	prio int32
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

func (ls *List) ToArray() []interface{} {
	var r []interface{}

	if ls.IsEmpty() {
		return r
	}
	cur := ls.head

	for cur != nil {
		r = append(r, cur.Data)
		cur = cur.next
	}
	return r
}

func (ls *List) SortedInsert(item interface{}, prio int32) {
	if ls.IsEmpty() {
		ls.head = &node{Data: item, next: nil, prio: prio}
	} else {
		item := &node{Data: item, next: nil, prio: prio}
		cur := ls.head
		var prev *node = nil
		for cur != nil {
			if cur.prio < item.prio {
				item.next = cur
				if prev == nil {
					ls.head = item
				} else {
					prev.next = item
				}
				return
			}
			prev = cur
			cur = cur.next
		}
		prev.next = item
	}
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
	m := New("baum")
	m.PushFront("meise")
	m.PushFront("maus")
	m.PushFront("haus")
	fmt.Printf("len of o: %v \n", m.Length())
	fmt.Printf("list:%v \n", m)
	m.Draw()
	m.Reverse()
	m.Draw()
}
