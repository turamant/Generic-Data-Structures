package polimorphism

import (

)

type FixedPriceJob struct {
	Description string
	FixedPrice float64
}

type HourlyJob struct {
	Description string
	HourlyRate float64
	NumberHour int
}

type JobInterface interface{
	Cost() float64
    GetDescription() string
}

func (fp *FixedPriceJob) Cost() float64 {
	return fp.FixedPrice
}

func (fp FixedPriceJob) GetDescription() string{
	return fp.Description
}

func (hj *HourlyJob) Cost() float64 {
	return hj.HourlyRate * float64(hj.NumberHour)
}

func (hj HourlyJob) GetDescription() string{
	return hj.Description
}

func TotalJobPrice(jobs []JobInterface) float64{
	result := 0.0
	for _, job := range jobs{
		result += job.Cost()
	}
	return result
}