package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

func (q *Queue) enqueue(value int) {
	newNode := &Node{value: value}

	if q.rear != nil {
		q.rear.next = newNode
	}
	q.rear = newNode

	if q.front == nil {
		q.front = newNode
	}
	q.size++
}

func (q *Queue) display() {
	temp := q.front
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func (q *Queue) displayPrimes(){
	temp := q.front
	for temp != nil{
		if isPrime(temp.value){
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

func main() {
	q := &Queue{}
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(41)
	q.enqueue(37)
	q.enqueue(4)
	q.display()

	q.displayPrimes()
}
