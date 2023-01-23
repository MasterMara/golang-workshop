package main

import "fmt"

type address struct {
	id       int
	location string
}

type user struct {
	id      int
	name    string
	address address
}

type userx struct {
	Name    string
	Age     int
	Address addressx
}

type addressx struct {
	City    string
	ApartNo int
}

func main() {

	myuser := user{
		id:   3,
		name: "Mustafa",
		address: address{
			id:       10,
			location: "Istanbul",
		},
	}

	fmt.Println(myuser)

	//Nested Struct with Anonmous struct
	//İç içe Struct kullanımı

	mySecondUser := userx{
		Name: "Ahmet",
		Age:  34,
		Address: struct { //I create a anonymous struct and also ı would create some address class and implement it
			City    string
			ApartNo int
		}{City: "Istanbul", ApartNo: 25},
	}

	fmt.Println(mySecondUser)
}

/*Some Notes By Mustafa KARACABEY:



 */
