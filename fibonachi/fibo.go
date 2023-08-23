package fibonachi

import (

)

func Fibo1(n int) int{
	x, y := 0, 1
	for i:=0; i < n; i++{
		x, y = y, x + y
	}
	return x
}

func Fibo2(n int) int{
	if n < 2{
		return n
	}
	return Fibo2(n-1) + Fibo2(n-2)
}

func FibonacciChanel(n int, c chan int) {
	x, y := 0, 1
	for i:=0; i<n;i++{
		c <- x
		x, y = y, x+y
	}
	close(c)
}

