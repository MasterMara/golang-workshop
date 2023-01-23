package main

import "fmt"

type myStruct struct {
	something int
}

func (m *myStruct) structGoesBoom() {
	m = nil
}

func problem5() {
	x := myStruct{1999} // x is copy
	fmt.Printf("Address:%p\n", &x)

	x.structGoesBoom() //Does not change anything
	fmt.Printf("%#v\n", x)

}
