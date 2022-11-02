package main

import "fmt"

type Queue struct {
	Que  []int
	head int
	tail int
}

func (q *Queue) queue() {
	q.head = 0
	q.tail = -1
}
func (q *Queue) enQue(data int) {

	q.Que = append(q.Que, data)
	q.tail++
}

func (q *Queue) deQue() {

	i := 0
	for i != q.tail {
		q.Que[i] = q.Que[i+1]
		i++
	}
	q.tail--

}

func (q *Queue) display() {

	fmt.Println(q.Que[:q.tail+1])
}

func main() {

	list := Queue{}
    list.queue()
	list.enQue(1)
    list.enQue(2)
	list.enQue(3)
	list.enQue(4)
	list.enQue(5)
	list.deQue()
	list.display()
}
