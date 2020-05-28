package main

import (
	"fmt"
	"math"
)

func countDigit(number int) float64 {
	if number == 0 {
		return 0
	}

	return 1 + countDigit(number/10)
}

func check(number int) bool {
	length := countDigit(number)
	dup := number
	sum := 0

	for {
		sum += math.Pow(float64(dup%10), length)
		dup /= 10

		if dup == 0 {
			break
		}
	}

	return (number == sum)
}

func main() {
	var number int
	fmt.Print("Enter any number : ")
	fmt.Scan(&number)

	result := check(number)
	fmt.Println("result")

}
