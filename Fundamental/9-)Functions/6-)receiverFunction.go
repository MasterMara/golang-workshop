package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) getName() string {
	return s.name
}

func main() {

	student1 := student{
		name: "Mustafa",
		age:  23,
	}

	fmt.Println("My name is : ", student1.getName())

}

/*Some Notes By Mustafa KARACABEY:



 */
