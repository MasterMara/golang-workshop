package main

import "fmt"

var ptr *int

type User struct {
	Name string
}

func pointerToNumberAsCopy(myNumber *int) {
	myNumber = nil // it change value of it :) but not original because it copy as usual :)
	fmt.Printf("myNumber Memory Adresses : %p\t myNumber value : %p, myNumber dereferrencing : nil\n", &myNumber, myNumber)
}

func pointerToNumberAsTry(myNumber *int) { //This is again copy ?? //Bir alltaki fonksiyon ile neden aynı çıktıyı verdi
	*myNumber = 45
	fmt.Printf("myNumber Memory Adresses : %p\t myNumber value : %p, myNumber dereferrencing : %d\n", &myNumber, myNumber, *myNumber)
}

func pointerToRef(myNumber **int) { //Pointer tutan pointer olur :)
	**myNumber = 45
	fmt.Printf("myNumber Memory Adresses : %p\t myNumber value : %p, myNumber dereferrencing : %d\n", &myNumber, myNumber, myNumber)
}

func changeName(myUserPtr *User) {
	myUserPtr.Name = "Kadir" // Go Auto Change this -->  (*myUserPtr).Name = "Kadir"
}
func changeNameAsRef(myUserPtr **User) {
	(**myUserPtr).Name = "Kadir" // Go Auto Change this --> (*myUserPtr).Name = "Kadir"
}

func main() {

	a := 10
	ptr = &a

	fmt.Printf("a Memory Adresses : %p\t a value : %d\n", &a, a)
	fmt.Printf("Ptr Memory Adresses : %p\t ptr value : %p, Ptr dereferrencing : %d\n", &ptr, ptr, *ptr)
	fmt.Println("********************************")

	*ptr = 20 // It will chabnge original value of course :)
	fmt.Printf("a Memory Adresses : %p\t a value : %d\n", &a, a)
	fmt.Printf("Ptr Memory Adresses : %p\t ptr value : %p, Ptr dereferrencing : %d\n", &ptr, ptr, *ptr)
	fmt.Println("**********************************")

	pointerToNumberAsCopy(ptr) //Nothing Will be Change :) because myNumber pointer is copy of actual pointer
	fmt.Printf("Ptr Memory Adresses : %p\t ptr value : %p\t, Ptr dereferrencing : %d\n", &ptr, ptr, *ptr)

	fmt.Println("**********************************")

	fmt.Printf("Ptr Memory Adresses : %p\t ptr value : %p, Ptr dereferrencing : %d\n", &ptr, ptr, *ptr)
	pointerToNumberAsTry(ptr)
	fmt.Printf("Ptr Memory Adresses : %p\t ptr value : %p, Ptr dereferrencing : %d\n", &ptr, ptr, *ptr)
	fmt.Println("**********************************")

	fmt.Printf("Ptr Memory Adresses : %p\t ptr value : %p, Ptr dereferrencing : %d\n", &ptr, ptr, *ptr)
	pointerToRef(&ptr)
	fmt.Printf("Ptr Memory Adresses : %p\t ptr value : %p, Ptr dereferrencing : %d\n", &ptr, ptr, *ptr)

	fmt.Println("************************************************************")

	myUser := User{
		Name: "Mustafa",
	}

	fmt.Printf("MyStruct Addresses  : %p\t MyStruct Value : %s\n", &myUser, myUser.Name)

	myUserPtr := &myUser
	fmt.Printf("myUserPtr Addresses  : %p\t myUserPtr points : %p\t myUserPtr Dereferring : %s\n", &myUserPtr, myUserPtr, (*myUserPtr).Name)
	fmt.Println("************************************************************")

	changeName(myUserPtr)
	fmt.Printf("myUserPtr Addresses  : %p\t myUserPtr points : %p\t myUserPtr Dereferring : %s\n", &myUserPtr, myUserPtr, (*myUserPtr).Name)
	fmt.Println("************************************************************")

	changeNameAsRef(&myUserPtr)
	fmt.Printf("myUserPtr Addresses  : %p\t myUserPtr points : %p\t myUserPtr Dereferring : %s\n", &myUserPtr, myUserPtr, (*myUserPtr).Name)
}

/*Some Notes By Mustafa KARACABEY:

1-)

*/
