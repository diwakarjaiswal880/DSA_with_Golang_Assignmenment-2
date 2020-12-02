package main

import "fmt"

// Queue represents a queue
type Queue struct {
	items []int
}

// Enqueue adds integer at the back
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes the integer at the front of the queue
// and RETURNs the removed integer
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(30)
	myQueue.Enqueue(90)
	myQueue.Enqueue(40)
	
	fmt.Println(myQueue)
	dequeue_item:=myQueue.Dequeue()
	fmt.Println("Dequeue Item :", dequeue_item)
	
	fmt.Println(myQueue)


}
