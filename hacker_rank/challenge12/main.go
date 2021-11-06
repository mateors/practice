package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here

	//range check between s and t
	var appleCounter, orangeCounter int

	//apple counter
	for _, num := range apples {
		nn := a + num
		if nn >= s && nn <= t {
			appleCounter++
		}
	}

	//orange counter
	for _, num := range oranges {
		nn := b + num
		if nn >= s && nn <= t {
			orangeCounter++
		}
	}

	fmt.Println(appleCounter)
	fmt.Println(orangeCounter)

}

func main() {

	// var s int32 = 7
	// var t int32 = 10
	// var a int32 = 4
	// var b int32 = 12
	// var apples = []int32{2, 3, -4}
	// var oranges = []int32{3, -2, -4}

	var s int32 = 7
	var t int32 = 11
	var a int32 = 5
	var b int32 = 15
	var apples = []int32{-2, 2, 1} //m=3
	var oranges = []int32{5, -6}   //n=2
	countApplesAndOranges(s, t, a, b, apples, oranges)

}
