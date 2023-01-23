package main

import "fmt"

// 1.WAY = classical functions with return nothing
func sayHello() {
	fmt.Println("Hi, Welcome Functions. !!!")
}

// 2.WAY = Return single value
func sumNumbers(number1 int, number2 int) int {
	return number1 + number2
}

// 3.WAY = Return multiple values
func swapNumbers(number1 int, number2 int) (int, int) {
	return number2, number1
}

func main() {

	var (
		number1    = 10
		number2    = 20
		newNumber1 int
		newNumber2 int
	)

	sayHello()

	fmt.Println("2 + 3 is = ", sumNumbers(10, 20))

	newNumber1, newNumber2 = swapNumbers(number1, number2)
	fmt.Printf("Normal a = %d \t b = %d\n new a = %d \t new b = %d ", number1, number2, newNumber1, newNumber2)
}

/*Some Notes By Mustafa KARACABEY:



 */
