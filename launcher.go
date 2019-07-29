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
	x.Enqueue("http://yahoo.com")
	x.Enqueue("http://google.com")

	ch := make(chan string)
	finishedch := make(chan bool)

	url := x.Dequeue()
	fmt.Println("Dequeued Item", url)

	go myutils.ScrapeURL(url, ch, finishedch)

	for {
		select {
		case url := <-ch:
			//	x.Enqueue(url)
			fmt.Println("Enqueued URL ", url)
		case finished := <-finishedch:
			fmt.Println("Finished", finished)

		}
	}

}
