package main

import (
	// "fmt"

	// "github.com/turamant/genericdatastruct/fibonachi"
	// "github.com/turamant/genericdatastruct/sumchanel"
	// "github.com/turamant/genericdatastruct/chpt1"
	// "github.com/turamant/genericdatastruct/counter"

	"fmt"

	"github.com/turamant/genericdatastruct/polimorphism"
)


func main(){

	job1 := polimorphism.FixedPriceJob{Description: "Do it ..",
	                                   FixedPrice: 500.40}
	job2 := polimorphism.HourlyJob{Description: "For nine hour",
	                               HourlyRate: 21.30,
								   NumberHour: 11}
	jobs := []polimorphism.JobInterface{&job1,&job2}
	totalcost := polimorphism.TotalJobPrice(jobs)
	fmt.Printf("Total: %0.2f", totalcost)

	// count1 := counter.Counter(&counter.Count{})
    // count1 := counter.Counter(counter.NewCount())
	// fmt.Println("Inicialisation: ",count1.GetCount())
	// for i:=0; i< 10;i++{
	// 	count1.Increment()
	// }
	// count1.Decrement()
	// result := count1.GetCount()
	// fmt.Println("Result: ",result)

	// res := fibonachi.Fibo1(13)
	// fmt.Println(res)
	// s := []int{17,2,8,-19,14,0}
	// c := make(chan int)
	// go sumchanel.Sum(s[:len(s)/2], c) // Суммируем полвину слайса
	// go sumchanel.Sum(s[len(s)/2:], c) // Суммирует вторую половину слайса
	// x, y := <-c, <-c // дожидаемся пока в канале не появятся два значения
	// fmt.Println(x,y,x+y)
	// c2 := make(chan int, 14)
    // go fibonachi.FibonacciChanel(cap(c2), c2)
	// for i := range c2 {
	// 	fmt.Println(i)
	// }
	// students := []string{}
	// students = chpt1.AddStudent[string](students, "Mohhamed Ali")
	// students = chpt1.AddStudent[string](students, "Shakil O`Nill")
	// students = chpt1.AddStudent[string](students, "Lex valensa")
	// fmt.Println(students)

	// studentsID := []int{}
	// studentsID = chpt1.AddStudent[int](studentsID, 113)
	// studentsID = chpt1.AddStudent[int](studentsID, 115)
	// fmt.Println(studentsID)

	// studentsStruct := []chpt1.Student{}
	// studentsStruct = chpt1.AddStudent[chpt1.Student](studentsStruct, chpt1.Student{Name: "Ivan", ID: 101, Age: 21.0})
	// studentsStruct = chpt1.AddStudent[chpt1.Student](studentsStruct, chpt1.Student{Name: "Petrov", ID: 102, Age: 23.5})
	// studentsStruct = chpt1.AddStudent(studentsStruct, chpt1.Student{Name: "Sidorov", ID: 105, Age: 19.0})
	// fmt.Println(studentsStruct)

	
}

	