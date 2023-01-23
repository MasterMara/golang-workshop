package main

import "fmt"

type person struct {
	name string
}

func (p *person) changeName(name string) {

	fmt.Printf("in function: p memory address is : %p \t p value is : %v\t p points : %p\n", &p, *p, p)

	p.name = name

}

func createNewPersonPtr(name string) *person {

	myPerson := person{name: name}
	fmt.Printf(" PTR myPerson memory address is : %p \t myPerson value is : %v\n", &myPerson, myPerson)

	return &myPerson
}

func createNewPerson(name string) person {
	myCopyPerson := person{name: name}

	fmt.Printf("myCopyPerson memory address is : %p \t myCopyPerson value is : %v\n", &myCopyPerson, myCopyPerson)
	return myCopyPerson
}

func main() {

	mustafa := person{name: "Mustafa"}
	fmt.Printf("mustafa memory address is : %p \t mustafa value is : %v\n", &mustafa, mustafa)

	mustafa.changeName("kadir")
	fmt.Printf("mustafa memory address is : %p \t mustafa value is : %v\n", &mustafa, mustafa)
	fmt.Println("****************************************************************************************")

	mehmet := createNewPersonPtr("Mehmet")
	fmt.Printf("mehmet memory address is : %p \t mehmet value is : %v\t mehmet points : %p\n", &mehmet, mehmet, mehmet)
	fmt.Println("****************************************************************************************")

	ahmet := createNewPerson("Ahmet")
	fmt.Printf("ahmet memory address is : %p \t ahmet value is : %v\t Is ahmet points? : NO ", &ahmet, ahmet)

}

/*Some Notes By Mustafa KARACABEY:
 */
