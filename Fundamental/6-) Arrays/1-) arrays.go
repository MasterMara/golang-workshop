package main

import "fmt"

func main() {

	// 1.WAY --> already we have default [0,0,0] values if without initialize
	var numbers [3]int
	var cities = [3]string{"Istanbul", "Ankara", "Izmir"}

	// 2.WAY --> := Array Initializer
	progrramingLanguages := [3]string{"go", "c", "c-sharp"}

	// 3.WAY --> Create with Makte function
	names := make([]string, 3)
	names[0] = "Mustafa"

	fmt.Println(numbers)
	fmt.Println(cities)
	fmt.Println(progrramingLanguages)
	fmt.Println(names)

}

/*Some Notes By Mustafa KARACABEY:

1-)Slice and Arrays are mutable data types in golang, t
this means the value of the elements in slice or array can be changed after initialization without re-allocations of memory

2-) Array is fixed size

3-) The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).


*/
