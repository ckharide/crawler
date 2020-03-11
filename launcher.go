package main

import (
	"fmt"
	"github.com/gocrawler/myutils"
)

func main() {
	init_main()
}

func init_main() {
	var x myutils.Queue
	//var g myutils.Graph kept for future reference if we want to use graph DS.

	x.NewQueue(1)
	x.Enqueue("http://www.google.com")
	ch := make(chan string)
	finishedch := make(chan bool)

	url := x.Dequeue()

	//g.AddNode(myutils.Vertex{URL: url})

	go myutils.ScrapeURL(url, ch, finishedch)

	for {
		select {
		case finalurl := <-ch:
			if(x.Enqueue(finalurl)) {
				fmt.Println("Enqueued URL ", finalurl)
				go myutils.ScrapeURL(finalurl, ch, finishedch)
			}

		case finished := <-finishedch:
			fmt.Println("Finished", finished)
		}
	}

}
