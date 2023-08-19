package main

import (
	"fmt"

	
)

type Student struct {
	Name string
	ID int
	Age float64
}


func addStudent(students []string, student string) []string{
	return append(students, student)
} 

func addStudentID(students []int, student int) []int {
	return append(students, student)
}

func addStudentStruct(students []Student, student Student) []Student {
	return append(students, student)
}


func main(){
	var students = make([]string, 0)
	// students := []string{}
	students = addStudent(students, "Mohhamed Ali")
	students = addStudent(students, "Shakil O`Nill")
	students = addStudent(students, "Lex valensa")
	fmt.Println(students)

	studentsID := []int{}
	studentsID = addStudentID(studentsID, 112)
	studentsID = addStudentID(studentsID, 113)
	studentsID = addStudentID(studentsID, 115)
	fmt.Println(studentsID)

	studentsStruct := []Student{}
	studentsStruct = append(studentsStruct, Student{Name: "Ivan", ID: 101, Age: 21.0})
	studentsStruct = append(studentsStruct, Student{Name: "Petrov", ID: 102, Age: 23.5})
	studentsStruct = append(studentsStruct, Student{Name: "Sidorov", ID: 105, Age: 19.0})
	fmt.Println(studentsStruct)
}