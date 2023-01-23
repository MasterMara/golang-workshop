package main

import (
	"fmt"
)

func main() {

	myUnbufChanngel := make(chan bool)

	go func() {

		<-myUnbufChanngel //Listen That Channel
		fmt.Println(myUnbufChanngel)
	}()

	myUnbufChanngel <- true //Send Data

}

/*Some Notes By Mustafa KARACABEY:

1-) https://www.velotio.com/engineering-blog/understanding-golang-channels


*/
