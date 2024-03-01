package main

import (
	"fmt"
	"math"
)

//in go interfaces are implemented implicitly

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length int
	width  int
}

func (r rect) area() float64 {
	return float64(r.length) * float64(r.width)
}

func (r rect) perimeter() float64 {
	return float64(2*r.length) + float64(2*r.width)
}

type circle struct {
	radius int
}

func (r circle) area() float64 {
	return math.Pi * float64(r.radius) * float64(r.radius)
}

func (r circle) perimeter() float64 {
	return 2 * math.Pi * float64(r.radius)
}

//from above u can see that the rect struct & circle struct will auto implements the shape interface as they  both use both methods of shape interface..

// Now i can make general methods for any struct implementng interface
func getArea(shp shape) {
	fmt.Println(shp.area())
}

func main() {
	rect1 := rect{
		length: 100,
		width:  50,
	}

	getArea(rect1) //i can pass any struct that implements the shape interface
}
