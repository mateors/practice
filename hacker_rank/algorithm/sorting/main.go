package main

import "fmt"

//find the minimum number from an array
func findMin(numbers []int32) int32 {

	var iMin int
	var min int32

	for i := 0; i < len(numbers); i++ {

		if i == 0 {
			iMin = i
			min = numbers[i]
		}

		if min > numbers[i] {
			min = numbers[i]
			iMin = i
		}

	}
	fmt.Println(min, "at index:", iMin)
	return min
}

func selectionSort(numbers []int32) {

	var iMin int
	size := len(numbers)

	for i := 0; i < size-1; i++ {

		iMin = i
		//fmt.Println(i, iMin, numbers[i])
		for j := i + 1; j < size; j++ {

			//fmt.Println(i, j, numbers[j])   //0,1,6
			if numbers[j] < numbers[iMin] { // 6 < 5
				iMin = j
			}
		}
		temp := numbers[i]
		numbers[i] = numbers[iMin]
		numbers[iMin] = temp
	}
}

func bubbleSort(numbers []int32) {

	size := len(numbers)
	for i := 0; i < size; i++ {

		for j := i + 1; j < size-1; j++ {

			fmt.Println(i, j, numbers[i], numbers[j])
			if numbers[i] > numbers[j] { //swap
				// temp := numbers[i]
				// numbers[i] = numbers[j]
				// numbers[j] = temp
				temp := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = temp
			}
		}

		//fmt.Println()
		break

	}

}

func main() {

	// numbers := []int32{5, 6, 7, 8, 9, 1}
	// selectionSort(numbers)
	// fmt.Println(numbers)

	numbers := []int32{2, 7, 4, 1, 5, 3}
	bubbleSort(numbers)
	fmt.Println(numbers)

}
