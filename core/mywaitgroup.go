package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(id int, g *sync.WaitGroup) {
	fmt.Printf("%d starting \r\n", id)
	time.Sleep(time.Second)
	fmt.Printf("%d ending \r\n", id)
	g.Done()
}

func main() {
	var g sync.WaitGroup
	for i := 0; i < 5; i++ {
		g.Add(1)
		go worker1(i, &g)
	}
	g.Wait()

	fmt.Println(" main exit")
}
