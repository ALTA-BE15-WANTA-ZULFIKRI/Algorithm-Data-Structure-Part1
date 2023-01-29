package main 

import ( 
              "fmt"
               "strconv"
)

func munculsekali(angka string) [] int {
	var result []int

	for i := 0; i < len(angka); i++ {
		count := 0

		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] {
				count++
			}
		}

		if count == 1 {
			num, _ := strconv.Atoi(string(angka[i]))
			result = append(result, num)
		}
	}

   return result
}

func main () {
        fmt.Println(munculsekali("1234123"))           // [4]
        fmt.Println(munculsekali("76523752"))      // [6 , 3]
        fmt.Println(munculsekali("12345"))              // [1 2 3 4 5]
        fmt.Println(munculsekali("1122334455")) // [] 
        fmt.Println(munculsekali("0872504"))       // [8 7 2 5 4 ]
}