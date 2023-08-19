package main

import (
	"fmt"

	
)
// type Stringer interface {
// 	String() string
// }

// type String string 
// func (s String) String() string {
// 	return string(s)
// }

// type Integer int

// func (i Integer) String() string {
// 	return fmt.Sprintf("%d", i)
// }

type Student struct {
	Name string
	ID int
	Age float64
}

// func (s Student) String() string {
// 	return fmt.Sprintf("Name:%s, ID:%d, Age:%0.1f", s.Name, s.ID, s.Age)
// }


func addStudent[T any](students []T, student T) []T{
	return append(students, student)
} 




func main(){
	students := []string{}
	students = addStudent[string](students, "Mohhamed Ali")
	students = addStudent[string](students, "Shakil O`Nill")
	students = addStudent[string](students, "Lex valensa")
	fmt.Println(students)

	studentsID := []int{}
	studentsID = addStudent[int](studentsID, 112)
	studentsID = addStudent[int](studentsID, 113)
	studentsID = addStudent[int](studentsID, 115)
	fmt.Println(studentsID)

	studentsStruct := []Student{}
	studentsStruct = addStudent[Student](studentsStruct, Student{Name: "Ivan", ID: 101, Age: 21.0})
	studentsStruct = addStudent[Student](studentsStruct, Student{Name: "Petrov", ID: 102, Age: 23.5})
	studentsStruct = addStudent(studentsStruct, Student{Name: "Sidorov", ID: 105, Age: 19.0})
	fmt.Println(studentsStruct)
}