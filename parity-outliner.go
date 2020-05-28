package main

import (
	"fmt"
	"math"
)

func findoutliner(arrayInt []int) int {
	numerOfOdds := 0
	numberOfEvens := 0
	var integerOutlier float64

	for i := 0; i < len(arrayInt); i++ {
		if arrayInt[i]%2 == 0 {
			numberOfEvens++
		} else {
			numerOfOdds++
		}
	}

	if numberOfEvens > numerOfOdds {
		integerOutlier = 1
	} else {
		integerOutlier = 0
	}

	for i := 0; i < len(arrayInt); i++ {
		if math.Abs(float64(arrayInt[i]%2)) == integerOutlier {
			return arrayInt[i]
		}
	}

	return 0
}

func main() {
	array1 := []int{2, 4, 0, 100, 4, 11, 2602, 36}

	fmt.Println(findoutliner(array1))
}
