package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//Using with struct

type CustomData interface{
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	Name string
	ID int
	Data T
}


func main() {
	user1 := User[int]{
		Name: "Munna Mia",
		ID: 22,
		Data: 231,
	}

	fmt.Println(user1)
}
