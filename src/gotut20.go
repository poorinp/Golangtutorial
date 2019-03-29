// package main

// import "fmt"

// func foo() {
// 	// using ordinary defer
// 	// defer fmt.Println("Done!")
// 	// defer fmt.Println("Are we done?")
// 	// fmt.Println("Doing some stuff, who knows what?")

// 	//using for loop
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// func main() {
// 	foo()
// }

// improve from EP19
package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
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