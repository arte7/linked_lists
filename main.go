package main

import (
	"data_structures/stack"
	"data_structures/d_list"
	"fmt"
)

func main() {
	fmt.Println("gsgaukdgus")
	s := stack.New()
	fmt.Printf("%v", s.Top())
	s.Push("kiwi")
	fmt.Printf("%v", s.Top())
	s.Push("brezel")
	fmt.Printf("%v \n", s.Top())
	fmt.Printf("%v \n", s.Pop())
	fmt.Printf("%v \n", s.Top())
	nums := d_list.New("1", "3", "2", "5")
	fmt.Printf("length: %v \n", nums.Length())
}
