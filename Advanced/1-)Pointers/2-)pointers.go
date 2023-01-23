package main

import "fmt"

//Pointers are just variables this is a proof not references
//value and ptr2 they are not same variable is just copy thats why address is different just copy of value
func destroy(value *int) {

	fmt.Printf("value memory Adress : %p\t value points Address : %p\t value value : %d\n", &value, value, *value)
	value = nil
	fmt.Printf("After value = nil inside function == value memory Adress : %p\t value points Address : %p\t value value : nil\n", &value, value)
}

func main() {

	// 1.WAY TO define pointer
	//var ptr *int

	// 2.WAY TO define pointer
	ptr2 := new(int)

	fmt.Printf("PTR2 memory Adress : %p\t PTR2 points Address : %v\t PTR2 value : %v\n", &ptr2, ptr2, *ptr2)

	*ptr2 = 10

	fmt.Printf("PTR2 memory Adress : %p\t PTR2 points Address : %p\t PTR2 value : %d\n", &ptr2, ptr2, *ptr2)

	destroy(ptr2)
	fmt.Println("after destroy function in main******************************************")
	fmt.Printf("Answer memory Adress : %p\t Answer points Address : %p\t answer value : %d\n", &ptr2, ptr2, *ptr2)

}

/*Some Notes By Mustafa KARACABEY:
1-) This experimental job is proof that pointers are just variables.

*/
