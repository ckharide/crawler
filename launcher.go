package main

import (
	"fmt"
	"github.com/gocrawler/myutils"
)

func main() {
	fmt.Println("Hello World New")
	testQueue()
}

func buildGraph(url string) {

}

func testQueue() {
	var x myutils.Queue
	var g myutils.Graph

	x.NewQueue(100)
	//x.Enqueue("http://yahoo.com")
	x.Enqueue("http://google.com")

	g.BuildGraph(100)

	ch := make(chan string)
	finishedch := make(chan bool)

	url := x.Dequeue()
	fmt.Println("Dequeued Item", url)
	g.AddNode(myutils.Vertex{URL: url})

	go myutils.ScrapeURL(url, ch, finishedch)

	for {
		select {
		case finalurl := <-ch:
			x.Enqueue(finalurl)
			//fmt.Println("Enqueued URL ", finalurl)
			g.AddNode(myutils.Vertex{URL: finalurl})
			g.AddEdge(&myutils.Vertex{URL: url}, &myutils.Vertex{URL: finalurl})
			//g.Print()
			fmt.Println("**************************\n")
			fmt.Println("Size ", g.GetNodes())
			/*case finished := <-finishedch:
			fmt.Println("Finished", finished)*/
		}
	}

}
