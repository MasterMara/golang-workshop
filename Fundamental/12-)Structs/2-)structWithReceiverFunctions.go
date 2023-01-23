package main

import "fmt"

type student struct {
	id   int
	name string
}

func createStudent() *student {

	myStudent := student{
		id:   2,
		name: "Kadir",
	}

	return &myStudent
}

//STUDENT Receiver Functions

func (s *student) getName() string {
	return s.name
}

//CallByValue
func (s student) changeIdByValue() {
	s.id = 100
}

//CallByReference
func (s *student) changeIdByReference() {
	s.id = 100
}

func main() {

	//1.WAY Create struct in main function
	myStudent := student{
		id:   1,
		name: "Mustafa",
	}

	fmt.Println("myStudentName is : ", myStudent.getName())

	//2.WAY Create Struct wiht a function
	myStudent2 := createStudent()

	//callByValue
	myStudent2.changeIdByValue()
	fmt.Println(myStudent2.id)

	//callByRefernece
	myStudent2.changeIdByReference()
	fmt.Println(myStudent2.id)

}

/*Some Notes By Mustafa KARACABEY:



 */
