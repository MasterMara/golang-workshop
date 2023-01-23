package main

import "fmt"

type book struct {
	id   int
	name string
}

func main() {

	//1.WAY to Create Struct
	var myBook book
	myBook.id, myBook.name = 3, "Osman Ghazi"
	fmt.Println(myBook)

	//2.WAY to create Struct
	myBook2 := book{}
	myBook2.id, myBook2.name = 4, "The Conquerer Mehmet"
	fmt.Println(myBook2)

	//3.WAY initialize structs when declaring
	myBook3 := book{
		id:   3,
		name: "Selim",
	}
	fmt.Println(myBook3)

	//4.WAY Anonymous Struct

	myAnonymousStruct := struct {
		Name string
		Age  int
	}{"Ahmet", 22}

	fmt.Println("My Anoymous Struct is  : ", myAnonymousStruct)

}

/*Some Notes By Mustafa KARACABEY:
1-) 	str := myStructx{} --> Empty struct Does not take up memory space
2-) struct is a concrete type while interfaces are abstract types.


*/
