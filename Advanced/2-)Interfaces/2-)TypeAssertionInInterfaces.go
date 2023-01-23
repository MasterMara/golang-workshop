package composition

import "fmt"

type shapexx interface {
	area() int
}

type rectanglexx struct {
	shortEdge int
	longEdge  int
}

type squarexx struct {
	edge int
}

func (r *rectanglexx) area() int {
	return r.longEdge * r.shortEdge
}

func (r *rectanglexx) getShortEdge() int {
	return r.shortEdge
}

func (s *squarexx) area() int {
	return s.edge * s.edge
}

func calculateArea(s shapexx) {

	fmt.Printf("s Type is : %T\t s value is : %v\n", s, s)

	//1.WAY TO CREATE TYPE ASSERTION
	myRect := s.(*rectanglexx) // Check error handling here :)

	//mysquare := s.(*square) -->This will create painc which means runtime error why because of  square has no other function just implements interfaces methods
	fmt.Println("My Rectanlge short value inside s shape functin with type assertion is : ", myRect.getShortEdge())
}

func main() {

	rect := rectanglexx{shortEdge: 10, longEdge: 20}

	calculateArea(&rect)

}

/*Some Notes By Mustafa KARACABEY:
1-) Type Assertion
	myRect := s.(*rectangle)

2-) mysquare := s.(*square) -->This will create painc which means runtime error why because of  square has no other function just implements interfaces methods



*/
