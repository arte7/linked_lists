package d_list

import "fmt"

type d_node struct {
	data interface{} // interface is some reference type
	next *d_node
	prev *d_node
}

type List struct {
	head *d_node
	last *d_node
}

func Empty() *List {
	return &List{head: nil, last: nil}
}

func New(data ...interface{}) *List {
	e := Empty()
	for _, d := range data {
		e.PushBack(d) // methodenaufruf auf pointer auf structuren oder auf strukturen
	}
	return e
}

func (ls *List) Front() interface{} {
	if ls.IsEmpty() {
		return nil
	}

	return ls.head.data
}

func (ls *List) IsEmpty() bool {
	return ls.head == nil && ls.last == nil
}

func (ls *List) PopFront() interface{} {
	if ls.IsEmpty() {
		return nil
	} else if ls.head.next == nil {
		d := ls.head
		ls.head = nil
		ls.last = ls.head
		return d.data
	} else {
		d := ls.head
		ls.head = ls.head.next
		return d.data
	}
}

func (ls *List) PushBack(d interface{}) {
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

func (ls *List) Draw() {
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

func (ls *List) PushDFront(d interface{}) {
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

func (ls1 *List) appendList(ls2 *List) *List {
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

// func (ls *List) Length() int {
// 	acc := 0
// 	cur := ls.head

// 	for {
// 		if cur == nil {
// 			return acc
// 		} else {
// 			acc += 1
// 			cur = cur.next
// 		}
// 	}
// }

func (ls *List) Length() int {
	if ls.IsEmpty() {
		return 0
	}
	
	acc := 2
	p1 := ls.head
	p2 := ls.last

	for {
		if p1 == p2 {
			acc -= 1
			return acc
		} else if p1.next == p2 {
			return acc
		} else {
			acc += 2
			p1 = p1.next
			p2 = p2.prev
		}
	}
}

func (ls *List) reverseDList() *List {
	if ls.IsEmpty() {
		return ls
	}
	cur := ls.last

	ls.last = ls.head
	ls.head = cur

	for {
		if cur == nil {
			return ls
		} else {
			prev := cur.prev

			cur.prev = cur.next
			cur.next = prev

			cur = cur.next
		}
	}
}

func main() {
	list := New("baum", "bert")
	ls := Empty()
	ls.PushBack("haus")
	ls.PushBack("maus")
	ls.PushDFront("keks")
	// fmt.Printf("ls: %v \n", ls)
	// fmt.Printf("head.next.data: %v \n", ls.head.next.data)
	// fmt.Printf("head.next.next: %v \n", ls.head.next.next)
	// fmt.Printf("last.data: %v \n", ls.last.data)

	ls.Draw()
	list.Draw()
	newList := list.appendList(ls)
	fmt.Printf("length: %v \n", newList.Length())
	newList.Draw()
	newList.reverseDList().Draw()
}
