package main

import "fmt"

// 1.WAY Variadic Parameter int
func sumValues(numbers ...int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// 2.WAY WİTH ARRAY
func sumArrayValues(array ...[]int) {

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func main() {

	myArray := []int{1, 2, 3, 4, 5}

	//1.WAY
	fmt.Println(sumValues(1, 2, 3, 4, 5))

	//2.WAY
	sumArrayValues(myArray)

}

/*Some Notes By Mustafa KARACABEY:



 */
