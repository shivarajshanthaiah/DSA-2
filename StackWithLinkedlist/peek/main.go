package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) push(value int) {
	newNode := &Node{
		value: value,
		next:  s.top,
	}
	s.top = newNode
	s.size++
}

func (s *Stack) peek() int {
	return s.top.value
}

func main() {
	stack := &Stack{}
	stack.push(20)
	stack.push(30)
	stack.push(10)
	stack.push(50)

	fmt.Println(stack.peek())
}
