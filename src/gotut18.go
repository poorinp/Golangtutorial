package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	// say("Hey")
	// say("There")

	// using goroutine to Hey
	// go say("Hey")
	// say("There")

	// using goroutine to both
	// go say("Hey")
	// go say("There")

	// using goroutine to both but using Sleep
	go say("Hey")
	go say("There")
	time.Sleep(time.Second)
}