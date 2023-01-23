package main

import "fmt"

func main() {

	// CASE 1 is normal declaration
	myDictionary := map[string]int{}

	myDictionary["one"] = 1
	myDictionary["two"] = 2
	myDictionary["three"] = 3

	for key, value := range myDictionary {
		fmt.Println("key:", key, "value:", value)
	}
	fmt.Println("************************")

	//Case 2 is with make but has no value
	myMap := make(map[string]int)
	myMap2 := make(map[int]int, 3)

	myMap2[0] = 1
	myMap2[1] = 2
	myMap2[2] = 3
	myMap2[3] = 4 //Even if size is overloaded there is no problem!!

	fmt.Println("My map len is ", len(myMap))

	for key, value := range myMap2 {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("Len is ", len(myMap2))

}

/*Some Notes By Mustafa KARACABEY:

1-) Maps data structure is like dictionary pr hash map, key-value pair


*/
