package main

import "fmt"

func main() {

	var a int = 10
	var p *int = &a
	var pp **int = &p

	// 1.WAY Pointers are just variables that stores another variable memory address as a value
	fmt.Printf("A value is : %v \t a memory address is : %p\n", a, &a)
	fmt.Printf("P Pointer value is : %v \t memory address is : %p\n", p, &p)
	fmt.Printf("PP Pointer value is : %v \t memory address is : %p\n", pp, &pp)

	// 2.WAY Deferencing Pointers
	fmt.Printf("Deferencing p value is : %d \t pp value is : %d\n", *p, **pp)
}

/*Some Notes By Mustafa KARACABEY:

1-)Perfect Reference: https://www.youtube.com/watch?v=yEiaCx0fR9k&ab_channel=GopherConUK
2-) No Pointer arithmetic in golang
3-) Garbage Collector hanles deallocation
4-) Pointers are variables that store the memory address of another variable
5-) * --> retrieves the indirect value(deference)
6-) & --> Take The address
7-) Every Variable has address and value
8-) Pointers arev values,which happen to store address
9-) When to Use Pointers:
	- Avoid copying the value on each method call.This can be more efficient if receiver is large struct
	- method can modify the value that is receiver points to

*/
