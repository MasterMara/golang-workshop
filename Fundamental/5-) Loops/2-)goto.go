package main

import "fmt"

func main() {

	// 1.Way goto is equelas assmebly jumps
	fmt.Println("Step1")
	goto STEP3

	fmt.Println("Step2") //Skip this

STEP3:
	fmt.Println("Step3")

}

/*Some Notes By Mustafa KARACABEY:

- Why go has goto ? -->  https://stackoverflow.com/questions/11064981/why-does-go-have-a-goto-statement


*/
