package main

import (
	"fmt"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {

	// Write your code here
	var numbers []int64

	for i, _ := range arr {
		//fmt.Println(number)

		var result int64
		for n, num := range arr {

			if i != n {
				result += int64(num)
				//fmt.Println(">>", num, result)
			}
			// if number != num {
			// 	result += int64(num)
			// 	fmt.Println(">>", num, result)
			// }
		}
		//
		numbers = append(numbers, result)
	}

	var max, min int64
	for i, n := range numbers {
		if i == 0 {
			max = n
			min = n
		}
		//fmt.Println(n)
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	fmt.Println(min, max)
	//sort.Ints(numbers)

}

func main() {

	//arr := []int32{1, 2, 3, 4, 5}
	//arr := []int32{1, 3, 5, 7, 9}
	//arr := []int32{9, 3, 5, 7, 19}
	//arr := []int32{1, 2, 5, 11, 9}
	//arr := []int32{19, 2, 5, 11, 9}
	//arr := []int32{5, 5, 5, 5, 5}
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr)

	//numbers := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	//miniMaxSum(numbers)
	// var max, min int32
	// for i, n := range numbers {
	// 	if i == 0 {
	// 		max = n
	// 		min = n
	// 	}
	// 	fmt.Println(n)
	// 	if n > max {
	// 		max = n
	// 	}
	// 	if n < min {
	// 		min = n
	// 	}
	// }
	// fmt.Println(min, max)

	// var one int64
	// one = 623958417 + 467905213 + 714532089 + 938071625
	// fmt.Println(one)
}
