package main

import (
	"fmt"

	
)




func addStudent(students []string, student string) []string{
	return append(students, student)
} 

func addStudentID(students []int, studentID int) []int {
	return append(students, studentID)
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
}