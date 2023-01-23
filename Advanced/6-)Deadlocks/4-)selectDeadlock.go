package main

import "fmt"

func main() {

	mychan1 := make(chan int)
	mychan2 := make(chan int)

	go func(c1 chan int) {
		c1 <- 1
	}(mychan1)

	go func(c2 chan int) {
		c2 <- 2
	}(mychan2)

	for {
		select {
		case c1 := <-mychan1:
			fmt.Println(c1)
		case c2 := <-mychan2:
			fmt.Println(c2)
		}

	}

}

/*Some Notes By Mustafa Karacabey

- Probably cause for still wait the deadlock cause of infinite for loop !.

*/
