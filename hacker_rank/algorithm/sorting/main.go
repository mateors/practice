package main

import (
	"fmt"
)

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
func merge2(aa, bb []int32, c *[]int32) {

	//var c [6]int32
	//var c = make([]int32, len(a)+len(b))

	aSize := len(aa)
	bSize := len(bb)
	fmt.Println("*", aa, bb, c)

	var i, j, k int = 0, 0, 0

	fmt.Println("#a,b:", aa, bb)

	for i < aSize && j < bSize {

		//fmt.Println(i, j, k)

		if aa[i] < bb[j] {
			fmt.Println(bb[j], "b is higher than a", aa[i])
			(*c)[k] = aa[i]
			i++

		} else {
			//fmt.Println(a[i], "a is higher than b", b[j])
			fmt.Println("aa,bb:", aa, bb)
			fmt.Println(bb[j], "b is less than a", aa[i], aa, bb, i, j)
			(*c)[k] = bb[j] //*****
			fmt.Println("aa,bb:", aa, bb)
			j++
		}
		k++

		fmt.Println("a,b:", aa, bb)
	}

	fmt.Println(">>", i, j, k, "Array:", aa, bb)
	for i < aSize {
		//fmt.Println(a[i])
		(*c)[k] = aa[i]
		fmt.Println("i<aSize:", i, aa[i], c)
		i++
		k++
	}

	for j < bSize {
		(*c)[k] = bb[j]
		j++
		k++
	}
	//fmt.Println("merge:", c)
}

func mergeSort(a []int32) {

	//var b = make([]int32, len(a))
	size := 9 //len(a)
	var last int = 1

	for last < 5 {

		lindx := 0
		for i := 0; i < size; i += last {

			lsi := lindx
			lei := lsi + last
			rsi := lei
			rei := lei + last //start + last + last
			if rei > size {
				continue
			}

			//aa := make([]int32, len(a[lsi:lei]))
			//merge2(a[lsi:lei], a[rsi:rei], &a) //error prone approach because of slice gotchas

			res := merge(a[lsi:lei], a[rsi:rei])
			ri := 0
			for k := lsi; k < rei; k++ {
				a[k] = res[ri]
				ri++
			}
			fmt.Println(a, lsi, "-", lei, ",", rsi, "-", rei, "==>", a[lsi:lei], a[rsi:rei], "=", a)
			lindx = rei
			//merge2(a[lsi:lei], a[rsi:rei], a)
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

	// a := []int32{3, 5, 7, 9}
	// b := []int32{2, 4, 6, 8}
	// result := merge(a, b)
	// fmt.Println(result)

	// a := []int32{9}
	// b := []int32{3}
	// c := []int32{9, 3, 7, 5, 6, 4, 8, 2, 1}

	// fmt.Println(c)
	// merge2(a, b, &c)
	// fmt.Println(c)

	a := []int32{9, 3, 7, 5, 6, 4, 8, 2, 1}
	mergeSort(a)

}
