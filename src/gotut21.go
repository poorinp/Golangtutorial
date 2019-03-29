package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup: ", r)
	}
	// wg.Done()
}

func say(s string) {
	// defer wg.Done()
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func main() {
	// using goroutine to both but using Sleep
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	// wg.Wait()
	wg.Add(1)
	go say("Hi")
	wg.Wait()
}