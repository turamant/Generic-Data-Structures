package counter


import (

)

type Counter interface{
	Increment()
	Decrement()
	Reset()
	GetCount() int
}

type Count struct {
	count int
}

func NewCount() *Count{
	return &Count{
		count: 0,
	}
}

func (c *Count) Increment(){
	c.count++
}

func (c *Count) Decrement(){
	c.count--
}

func (c *Count) Reset(){
	c.count = 0
}

func (c *Count) GetCount() int {
	return c.count
}