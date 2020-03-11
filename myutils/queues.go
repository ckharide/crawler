package myutils

import (
	"fmt"
)

type Node struct {
	url string
}
type Queue struct {
	head  int
	tail  int
	nodes []Node
	count int
}

var visited = make(map[string]bool)
var MAX_SIZE int

func (q *Queue) NewQueue(size int) {
	q.head = -1
	q.tail = -1
	q.count = 0
	q.nodes = make([]Node, size)
	MAX_SIZE = size
	fmt.Println("leng of the queue  ", len(q.nodes))
}

func (q *Queue) Enqueue(url string) bool {
	if q.count == MAX_SIZE {
		return false
	}

	if !visited[url] {
		if q.count < MAX_SIZE {
			var temp Node
			temp.url = url
			visited[url] = true
			q.head = (q.head + 1) % MAX_SIZE
			q.nodes[q.head] = temp
			q.count = q.count + 1

			if q.tail == -1 {
				q.tail = q.head
			}
			return true
		}
	}
	return false
}

func (q *Queue) Dequeue() string {
	if q.count == 0 {
		return ""
	} else {
		temp := q.nodes[q.tail]
		q.count = q.count - 1
		q.tail = (q.tail + 1) % MAX_SIZE
		if q.count == 0 {
			q.head = -1
			q.tail = -1
		}
		return temp.url
	}
	return ""
}

func (q *Queue) Size() int {
	return q.count
}
