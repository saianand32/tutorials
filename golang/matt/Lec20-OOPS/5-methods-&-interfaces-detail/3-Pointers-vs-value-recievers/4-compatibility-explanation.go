package main

type Point struct {
	x int
	y int
}

func (p *Point) Add(a, b int) {
	p.x = a
	p.y = b
}
func main() {
	// in prev slide what is mentioned is in go you  can give
	// a val to a pointer reciever method and go will know if to call its adreesable location or value based on if it is
	// a pointer or val reciever
	// in all cases except one -

	// example -
	// here i cant do this
	Point{1, 2}.Add(2, 3)

	// as Point is not addressable as its a literal u will have to do either of below

	p1 := Point{1, 2}
	p1.Add(3, 4) //go auto determines that address must be caled

	(&p1).Add(4, 5) //you can also specify that address call happens
}
