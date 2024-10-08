package main

import (
	"fmt"
)

/*
	Here we have a function. What is does just add two numbers.
	Now if we focus we can see as a type language , we can pass
	a single type of value as parameter. What if we want to add
	a floating point number or int number and any type of number
	we can't do this in a function normal way.

	So here comes Generics.

		=> func add[T int | float64](a T, b T) T {
		=> 	return a + b
		=> }
*/

type Num interface {
	int | float32 | float64
}

/*
Here You can use the type ""NUM"" to work with any types of data
OR
Use :: constraints.Ordered --> It will handle all the dirty work for us

	===>func add[T constraints.Ordered](a T, b T) T {
	===>	return a + b
	===>}
*/

func add[T Num](a T, b T) T {
	return a + b
}

/*
	Another Example:
		1. Create a USERID datatype
		2. Create X, Y variable based on USERID type
		3. Pass X, Y to Minus() funciton

		What the "~" does here it will allow a variable to pass
		if the underlaying type is matched.

*/

type USERID int

func Minus[T ~int | ~float64](a T, b T) T {
	return a - b
}

func main() {
	x := USERID(10)
	y := USERID(11)
	r := Minus(x, y)
	fmt.Println(r)
}
