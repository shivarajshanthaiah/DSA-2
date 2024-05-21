package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) pop() int {
	index := len(s.items) - 1
	element := s.items[index]

	s.items = s.items[:index]
	return element
}
func (s *Stack) display() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
}

func main() {
	stack := &Stack{}

	stack.push(20)
	stack.push(30)
	stack.push(10)
	stack.push(80)

	stack.display()

	stack.pop()
	fmt.Println("After popping")
	stack.display()
}
