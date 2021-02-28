package main

import (
	"fmt"
	"time"
)

func worker(c chan bool) {
	fmt.Println("start >> ")
	time.Sleep(time.Second)
	fmt.Println("end >>")

	c <- true
}

func main1() {
	c := make(chan bool)
	go worker(c)
	result := <-c
	fmt.Println(result)
}
