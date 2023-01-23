package main

import "fmt"

// Anonmouys function as an argument
func passedAnonmouysBro(myFunc func(age int) int) {

	fmt.Println(myFunc(15))
}

//function Returns anonmous Function
func returnedAnonmousFunctionBro(age int) func(age int, name string) {

	return func(age int, name string) {
		fmt.Println("My name:", name)
		fmt.Println("My age:", age)
	}

}

func main() {

	//1.WAY --> Without argument
	func() {
		fmt.Println("Hello World")
	}()

	//2.WAY -->With an argument
	func(name string) {
		fmt.Println("My name is: ", name)
	}("Mustafa")

	//3.WAY --> passing anonymous function as a parameter
	myFunc := func(age int) int {
		return age
	}

	passedAnonmouysBro(myFunc)

	//4.WAY Return Anonmous Function
	myReturnedFunc := returnedAnonmousFunctionBro(15)

	myReturnedFunc(10, "Mustafa")

}

/*Some Notes By Mustafa KARACABEY:

1-) Anonymous function has no NAMES !!!

*/
