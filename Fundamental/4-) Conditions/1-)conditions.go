package main

import (
	"fmt"
)

func main() {

	// 1.Way Exactly Classical if, else-if else logic
	age := 18

	if age >= 18 {
		fmt.Println("Your are Man")
	} else if age < 18 && age > 10 {
		fmt.Println("You are child")
	} else {
		fmt.Println("You are kid")
	}

}

/*Some Notes By Mustafa KARACABEY:

1-) There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions


*/
