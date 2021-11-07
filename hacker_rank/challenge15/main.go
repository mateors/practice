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

func main() {

	str := "ababaa"
	slice := suffinxArray(str)
	fmt.Println(slice)
}
