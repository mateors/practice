package main

import "fmt"

func plusMinus(arr []int32) {

	// Write your code here
	var positiveNumbers, negativeNumbers, zerosNumbers int32

	for _, number := range arr {

		if number > 0 {
			positiveNumbers++
		} else if number < 0 {
			negativeNumbers++
		} else {
			zerosNumbers++
		}
	}

	postiveRatio := float32(positiveNumbers) / float32(len(arr))
	negativeRatio := float32(negativeNumbers) / float32(len(arr))
	zerosRatio := float32(zerosNumbers) / float32(len(arr))

	fmt.Printf("%.6f\n", postiveRatio)
	fmt.Printf("%.6f\n", negativeRatio)
	fmt.Printf("%.6f\n", zerosRatio)

}

func main() {

	plusMinus([]int32{1, 1, 0, -1, -1})

}
