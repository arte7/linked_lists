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

func pushFront(data string, v *list) {
	node := createNode(data)
	node.next = v.head
	v.head = &node
}

func first(ls list) *node {
	return ls.head
}

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
	cur := ls.head

	for {
		if cur.next == nil {
			break
		} else {
			cur = cur.next
		}
	}
	return cur
}

// append doesn't work if you  append the same list to the same list
//
//	infiniteloop, is there any workaround switching out a new end? is this even a usecase?
func append(ls1 list, ls2 list) list {
	l := last(ls1)
	l.next = ls2.head
	return ls1
}

func reverse(ls list) list {
	rev := emptyList()
	cur := ls.head

	for {
		if cur == nil {
			return rev
		} else {
			pushFront(cur.data, &rev)
			cur = cur.next
		}
	}
}

func main() {
	m := create("baum") // is a list
	pushFront("meise", &m)
	pushFront("maus", &m)
	pushFront("haus", &m)
	fmt.Printf("len of o: %v \n", len(m))
	fmt.Printf("list:%v \n", m)
	fmt.Printf("rest:%v \n", rest(m))
	draw(m)
	draw(reverse(m))
}
