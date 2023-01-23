package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "mustafa"

	// 1.WAY Classical string manupulatins
	str_1 := str[:5]
	str_2 := str[len(str)-4:]

	// 2.WAY you can all the way throug inside the function
	isContainsMWord := strings.Contains(str, "m")

	fmt.Println(str_1)
	fmt.Println(str_2)
	fmt.Println("IsContains M :", isContainsMWord)

}

/*Some Notes By Mustafa KARACABEY:

1-) strings.Contains() --> You can understand the signatures of functions inside strings package



*/
