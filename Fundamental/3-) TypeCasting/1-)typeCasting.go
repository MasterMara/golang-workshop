package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		number1 int
		number3 float64
		number6 int64
		err     error
	)

	number2, number4, name1 := 2.2, 10, "125"

	// 1.WAY --> Float to int
	number1 = int(number2)

	// 2.WAY --> int to Float
	number3 = float64(number4)

	// 3.WAY --> String to int
	number6, err = strconv.ParseInt(name1, 10, 64)

	if err != nil {
		fmt.Println("Conversion has a problem !! for number6 ERROR:", err)
	}

	fmt.Println("Float to int typeCasting Number1 is: ", number1)
	fmt.Printf("int to float typeCasting Number3 is:%f\n", number3)
	fmt.Println("String to int typeCasting Number6 is: ", number6)
}

/*Some Notes By Mustafa KARACABEY:

1-) Golang does not support implicit type Casting. Which means compiles does automaticly.
2-) Some Conversion Functions:
	- int() --> Maybe Float to int

3-) strconv.atoi = strconv.ParseInt(number, 10, 0)

*/
