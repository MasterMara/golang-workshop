package main

import "fmt"

func main() {

	// 1.WAY [Row][Column]  or you can just define with var
	myMatrix := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	showMultiDimensionalArrayOutput(myMatrix)

}

func showMultiDimensionalArrayOutput(matrix [3][2]int) {

	for _, row := range matrix {
		for _, item := range row {
			fmt.Println(item)
		}
	}

}

/*Some Notes By Mustafa KARACABEY:




 */
