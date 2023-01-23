package main

import "fmt"

var ap *int

func main() {

	// 1.WAY PERFECT UNDERSTANDING OF POINTERS IN GOLANG
	fmt.Printf("ap memory addresses: %p\t ap points : %p\tp value :nill\n", &ap, ap)
	fmt.Println("**************************************************")
	a := 1 // define int
	b := 2 // define int

	fmt.Printf("a memory addresses : %p\t a value is : %d\n", &a, a)
	fmt.Printf("b memory addresses : %p\t b value : %d\n", &b, b)

	fmt.Println("**************************************************")

	ap = &a
	// set ap to address of a (&a)
	//   ap address: 0x2101f1018
	//   ap value  : 1

	fmt.Printf("ap memory addresses: %p\t ap points : %p\tp referred value :%d\n", &ap, ap, *ap)

	fmt.Println("**************************************************")
	*ap = 3
	// change the value at address &a to 3
	//   ap address: 0x2101f1018
	//   ap value  : 3
	fmt.Printf("ap memory addresses: %p\t ap points : %p\tap deferred value :%d\n", &ap, ap, *ap)
	fmt.Printf("a memory addresses : %p\t a value is : %d\n", &a, a)
	fmt.Println("**************************************************")
	a = 4
	// change the value of a to 4
	//   ap address: 0x2101f1018
	//   ap value  : 4
	fmt.Printf("a memory addresses : %p\t a value is : %d\n", &a, a)
	fmt.Printf("ap memory addresses: %p\t ap points : %p\tp referred value :%d\n", &ap, ap, *ap)

	fmt.Println("**************************************************")
	ap = &b
	// set ap to the address of b (&b)
	//   ap address: 0x2101f1020
	//   ap value  : 2
	fmt.Printf("ap memory addresses: %p\t ap points : %p\tp referred value :%d\n", &ap, ap, *ap)
	fmt.Printf("b memory addresses : %p\t b value : %d\n", &b, b)

	fmt.Println("**************************************************")

	ap2 := ap //Assing pointer to another pointer. It will point the same with ap pointer
	// set ap2 to the address in ap
	//   ap  address: 0x2101f1020
	//   ap  value  : 2
	//   ap2 address: 0x2101f1020
	//   ap2 value  : 2
	fmt.Printf("ap memory addresses: %p\t ap points : %p\tp referred value :%d\n", &ap, ap, *ap)
	fmt.Printf("ap2 memory addresses: %p\t ap2 points : %p\tap2 referred value :%d\n", &ap2, ap2, *ap2)

	fmt.Println("**************************************************")

	*ap = 5
	// change the value at the address &b to 5
	//   ap  address: 0x2101f1020
	//   ap  value  : 5
	//   ap2 address: 0x2101f1020
	//   ap2 value  : 5
	// If this was a reference ap & ap2 would now
	// have different values
	fmt.Printf("b memory addresses : %p\t b value : %d\n", &b, b)
	fmt.Printf("ap memory addresses: %p\t ap points : %p\tp referred value :%d\n", &ap, ap, *ap)
	fmt.Printf("ap2 memory addresses: %p\t ap2 points : %p\tap2 referred value :%d\n", &ap2, ap2, *ap2)

	fmt.Println("*********************STOP*****************************")
	ap = &a
	// change ap to address of a (&a)
	//   ap  address: 0x2101f1018
	//   ap  value  : 4
	//   ap2 address: 0x2101f1020
	//   ap2 value  : 5
	// Since we've changed the address of ap, it now
	// has a different value then ap2

	fmt.Printf("ap memory addresses: %p\t ap points : %p\tp referred value :%d\n", &ap, ap, *ap)
	fmt.Printf("ap2 memory addresses: %p\t ap2 points : %p\tap2 referred value :%d\n", &ap2, ap2, *ap2)

}

/*Some Notes By Mustafa KARACABEY:

1-) If a assing pointer to another pointer which means a:= b  a will point WHERE the b points

*/
