package myutils

/*import (
	"fmt"
)*/

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

func (q *Queue) NewQueue(size int) {
	q.head = 0
	q.tail = 0
	q.count = 0
	q.nodes = make([]Node, size)
}

const MAX_SIZE = 100

func (q *Queue) Enqueue(url string) {
	if !visited[url] {
		if q.count <= MAX_SIZE {
			var temp Node
			temp.url = url
			visited[url] = true
			q.nodes[q.tail] = temp
			q.tail = q.tail + 1
			q.count = q.count + 1

		}
	}
}

func (q *Queue) Dequeue() string {
	if(q.count == 0) {
		q.head = 0
		q.tail = 0
	}
	if q.count > 0 {
		temp := q.nodes[q.head]
		q.head = q.head + 1
		q.count = q.count - 1
		return temp.url
	}
	return ""
}
