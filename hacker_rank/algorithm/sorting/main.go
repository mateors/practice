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

//a,b=input array (unsorted)
//c = output array (sorted)
func merge2(a, b, c []int32) {

	//var c [6]int32
	//var c = make([]int32, len(a)+len(b))
	aSize := len(a)
	bSize := len(b)

	var i, j, k int = 0, 0, 0

	for i < aSize && j < bSize {

		//fmt.Println(i, j, k)
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
}

func mergeSort(a []int32) {

	//var b = make([]int32, len(a))
	size := len(a)

	// for i := 0; i < len(a); i += 2 {

	// 	if i+1 < size {
	// 		fmt.Println(a[i], a[i+1])
	// 		//merge(a, b []int32)
	// 		if a[i] < a[i+1] {
	// 			b[i] = a[i]
	// 			b[i+1] = a[i+1]
	// 		} else {
	// 			b[i] = a[i+1]
	// 			b[i+1] = a[i]
	// 		}

	// 	} else if i+1 == size {
	// 		fmt.Println("*", a[size-1])
	// 		b[size-1] = a[size-1]
	// 	}
	// 	fmt.Println("1=>", b)
	// }

	//2nd try
	//var m, n, p int

	fmt.Println(a[0:1], a[1:2])
	fmt.Println(a[2:3], a[3:4])
	fmt.Println(a[4:5], a[5:6])
	fmt.Println(a[6:7], a[7:8])
	fmt.Println(a[8:9])
	var last int = 1

	for last < 5 {

		lindx := 0
		for i := 0; i < size; i += last {

			// m := i
			// start := i
			// if last > size {
			// 	last = size
			// }
			// end := start + last
			// if end > size {
			// 	end = size
			// }
			// left := a[start:end]
			// right := a[m+last]
			// if end > size {
			// 	continue
			// }
			//fmt.Println(m, start, end)
			//fmt.Println(m, left, right)
			//fmt.Println(a[i:i+last], a[i+last:i+last+last])
			//if (i>1){}

			// fmt.Println(a[0:1], a[1:2])
			// fmt.Println(a[2:3], a[3:4])
			// fmt.Println(a[4:5], a[5:6])
			// fmt.Println(a[6:7], a[7:8])
			// fmt.Println(a[8:9])

			//start := lindx
			// if start+last+last > size {
			// 	continue
			// }
			//fmt.Println(start, "-", i+last, ",", i+last, "-", i+last+last)
			lsi := lindx
			lei := lsi + last
			rsi := lei
			rei := lei + last //start + last + last
			if rei > size {
				continue
			}
			//fmt.Println(start, "-", start+last, ",", start+last, "-", start+last+last, "==>", a[lsi:lei], a[rsi:rei])
			fmt.Println(lsi, "-", lei, ",", rsi, "-", rei, "==>", a[lsi:lei], a[rsi:rei])
			//lindx = start + last + last
			lindx = rei

		}
		last = last * 2
	}

}

func main() {

	// numbers := []int32{5, 6, 7, 8, 9, 1}
	// selectionSort(numbers)
	// fmt.Println(numbers)

	//numbers := []int32{2, 7, 4, 1, 5, 3}
	//bubbleSort(numbers)
	//fmt.Println(numbers)
	// a := []int32{1, 3, 7, 9}
	// b := []int32{2, 5, 6}
	// a := []int32{9}
	// b := []int32{3}
	// result := make([]int32, len(a)+len(b))
	// merge2(a, b, result)
	// fmt.Println(result)

	a := []int32{9, 3, 7, 5, 6, 4, 8, 2, 1}
	mergeSort(a)

}
