package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Create a maping funciton like JS in go

// Normal way: using function

func mapValue(arr []int, mapFunk func(int) int) []int {
	var newValuesArr []int

	for _, v := range arr {
		updateValue := mapFunk(v)
		newValuesArr = append(newValuesArr, updateValue)
	}

	return newValuesArr
}

// Let's update it Generics...
func mapValueUpdate[T constraints.Ordered](arr []T, mapFunk func(T) T) []T {
	var newValuesArr []T

	for _, v := range arr {
		updateValue := mapFunk(v)
		newValuesArr = append(newValuesArr, updateValue)
	}

	return newValuesArr
}

func main() {
	test := mapValueUpdate([]float64{1.1, 2.9, 3, 4, 5}, func(n float64) float64 { return n * 2 })
	fmt.Println(test)
}
