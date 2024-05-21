package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) display() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
	// fmt.Println(s.items)
}

func main(){
	stack := &Stack{}

	stack.push(20)
	stack.push(10)
	stack.push(40)
	stack.push(70)

	stack.display()
}
