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

func (s *Stack) pop() int {
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}

func (s *Stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}
func main() {

	stack := &Stack{}
	stack.push(10)
	stack.push(30)
	stack.push(45)
	stack.push(20)

	stack.display()
	stack.pop()
	fmt.Println("After popping")
	stack.display()
}
