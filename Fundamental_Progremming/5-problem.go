package main

import "fmt"

func RemoveDuplicates(array []int) int {
	// your code here
	result := 0
	for i := 0; i < len(array); i++ {
		if i == 0 || array[i] != array[i-1] {
			array[result] = array[i]
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) // 4
	fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9})) // 6
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         // 2
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         // 2
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11}))     // 4

}
