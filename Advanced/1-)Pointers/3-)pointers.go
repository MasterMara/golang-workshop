package main

import "fmt"

type studentStr struct {
	id int
}

//Pointers are pass by copy like anything else
func (s *studentStr) goesBoom() {

	fmt.Printf("s memory Adress : %p\t s points : %p\t s value is : %v\n", &s, s, s)
	s = nil // s is copy of actual value !!!
	fmt.Printf(" after nil s memory Adress : %p\t s points : %p\t s value is : %v\n", &s, s, s)
}

func main() {

	mustafa := studentStr{id: 10}

	fmt.Printf("mustafa memory Adress : %p\t mustafa value is : %v\n", &mustafa, mustafa)

	mustafa.goesBoom()

	fmt.Printf(" after boom in main : mustafa memory Adress : %p\t mustafa value is : %v\n", &mustafa, mustafa)

}

/*Some Notes By Mustafa KARACABEY:
1-) INVESTIGATE THIS EXAMPLE VERY DEEPLY
*/
