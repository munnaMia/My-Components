package main

import (
	"fmt"
	"math"
)


type shape interface{
	area() float32
	perimeter() float32
}


// Rectangler area
type rect struct {
	width, height float32
}

func (r rect) area() float32 {
	return r.height * r.width
}

func (r rect) perimeter() float32 {
	return 2 * (r.height + r.width)
}

// Circle area
type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}


func printdata (s shape){
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}


func main (){
	var squar rect = rect{
		height: 23.3,
		width: 33.3,
	}

	var circle circle = circle{
		radius: 10.2,
	}


	printdata(circle)
	printdata(squar)

	
	
}