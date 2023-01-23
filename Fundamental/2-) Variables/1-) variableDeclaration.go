package main

import "fmt"

func main() {

	// 1.WAY
	var isTrue bool             //--> Without any Assignment
	var name string = "Mustafa" // With Initialize

	// 2.WAY
	age := 18

	// 3.WAY --> Multiple variable declaration without Initialize and with Initialize
	var (
		firstName string
		lastName  string
		job       = "Software Engineer"
	)

	// 4.WAY --> Multiple variable declaration with single line  without Initialize and with Initialize
	a, b := 5, 10

	fmt.Println("My Variable value is :", isTrue)
	fmt.Println("My name is :", name)
	fmt.Println("My Age is : ", age)
	fmt.Println("My FirstName is : ", firstName)
	fmt.Println("My LastName is : ", lastName)
	fmt.Println("My Job is : ", job)
	fmt.Println("a is :", a)
	fmt.Println("b is : ", b)

}

/*Some Notes By Mustafa KARACABEY:

1-) Default value is false for "bool"
2-) var = Can be used inside and outside of functions, Variable declaration and value assignment can be done separately
	:= Can only be used inside functions, Variable declaration and value assignment cannot be done separately (must be done in the same line)


*/
