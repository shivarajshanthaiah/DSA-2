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

func (s *Stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func (s *Stack) displayPrimes() {
	temp := s.top
	for temp != nil {
		if isPrime(temp.value) {
			fmt.Println(temp.value)
		}
		temp = temp.next
	}
	fmt.Println()
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (s *Stack) middle() int {
	mid := s.size/2
	temp := s.top

	for i :=0;i<mid;i++{
		temp = temp.next
	}
	return temp.value
}

func main() {

	stack := &Stack{}
	stack.push(20)
	stack.push(31)
	stack.push(7)
	stack.push(33)
	stack.push(15)
	stack.push(41)

	stack.display()
	stack.displayPrimes()

	fmt.Println(stack.middle())

}
