package main

import (
	"errors"
	"fmt"
	"math"
)

// 1.WAY = Simple Error Define
func getSqrt(number int) (int, error) {

	if number < 0 {
		return -1, errors.New("Negative Sqrt is not valid")
	}

	return int(math.Sqrt(float64(number))), nil
}

func main() {

	myNumber := -1

	value, err := getSqrt(myNumber)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("My %d sqrt is : %d", myNumber, value)
	}

}

/*Some Notes By Mustafa KARACABEY:

1-) Reference: https://go.dev/blog/error-handling-and-go
2-) Error Type is a built in type nad universe block
3-) Panic is runtime error

*/
