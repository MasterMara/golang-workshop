package composition

import "fmt"

type shape interface {
	area() int
}

type rectangle struct {
	shortEdge int
	longEdge  int
}

type square struct {
	edge int
}

func (r *rectangle) area() int {
	return r.longEdge * r.shortEdge
}

func (r *rectangle) getShortEdge() int {
	return r.shortEdge
}

func (s *square) area() int {
	return s.edge * s.edge
}

func calculateArea(s shape) {

	fmt.Printf("s Type is : %T\t s value is : %v\n", s, s)

	//1.WAY TO CREATE TYPE ASSERTION
	myRect := s.(*rectangle) // Check error handling here :)

	//mysquare := s.(*square) -->This will create painc which means runtime error why because of  square has no other function just implements interfaces methods
	fmt.Println("My Rectanlge short value inside s shape functin with type assertion is : ", myRect.getShortEdge())
}

func main() {

	rect := rectangle{shortEdge: 10, longEdge: 20}

	calculateArea(&rect)

}

/*Some Notes By Mustafa KARACABEY:
1-) Type Assertion
	myRect := s.(*rectangle)

2-) mysquare := s.(*square) -->This will create painc which means runtime error why because of  square has no other function just implements interfaces methods



*/
