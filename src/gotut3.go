package main

import "fmt"

// add() form1
// func add(x float64,y float64) float64 {
// 	return x+y
// }
// add() form2
// func add(x,y float64) float64 {
// 	return x+y
// }
// add() form3 change float64 to floate32
func add(x,y float32) float32 {
	return x+y
}

func multiple(a,b string) (string, string) {
	return a,b
}

func main() {
	// declare var form1
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5
	// declare var form2
	// var num1, num2 float64 = 5.6, 9.5
	// declare var form3
	// num1, num2 := 5.6, 9.5

	// Add() func
	// fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1, w2))
}