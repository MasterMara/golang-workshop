package main

import "fmt"

func main() {

	fruits := [3]string{"apple", "orange", "banana"}
	counter := 3

	// 1.WAY --> Tree Component Loop = Classical loop
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// 2.WAY --> Range Loop with index and withoud index
	for index, fruit := range fruits {
		fmt.Printf("Index is : %d and value is :%s\n", index, fruit)
	}

	for _, fruit := range fruits {
		fmt.Println("Fruit is : ", fruit)
	}

	// 3.WAY --> for usage for infinite loops
	for {
		fmt.Println("Your Counter value is :", counter)
		if counter > 5 {
			break
		}
		counter++
	}

}

/*Some Notes By Mustafa KARACABEY:

1-) Continue, and Break keywords are same
	Continue: ignore if conditions ok
	Break: Finished the loop if conditions ok



*/
