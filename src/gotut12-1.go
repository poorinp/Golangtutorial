package main 

import "fmt"

// simple loop
// func main() {
// 	for i:=0 ; i<10 ; i++ {
// 		fmt.Println(i)
// 	}
// } 

// another way to use loop
// func main() {
// 	i := 0
// 	for i < 10 {
// 		fmt.Println(i)
// 		i+=1
// 	}
// } 

// using +=5 instead of +=1
// func main() {
// 	i := 0
// 	for i < 10 {
// 		fmt.Println(i)
// 		i+=5
// 	}
// } 

// infinity loop
// func main() {
// 	for {
// 		fmt.Println("Do Stuff")
// 	}
// } 

// increase round by 3 to under 25
// func main() {
// 	x := 5 
// 	for {
// 		fmt.Println("Do stuff", x)
// 		x+=3
// 		if x > 25 {
// 			break
// 		}
// 	}
// }

// same as above but different pointer
func main() {
	a := 3
	for x := 5 ; a < 25 ; x +=3 {
		fmt.Println("Do stuff", x)
		a += 4 
	}
}