package main

import (
	"fmt"
	"github.com/gocrawler/myutils"
)

func main() {
	fmt.Println("Hello World New")
	testQueue()
}

func testQueue() {
	var x myutils.Queue
	x.NewQueue(10)
	x.Enqueue("http://google.com")
	x.Enqueue("http://yahoo.com")

	y := x.Dequeue()
	fmt.Println("Dequeued Item", y)
}
