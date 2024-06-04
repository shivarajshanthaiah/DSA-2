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
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (s *Stack) min() int {
	minVal := s.top.value
	temp := s.top.next
	for temp != nil {
		if temp.value < minVal {
			minVal = temp.value
		}
		temp = temp.next
	}
	return minVal
}

func main() {

	s := &Stack{}
	s.push(30)
	s.push(20)
	s.push(10)
	s.push(34)
	s.push(76)

	s.display()
	fmt.Println(s.min())
}
