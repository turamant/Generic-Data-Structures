package sumchanel

import (
	
)

func Sum(s []int, c chan int){
	//Суммирует элементы слайса и
    //кладет результат в канал c.
	sum := 0
	for _, v := range s {
		sum += v
	}
	c<-sum
}
