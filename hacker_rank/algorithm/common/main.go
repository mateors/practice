package main

import "fmt"

func main() {

	//string
	//int
	//float
	//bool
	//char
	//datetime

	//struct
	//array
	//slice

	//array
	//index -> 0,1,2,3,4
	//value assign 0-3
	//index value
	//data type
	//0,1,2,3
	// var students [4]string
	// students[0] = "Rezaul"
	// students[2] = "Fatema" //
	// students[1] = "Mazid"  //
	// students[3] = "Moaz"   //

	// var students []string //variable declaration
	//fmt.Println(students[1:3])
	var students []string
	//fmt.Printf("%T", students)
	//fmt.Println(reflect.TypeOf(students).Kind())
	students = append(students, "A")
	students = append(students, "B")
	students = append(students, "C")
	students = append(students, "D") //4
	students = append(students, "E") //previusCapacity*2=>4*2=8
	students = append(students, "F")
	students = append(students, "G")
	students = append(students, "H")
	students = append(students, "E") //previusCapacity*2*8=16
	students = append(students, "E")
	students = append(students, "F") //16*2=32
	//fmt.Println(students[0:])
	//fmt.Println(students[2:])
	//fmt.Println(students[:4])
	//fmt.Println(students[:])
	//50/10
	//fmt.Println(len(students[2:]))
	fmt.Println(cap(students))

}
