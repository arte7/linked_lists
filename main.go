package main

import (
	"data_structures/stack"
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
}
