package main

import "fmt"

func createSliceMap() map[string][]int {

	myMustafaMap := map[string][]int{}

	myMustafaMap["Mustafa"] = []int{1, 2, 3, 4, 5}

	return myMustafaMap
}

func showMapValue(myMap map[string][]int) {

	for key, value := range myMap {
		fmt.Println("Key", key, "Value:", value)
	}
}

func main() {

	mySliceMap := createSliceMap()
	showMapValue(mySliceMap)
}

/*Some Notes By Mustafa KARACABEY:




 */
