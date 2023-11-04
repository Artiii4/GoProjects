package calculations

import "fmt"

func recursiveSearching(n int64) int64 {
	var result int64
	if n == 1 {
		return result
	} else {
		result = n * recursiveSearching(n-1)
	}
	return -1
}
func Calculate(n int64, check bool) int64 {
	//check of n
	if check == true {
		fmt.Println("Start calculations...\nCalculate <", n, ">!")
	}
	result := recursiveSearching(n)
	fmt.Println(result)
	return result
}
