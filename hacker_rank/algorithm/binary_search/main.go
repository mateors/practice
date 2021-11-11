package main

import "fmt"

//Book :: A Common Sense Guide to Data Structures & Algorthims
func binarySearch(array []int32, searchValue int32) int {

	lowerBound := 0
	upperBound := len(array) - 1

	for lowerBound <= upperBound {

		midpoint := (upperBound + lowerBound) / 2
		midpointValue := array[midpoint]

		if searchValue == midpointValue {
			return midpoint

		} else if searchValue < midpointValue {
			upperBound = midpoint - 1

		} else if searchValue > midpointValue {
			lowerBound = midpoint + 1
		}
	}
	return -1
}

func findNumberOfSteps(arrayCount int32) int {

	var count int
	for {

		steps := arrayCount / 2
		arrayCount = steps
		//fmt.Println(">", steps)
		if 1 > steps {
			break
		}
		count++
	}
	return count
}

func main() {

	//as it is a binary search algorithm implementation
	// so array must be sorted or ordered list
	//array := []int32{3, 17, 75, 80, 202}
	//indxPosition := binarySearch(array, 75)
	//fmt.Println(indxPosition)

	steps := findNumberOfSteps(100000)
	fmt.Println("steps:", steps)

}
