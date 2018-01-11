package main

import (
	"fmt"
)

func main() {
	slice := []int{}
	for i := 0; i < 11; i++ {
		slice = append(slice, i)
	}

	for _, number := range slice {
		var end string
		if number%2 == 0 {
			end = "even"
		} else {
			end = "odd"
		}
		fmt.Println(number, " is ", end)
	}
}
