package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here

	var max int32
	for i, c := range candles {

		if i == 0 {
			max = c
		}

		if max < c {
			max = c
		}
	}

	var counter int32
	for _, v := range candles {
		if max == v {
			counter++
		}
	}
	return counter
}

func main() {

	//candles := []int32{4, 4, 1, 3}
	//candles := []int32{3, 2, 1, 3}
	candles := []int32{30, 1, 100, 100}
	count := birthdayCakeCandles(candles)
	fmt.Println(count)

}
