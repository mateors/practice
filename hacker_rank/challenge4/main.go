package main

import (
	"fmt"
	"math"
)

// a square matrix
//what is square matrix? => https://www.mathsisfun.com/definitions/square-matrix.html
//absolute difference
//the sums of its diagonals.
func diagonalDifference(arr [][]int32) int32 {

	// Write your code here
	for i, ar := range arr {
		fmt.Println(i, ar)
	}

	//variable declaration
	var l2r, r2l int32
	var counter int = 0

	//left to right diagonal
	for i := 0; i < len(arr); i++ {
		l2r += arr[i][counter]
		counter++
	}

	//right to left diagonal
	counter = len(arr) - 1
	for i := 0; i < len(arr); i++ {
		r2l += arr[i][counter]
		counter--
	}
	//fmt.Println(l2r)
	//fmt.Println(r2l)
	diagonalDifference := l2r - r2l
	result := math.Abs(float64(diagonalDifference))
	return int32(result)
}

//Abs() is a function from Go language.
//This Abs() function returns the absolute (positive) value
//of any given number (positive or negative)

func main() {

	var arr [][]int32
	//arr = append(arr, []int32{1, 2, 3}, []int32{4, 5, 6}, []int32{9, 8, 9})
	//arr = append(arr, []int32{11, 2, 4}, []int32{4, 5, 6}, []int32{10, 8, -12})
	//arr = append(arr, []int32{11, 2}, []int32{4, 5})
	arr = append(arr, []int32{11, 2, 3, 4}, []int32{4, 5, 5, 7}, []int32{11, 2, 3, 4}, []int32{4, 5, 5, 7})

	result := diagonalDifference(arr)
	fmt.Println(result)

}
