package main

import "fmt"

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	// Write your code here
	var k1, k2, i int32
	var response string = "NO"

	for i = 1; i < 10000; i++ {

		if i == 1 {
			k1 = x1 + v1
			k2 = x2 + v2
		}
		//fmt.Println(i, "->", k1, k2)
		if k1 == k2 {
			return "YES"
		}
		k1 += v1
		k2 += v2
	}
	return response
}

func main() {

	var x1, v1, x2, v2 int32
	// x1 = 0
	// v1 = 3
	// x2 = 4
	// v2 = 2
	// /4523 8092 9419 8076
	x1 = 4523
	v1 = 8092
	x2 = 9419
	v2 = 8076
	res := kangaroo(x1, v1, x2, v2)
	fmt.Println(res)

}
