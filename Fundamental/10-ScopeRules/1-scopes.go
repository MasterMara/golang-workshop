package main

import "fmt"

/* global variable declaration */
var g int

func main() {

	/* local variable declaration */
	var a, b int

	/* actual initialization */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)

}

/* function to add two integers */
func sum(a, b int) int {
	fmt.Printf("value of a in sum() = %d\n", a)
	fmt.Printf("value of b in sum() = %d\n", b)

	return a + b
}

/*
	1-)A program can have the same name for local and global variables but the value of the local variable inside a function takes preference
	2-)Formal parameters are treated as local variables with-in that function and they take preference over the global variables. For example
	3-) In addition to explicit blocks in the source code, there are implicit blocks (SCOPES):
		The universe block encompasses all Go source text.
		Each package has a package block containing all Go source text for that package.
		Each file has a file block containing all Go source text in that file.
		Each "if", "for", and "switch" statement is considered to be in its own implicit block.
		Each clause in a "switch" or "select" statement acts as an implicit block.
*/
