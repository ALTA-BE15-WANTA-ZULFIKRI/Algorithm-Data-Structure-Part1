package main 

import (
            "fmt"
)

func DrawXYZ(N int) string {
	var result string

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j % 2 == 0 {
				result += "Y Z X "
			} else {
				result += "Z Y X "
			}
		}
		result += "\n"
	}

	return result
}

func main () {
fmt.Println(DrawXYZ(3))


fmt.Println(DrawXYZ(5))

}