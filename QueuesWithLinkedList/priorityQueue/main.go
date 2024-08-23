package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type PriorityQueue struct {
	front *Node
}

func (pq *PriorityQueue) insert(val int) {
	newNode := &Node{Val: val}

	if pq.front == nil || val > pq.front.Val {
		newNode.Next = pq.front
		pq.front = newNode
	} else {
		temp := pq.front
		for temp.Next != nil && temp.Next.Val > val {
			temp = temp.Next
		}
		newNode.Next = temp.Next
		temp.Next = newNode
	}
}

func (pq *PriorityQueue) display() {
	temp := pq.front
	for temp != nil {
		fmt.Printf("%d -> ", temp.Val)
		temp = temp.Next
	}
	fmt.Println("nil")
}

func main() {
	pq := &PriorityQueue{}

	pq.insert(5)
	pq.insert(1)
	pq.insert(3)
	pq.insert(7)
	pq.insert(6)

	pq.display()
}
