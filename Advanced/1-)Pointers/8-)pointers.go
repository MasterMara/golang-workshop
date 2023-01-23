package main

import "fmt"

func main() {

	// 1.WAY Pointers holds pointers hold pointers ...
	a := 1
	fmt.Printf("a : %#v\t  a memory adresses : %p\n", a, &a)

	b := &a
	fmt.Printf("b : %#v\t b value is : %p\t b memory adresses : %p\n", b, b, &b)

	c := &b
	fmt.Printf("c : %#v\t c value is : %d\t c memory adresses :%p\n", c, **c, &c)

	d := &c
	fmt.Printf("d : %#v\t d value is :%d\t d memory adresses : %p\n", d, ***d, &d)

}

/*Some Notes By Mustafa KARACABEY:

1-) %#v --> (*int)(0xc000016088)  like this

*/
