package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
}

func main() {

	name := "mustafa"
	person := person{name: "Mustafa"}

	mySlice := make([]int, 0, 2)

	mySlice = append(mySlice, 1, 2)

	fmt.Println(reflect.ValueOf(name), reflect.TypeOf(name))
	fmt.Println(reflect.ValueOf(mySlice), reflect.TypeOf(mySlice))

	myPersonTypeOf := reflect.TypeOf(person)

	fmt.Println("My Name  TYPE is : ", reflect.TypeOf(name))

	fmt.Println("*********************************************")

	fmt.Println("My person Type is : ", reflect.TypeOf(person))
	fmt.Println("My  person Kind is : ", myPersonTypeOf.Kind())
}

/*Some Notes By Mustafa KARACABEY:

Reflection
1-) Types: type myInt int -->Typeis myInt and Kind is int
2-) Kinds: string, struct , int etc.
3-) Values
4-)Reflection is a program's runtime, especially through tips defined as the study of its own structure


*/
