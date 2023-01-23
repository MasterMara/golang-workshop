package main

import "fmt"

const MAX = 3

func swapValuesByReferences(a, b *int) {

	var temp int

	temp = *a
	*a = *b
	*b = temp

}

func main() {

	// 1-WAY: Define Pointer Array

	myArray := []int{100, 200, 300}

	for i := 0; i < MAX; i++ {
		fmt.Println(myArray[i])
	}

	var myPtrArray [MAX]*int

	for i := 0; i < MAX; i++ {
		myPtrArray[i] = &myArray[i]
	}

	for j := 0; j < MAX; j++ {
		fmt.Printf("Value of a[%d] = %d\n", j, *myPtrArray[j])
	}

	//2.WAY SWAP WİTH CALL BY REFERENCE
	a, b := 10, 20

	fmt.Printf("a value is : %v\t a memory address is : %p\n", a, &a)
	fmt.Printf("b value is : %v\t b memory address is : %p\n", b, &b)

	swapValuesByReferences(&a, &b)

	fmt.Println("*****************AFTER SWAP*****************")

	fmt.Printf("a value is : %v\t a memory address is : %p\n", a, &a)
	fmt.Printf("b value is : %v\t b memory address is : %p\n", b, &b)

	// 3.WAY pointers hold pointer
	var numberx int = 10
	var myptr1 *int
	var myptr2 **int

	myptr1 = &numberx
	myptr2 = &myptr1

	fmt.Println("*****************POINTER HOLDS POINTER RESULT*****************")

	fmt.Printf("Numberx value is : %v\t Numberx memory address is : %p\n", numberx, &numberx)
	fmt.Printf("myptr1 value is : %v\t myptr1 memory address is : %p\t myprtr1 pointer address is : %p\n", *myptr1, &myptr1, myptr1)
	fmt.Printf("myptr2 value is : %v\t myptr2 memory address is : %p\t myprtr1 pointer address is : %p\n", **myptr2, &myptr2, myptr2)

}

/*Some Notes By Mustafa KARACABEY:

1-) we take a memoryAddress and if  myPtrArray changes then myArray will change
2-) If one pointer points another pointer which means that pointer points other pointers memory address.

*/
