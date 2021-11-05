package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var result []int32
	//var a [2]string
	var ac, bc int32

	for i := 0; i < 3; i++ {
		alice := a[i]
		bob := b[i]
		if alice > bob {
			ac++
		} else if alice < bob {
			bc++
		}
	}

	//result[0] = ac
	//result[1] = bc
	result = append(result, ac, bc)
	return result
}

func main() {

	resA := compareTriplets([]int32{1, 2, 3}, []int32{2, 3, 4})
	fmt.Println(resA)

}
