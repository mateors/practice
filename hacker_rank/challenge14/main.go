package main

import "fmt"

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

// for _, num := range b {

// 	var i int32
// 	for i = 1; i <= num; i++ {

// 		if num%i == 0 {
// 			fmt.Println(i, num)
// 		}
// 	}
// }

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here

	var i, ans int32
	for i = 1; i <= 100; i++ {

		var flag bool = true
		for _, num := range a {

			fmt.Println(i, num, flag)
			if i%num != 0 {
				flag = false
				fmt.Println(i, "first break")
				break
			}

		}

		if flag {
			for _, num := range b {

				fmt.Println(">>", num, i, flag)
				if num%i != 0 {
					flag = false
					fmt.Println(i, "second break")
					break
				}

			}
		}

		if flag {
			ans++
			fmt.Println("**", i, ans)
		}
	}
	return ans
}

func main() {

	//start time: 14:24
	//end time: 18:49

	// var a = []int32{2, 6}
	// var b = []int32{24, 36}
	var a = []int32{2, 4}
	var b = []int32{16, 32, 96}
	res := getTotalX(a, b)
	fmt.Println(res)
}
