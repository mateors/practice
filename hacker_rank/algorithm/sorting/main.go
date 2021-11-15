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

func mergeSort() {

}

func merge(a, b []int32) []int32 {

	//var c [6]int32
	var c = make([]int32, len(a)+len(b))
	aSize := len(a)
	bSize := len(b)

	var i, j, k int = 0, 0, 0

	for i < aSize && j < bSize {

		fmt.Println(i, j, k)
		if a[i] < b[j] {
			c[k] = a[i]
			//fmt.Println(a[i], "b is higher than a", b[j])
			i++

		} else {
			//fmt.Println(b[j], "a is higher than b", a[i])
			c[k] = b[j]
			j++
		}
		k++
	}

	//fmt.Println(">>", i, j, k)
	for i < aSize {
		c[k] = a[i]
		i++
		k++
	}

	for j < bSize {
		c[k] = b[j]
		j++
		k++
	}
	return c
}

func main() {

	// numbers := []int32{5, 6, 7, 8, 9, 1}
	// selectionSort(numbers)
	// fmt.Println(numbers)

	//numbers := []int32{2, 7, 4, 1, 5, 3}
	//bubbleSort(numbers)
	//fmt.Println(numbers)
	a := []int32{1, 3, 7, 9}
	b := []int32{2, 5, 6}
	c := merge(a, b)
	fmt.Println(c)

}
