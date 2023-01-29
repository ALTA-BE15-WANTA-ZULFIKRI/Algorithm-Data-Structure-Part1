package main 

import (
   "fmt"
) 

func MeanMedian(arrayInput [] float64) (float64, float64)  {
	var mean, median float64
	sum := 0.0

	for _, v := range arrayInput {
		sum += v
	}

	mean = sum / float64(len(arrayInput))

	if len(arrayInput)%2 == 0 {
		median = (arrayInput[len(arrayInput)/2] + arrayInput[len(arrayInput)/2-1]) / 2
	} else {
		median = arrayInput[len(arrayInput)/2]
	}

	return mean, median
}

func main() {
fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))                      // 2.5 2.5 
fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))                 // 3 3
fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))             // 10.4 9 
fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}