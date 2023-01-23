package main

import "fmt"

//1.WAY --> These are custom types.
type city string
type age int

func main() {

	//We create a city type
	myFirstCity := city("Istanbul")
	myAge := age(15)

	fmt.Println(myFirstCity, myAge)

}

/*Some Notes By Mustafa KARACABEY:

1-) In Golang you can create  custom types


*/
