package main

import "fmt"




func addStudent(students []string, student string) []string{
	return append(students, student)
} 

func main(){
	var students = make([]string, 0)
	// students := []string{}
	students = addStudent(students, "Mohhamed Ali")
	students = addStudent(students, "Shakil O`Nill")
	students = addStudent(students, "Lex valensa")
	fmt.Println(students)
}