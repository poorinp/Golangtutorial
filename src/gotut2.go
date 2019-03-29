package main

// import multiple package
import ("fmt"
		"math/rand")

// Use foo in main()
// func foo() {
// 	fmt.Println("The square root of 4 is", math.Sqrt(4))
// }

func main() {
	// foo()
	// square root function
	// fmt.Println("The square root of 4 is", math.Sqrt(4))

	// Random function
	fmt.Println("A number from 1-100", rand.Intn(100))
}