package pkg

import (
	"fmt"
)

func recursiveSearching(n int64) int64 {
	var result int64 = 1
	if n == 1 {
		return result
	} else {
		result = n * recursiveSearching(n-1)
	}
	return result
}
func Calculate(n int64, check bool) int64 {
	if n < 1 {
		fmt.Println("ERROR! Factorial is considered for positive number")
		return 0
	}
	if check == true {
		fmt.Println("Start calculations...\nCalculate <", n, ">!")
	}
	result := recursiveSearching(n)
	if check == true {
		fmt.Println("Calculations complete!")
	}
	return result
}
