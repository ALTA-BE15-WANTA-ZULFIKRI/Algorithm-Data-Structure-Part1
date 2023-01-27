package main

import "fmt"

func ArrayUnique(arrayA, arrayB []int) []int {
	var arrayC []int

	for _, v := range arrayA {
		arrayC = append(arrayC, v)
	}

	for _, v := range arrayB {
		isExist := false

		for _, v2 := range arrayA {
			if v == v2 {
				isExist = true
				break
			}
		}

		if !isExist {
			arrayC = append(arrayC, v)
		}
	}

	return arrayC
}

func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59}))
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{2, 8}))
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))
}
