package main

import "fmt"

type node struct {
	data string
	next *node
}

type list struct {
	head *node
}

func create(data string) list {
	n := node{data: data, next: nil}
	return list{head: &n}
}

func createNode(data string) node {
	return node{data: data, next: nil}
}

func emptyList() list {
	return list{head: nil}
}

func isEmpty(ls list) bool {
	return ls.head == nil
}

func pushFront(data string, v list) list {
	l := create(data)
	l.head.next = v.head
	return l
}

func first(ls list) *node {
	return ls.head
}

// is -if isEmpty(ls) { return ls } and return emptyList()- the same?
func rest(ls list) list {
	if isEmpty(ls) {
		return ls
	} else {
		if ls.head.next == nil {
			return emptyList()
		} else {
			return list{head: ls.head.next}
		}
	}
}

func draw(ls list) {
	for {
		if isEmpty(ls) {
			fmt.Print("■\n")
			return
		} else {
			fmt.Printf("%s -> ", first(ls).data)

			ls = rest(ls)
		}
	}
}

func len(ls list) int {
	acc := 0

	for {
		if isEmpty(ls) {
			return acc
		} else {
			acc += 1
			ls = rest(ls)
		}
	}
}

func last(ls list) *node {
	// can I remove this when I add it to the loop?
	// if ls.head == nil {
	// 	return nil
	// }

	cur := ls.head

	for {
		if cur == nil || cur.next == nil {
			return cur
		} else {
			cur = cur.next
		}
	}
}

func append(ls1 list, ls2 list) list {
	l := last(ls1)
	l.next = ls2.head
	return ls1
}

// ls = 1 > 2 > 3 > 4
// reverse(rest(ls)) = 4 > 3 > 2
// func reverse(ls list) list {
// 	if isEmpty(ls) {
// 		return emptyList()
// 	} else {
// 		if ls.head.next == nil {
// 			return ls
// 		} else {

// 			// h := first(ls)
// 			// t := rest(ls)
// 			// return pushFront(h.data, v list)
// 		}
// 	}
// }

func main() {
	m := create("baum") // is a list
	o := pushFront("meise", m)
	n := create("maus")
	p := pushFront("haus", n)
	fmt.Printf("len of o: %v \n", len(o))
	fmt.Printf("list:%v \n", o)
	fmt.Printf("first:%v \n", first(o))
	fmt.Printf("rest:%v \n", rest(o))
	fmt.Printf("last: %v \n", last(o))
	draw(o)
	draw(p)
	// fmt.Printf("append: %v \n", append(p, o))
	draw(append(p, o))
}
