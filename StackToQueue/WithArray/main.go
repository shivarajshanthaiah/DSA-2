package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	topIndex := len(s.items) - 1
	item := s.items[topIndex]
	s.items = s.items[:topIndex]
	return item
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

type Queue struct {
	stack1 Stack
	stack2 Stack
}

func (q *Queue) enqueue(item int) {
	q.stack1.push(item)
}

func (q *Queue) dequeue() int {
	if q.stack2.isEmpty() {
		for !q.stack1.isEmpty() {
			q.stack2.push(q.stack1.pop())
		}
	}
	return q.stack2.pop()
}

func (q *Queue) display() {

	for i := len(q.stack2.items) - 1; i >= 0; i-- {
		fmt.Print(q.stack2.items[i], " ")
	}
	for i := 0; i < len(q.stack1.items); i++ {
		fmt.Print(q.stack1.items[i], " ")
	}
	fmt.Println()
}

func main() {
	queue := &Queue{}

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.display()

	queue.dequeue()
	queue.display()

}
