package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	var output strings.Builder

	for _, char := range input {
		output.WriteString(string(int(char) + offset))
	}

	return output.String()
}

func main() { 
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}