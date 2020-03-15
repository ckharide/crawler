package main

import (
	"fmt"
	"github.com/gocrawler/myutils"
)

func main() {
	_Init()
}

func _Init() {
	var queue myutils.Queue
	//var g myutils.Graph kept for future reference if we want to use graph DS.

	queue.NewQueue(5)
	queue.Enqueue("http://www.google.com")
	ch := make(chan string)
	finishedch := make(chan bool)

	url := queue.Dequeue()

	//g.AddNode(myutils.Vertex{URL: url})

	go myutils.ScrapeURL(url, ch, finishedch)

	//defer print(&queue)
	for {
		select {
		case finalurl := <-ch:
			if queue.Enqueue(finalurl) {
				fmt.Println("Enqueued URL ", finalurl)
				go myutils.ScrapeURL(finalurl, ch, finishedch)
			}

		case finished := <-finishedch:
			fmt.Println("Finished", finished)
			queue.Echo()

		}
	}

}
