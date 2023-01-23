package main

import "fmt"

func exampleChan1Usage(myChan chan string) {

	myChan <- "MUSTAFA"
}

func main() {

	chan1, chan2 := make(chan string), make(chan int)

	go exampleChan1Usage(chan1)

	select {

	case msg := <-chan1:
		fmt.Println("Chan 1 is okey", msg)
	case msg := <-chan2:
		fmt.Println("Chan2 data is okey", msg)

	}

}

/*Some Notes By Mustafa KARACABEY:

1-)  A select is only used with channels.
2-) A switch is used with concrete types.
3-) Note that a switch can also go over types for interfaces when used with the keyword .(type)



*/
