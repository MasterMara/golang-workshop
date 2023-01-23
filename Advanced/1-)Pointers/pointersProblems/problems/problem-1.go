package main

import "fmt"

// What Is The myPtrAddress and myPtrAddress  of myPtr ?
func problem1() {
	myPtr := new(int)
	fmt.Printf("myPtrAddress : %p, myPtrAddress : %v", myPtr, *myPtr)
}

/*
a-) Does not Compile
b-)myPtrAddress : 0xc0000aa058, myPtrValue : 0

*/
