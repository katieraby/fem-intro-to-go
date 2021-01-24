package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

//func say(s string) {
//	for i := 0; i < 3; i++ {
//		fmt.Println(s)
//		time.Sleep(time.Millisecond * 300)
//	}
//}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC")
	}
}
func printStuff() {
	defer waitGroup.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	waitGroup.Add(1) // add to a queue
	go printStuff()
	waitGroup.Wait()
//	go say("Hello")
//	say("There")
}
