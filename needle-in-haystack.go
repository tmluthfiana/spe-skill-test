package main

import (
	"fmt"
	"reflect"
)

func findNeedle(val interface{}, array interface{}) (index int) {
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				// exists = true
				return
			}
		}
	}

	return
}

func main() {
	arrayTest := []string{"red", "blue", "yellow", "black", "grey"}
	result := findNeedle("blue", arrayTest)
	fmt.Println(result)
}
