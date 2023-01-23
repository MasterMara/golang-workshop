package main

import (
	"fmt"
	"time"
)

func main() {

	// 1.WAY Closure  Funtion
	for i := 0; i < 10; i++ {

		func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("*******************")

	//2.WAY With goroutine Closure Function
	for i := 0; i < 10; i++ { // Guess result ?
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 5)
}

/*Some Notes By Mustafa KARACABEY:

1-) Closure function will remember outer scope value because it gives reference which means memory address


*/
