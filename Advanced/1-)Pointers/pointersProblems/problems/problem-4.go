package main

import "fmt"

type myStructx struct {
	something int
}

func mainx() {

	//Empty struct Does not take up memory space
	str := myStructx{}

	fmt.Printf("Address of str %p\t value of str %v", &str, str)

}
