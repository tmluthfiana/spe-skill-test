package main

import (
	"fmt"
)

func findDiff(array1, array2 []int) (result []int) {

	for _, val1 := range array1 {
		for _, val2 := range array2 {
			if val1 != val2 {
				result = append(result, val1)
			}
		}
	}

	return
}

func main() {
	array1 := []int{1, 5, 5, 5, 5, 3}
	array2 := []int{5}
	result := findDiff(array1, array2)
	fmt.Println(result)
}
