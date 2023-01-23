package main

import "fmt"

func main() {

	// Case 1 create Slice
	var mySlice1 []string
	fmt.Println("My slice Cap is ", cap(mySlice1), "Len is ", len(mySlice1))

	// Case2 create slice
	mySlice2 := []string{}
	fmt.Println("My slice Cap with := declaration is ", cap(mySlice2))

	// Case 3 Create Slice from array
	var myArray = [6]int{1, 2, 3, 4, 5, 6}
	mySlice3 := myArray[2:5]
	fmt.Println("My Slice3 cap is :", cap(mySlice3), "Len is ", len(mySlice3), "My slice3 is :", mySlice3)

	//Case 4 create Slice with make
	mySlice4 := make([]int, 1, 3) // 1 is len, 3 is cap (will create with default value)
	for _, item := range mySlice4 {
		fmt.Println(item)
	}
	mySlice4 = append(mySlice4, 10) //Add data to end of slice

	//Case  5 with omitted capacity cap is equal to len
	myslice22 := make([]int, 5)
	fmt.Println(myslice22)

	// Example Case Slice can extend !!!
	slc_2 := make([]int, 0)

	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))
}

/*Some Notes By Mustafa KARACABEY:

1-)- Reference : https://www.w3schools.com/go/go_slices_modify.php
2-) Reference: https://www.developer.com/languages/arrays-slices-golang/#:~:text=Slices%20in%20Go%20and%20Golang,the%20start%20and%20end%20index.
2-) Slices are similar to arrays, but are more powerful and flexible.
3-) However, unlike arrays, the length of a slice can grow and shrink as you see fit.
4-) Using the []datatype{values} format
	Create a slice from an array
	Using the make() function
5-) Array has fixed capacity but slice is not var myArray [3]string{}
6-)len() function - returns the length of the slice (the number of elements in the slice)
7-)cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
8-) In My Opinion Slice = List similar
9-) Memory Efficiency When using slices, Go loads all the underlying elements into the memory.
10-) If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
    The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
11-) To define a slice, you can declare it as an array without specifying its size.
12-) If a slice is declared with no inputs, then by default, it is initialized as nil. Its length and capacity are zero.


*/
