package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var results []int32
	var passGrade int32 = 40

	for _, mark := range grades {

		//fmt.Println(i, mark)
		for j := mark; j < mark+5; j++ {

			if j%5 == 0 {

				difference := j - mark
				if j >= passGrade && difference < 3 {
					results = append(results, j)
					//fmt.Println(j, "is devisible by 5, diff:", difference)
				} else {
					results = append(results, mark)
				}
			}
		}
	}
	return results
}

func main() {

	grades := []int32{73, 67, 38, 33}
	finalGrades := gradingStudents(grades)
	fmt.Println(grades, "-->", finalGrades)
}
