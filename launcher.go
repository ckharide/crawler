package main

import (
	"fmt"
	"github.com/gocrawler/myutils"
)

func main() {
	init_main()
}

func buildGraph(url string) {

}

func init_main() {
	var x myutils.Queue
	//var g myutils.Graph

	x.NewQueue(10)
	x.Enqueue("http://www.google.com")

	ch := make(chan string)
	finishedch := make(chan bool)

	url := x.Dequeue()
	//fmt.Println("Dequeued Item", url)
	//g.AddNode(myutils.Vertex{URL: url})

	go myutils.ScrapeURL(url, ch, finishedch)

	for {
		select {
		case finalurl := <-ch:
			x.Enqueue(finalurl)
			fmt.Println("Enqueued URL ", finalurl)
			go myutils.ScrapeURL(finalurl, ch, finishedch)

		case finished := <-finishedch:
			fmt.Println("Finished", finished)
		}
	}

}
