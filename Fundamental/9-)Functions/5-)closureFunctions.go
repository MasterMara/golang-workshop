package main

import "fmt"

// Function return function
func intSequence() func() int {

	counter := 0 // Local Scope

	//Anonymous Func It will take i reference so never forget :)
	return func() int {
		counter++
		return counter
	}
}

func main() {

	myValue := 10

	//1.WAY This function return function
	myIntSequence := intSequence()
	fmt.Println(myIntSequence())
	fmt.Println(myIntSequence())

	//2.WAY Closure function --> Closure functions access outer variable here is = myValue
	myCounter := func() int {
		myValue++
		return myValue
	}

	fmt.Println(myCounter())
	fmt.Println(myCounter())

}

/*Some Notes By Mustafa KARACABEY:

1-) Anonmous function can form closure in golang.
2-) closure is just a function value that references variables from outside of its own functions body so
	inner function that remembers and has access to variables in the local scope in which it was created even after the outer
	function is done executing

3-) Closure function acces outer variable :)
4-) inner function that remembers and has access to variables in the local scope in which it was created even after the outer


*/
