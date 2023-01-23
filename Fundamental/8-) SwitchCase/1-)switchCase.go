package main

import "fmt"

func main() {

	number := 10

	// 1.WAY Classical switch case usage
	switch number { // You can assign number here also

	case 1:
		fmt.Println("My Number is: ", number)
	case 2:
		fmt.Println("My Number is: ", number)
	default:
		fmt.Println("Default is: ", number)
	}

	// 2.WAY Speacial usage of interface
	var a interface{}
	a = 5
	switch a.(type) {
	case int:
		fmt.Println("an int.")
	case int32:
		fmt.Println("an int32.")
	}
	// in this case it will print "an int."

}

/*Some Notes By Mustafa KARACABEY:

1-) In Golang Switch case has no need break :)



*/
