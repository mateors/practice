package main

import (
	"fmt"
)

func suffinxArray(str string) []string {

	var result []string

	length := len(str)
	//result = append(result, str)
	for i := 1; i < length; i++ {
		result = append(result, str[i:length])
	}

	//fmt.Println(result)
	//sort.Strings(result)
	return result
}

func longestCommonPrefix(str, compare string) int {

	var count int
	for i := 0; i < len(str); i++ {

		a := str[i]
		c := compare[i]
		//fmt.Println(a, c)
		if a != c {
			return count
		}
		count++
	}
	return count
}

func fizzBuzz(n int32) {
	// Write your code here
	var i int32
	for i = 1; i <= n; i++ {

		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 && i%3 != 0 {
			fmt.Println("Buzz")
		} else if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}

type csort struct {
	name     string
	raminder int
	length   int
}

func RemainderSorting(strArr []string) []string {

	var sortedAray []string

	remainder := []int{}
	lenaray := []int{}

	//csort:=&{}
	rem := make(map[string]int)

	//var min

	for _, name := range strArr {

		length := len(name)
		mod3 := length % 3

		remainder = append(remainder, mod3)

		rem[name] = mod3
		lenaray = append(lenaray, length)
		//fmt.Println(name, length, mod3)
	}

	rrr := make(map[int][]string)
	for key, val := range rem {

		fmt.Println(key, val)

		var a []string
		s, ok := rrr[val]
		if ok {
			rrr[val] = a
		} else {
			a = append(a, s)
			rrr[val] = a
		}

	}

	// fmt.Println(rem)
	// fmt.Println(lenaray)
	// fmt.Println(remainder)
	//sort.Ints(remainder)
	//rrr := make(map[int][]string)

	// for i, val := range remainder {

	// }
	return sortedAray

}

func main() {

	// str := "ababaa"
	// slice := suffinxArray(str)
	// fmt.Println(slice)

	//c := longestCommonPrefix("a", "ababaa")
	//fmt.Println(c)

	//fizzBuzz(15)
	//remainder := 4 % 3
	//fmt.Println(remainder)

	strArr := []string{"Colorado", "Utah", "Wisconsin", "Oregon"}

	sA := RemainderSorting(strArr)
	fmt.Println(sA)
}
