package composition

import "fmt"

type shape interface {
	area() int
	perimeter() int
}

type rectangle struct {
	name string
	x    int
	y    int
}

type square struct {
	name string
	x    int
}

func (r *rectangle) area() int {
	return r.x * r.y
}

func (r *rectangle) perimeter() int {
	return 2 * (r.x + r.y)
}

func (s *square) area() int {
	return s.x * s.x
}

func (s *square) perimeter() int {
	return s.x ^ 2
}

func calculateTogether(shapes []shape) {
	for _, shape := range shapes {
		fmt.Println(shape.area())
		fmt.Println(shape.perimeter())
	}
}

func main() {

	myRectangle := rectangle{
		name: "rectangle",
		x:    12,
		y:    32,
	}

	mySquare := square{
		name: "square",
		x:    10,
	}

	shapes := make([]shape, 0, 2)

	shapes = append(shapes, &myRectangle)
	shapes = append(shapes, &mySquare)

	calculateTogether(shapes)
}

/*Some Notes By Mustafa KARACABEY:




 */
