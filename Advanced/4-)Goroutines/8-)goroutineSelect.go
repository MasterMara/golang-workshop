package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every500ms"
			time.Sleep(time.Millisecond * 500)
		}

	}()

	go func() {
		for {
			c2 <- "every2second"
			time.Sleep(2 * time.Second)
		}

	}()

	for {
		select { // Hangi Kanaldaki Data Hazırsa onu Daha hızlı şekilde yazdıracaktır
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}

/*Some Notes By Mustafa KARACABEY:





 */
