package main

import "fmt"

// 1.WAY Recursive funtion example
func takeFactorial(number int) int {

	if number == 0 {
		return 1
	}

	return takeFactorial(number-1) * number
}

func main() {

	fmt.Println("5! is : ", takeFactorial(5))
}

/*Some Notes By Mustafa KARACABEY:

1-) Recursive Function is whole different world. This is a just example

*/
