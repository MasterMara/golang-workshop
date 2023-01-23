package composition

import "fmt"

type shape interface {
	area() int
	perimeter() int
}

type rectangle struct {
	shortEdge int
	longEdge  int
	sides     int
}

type square struct {
	edge int
}

func (r *rectangle) area() int {

	return r.longEdge * r.shortEdge
}

func (r *rectangle) perimeter() int {

	return (r.longEdge + r.shortEdge) * 2
}

func (r *rectangle) getSides() int {

	return r.sides
}

func (s *square) area() int {

	return s.edge * s.edge
}

func (s *square) perimeter() int {

	return s.edge * 4
}

func showAreaAndPerimeter(s shape) {

	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
func main() {

	rectangle := rectangle{shortEdge: 10, longEdge: 20}
	square := square{edge: 10}

	showAreaAndPerimeter(&rectangle)
	showAreaAndPerimeter(&square)

}

/*Some Notes By Mustafa KARACABEY:


1-) the io.Reader and io.Writer is most commongly use interfaces in go
2-) interface is an abstract type does not care implementation. Just define function signature
3-) Normally you should not declare an interface in tthe same package that declares the types that implement it
4-) Rectangle is a Shape and square is a shape
5-)A Type assertion provides access to an interface  values underlying concrete types
6-) GOOD RESOURCE : https://www.youtube.com/playlist?list=PLnhcQHG_PcWPmQSiagT5hPxQOxWAIkRDl


*/
