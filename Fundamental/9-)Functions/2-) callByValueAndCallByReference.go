package main

import "fmt"

// 1.WAY CallByValue
func swapWithCallByValue(number1, number2 int) {

	var temp int

	temp = number1
	number1 = number2
	number2 = temp
}

// 1.WAY CallByReference
func swapWithCallByReference(number1 *int, number2 *int) {

	var temp int

	temp = *number1
	*number1 = *number2
	*number2 = temp
}

func main() {

	number1, number2 := 5, 10

	// callByValue
	swapWithCallByValue(number1, number2)

	fmt.Printf(" Call by Value --> Number1 is = %d \t number2 is = %d \n", number1, number2)

	// callByReference
	swapWithCallByReference(&number1, &number2)

	fmt.Printf(" Call by Reference --> Number1 is = %d \t number2 is = %d ", number1, number2)
}

/*Some Notes By Mustafa KARACABEY:



 */
