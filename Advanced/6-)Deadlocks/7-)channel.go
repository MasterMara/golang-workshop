package main

import (
	"fmt"
	"time"
)

func main() {

	myChannel := make(chan string) //Create Channel

	go count("Fish", myChannel)

	//DEADLOCK !!!!!
	for { //Bu durumda goroutines bitmesine rağmen channelden hala data beklerse deadlock durumuna girer
		msg := <-myChannel //Channeldaki Data Buraya Geliyor
		fmt.Println(msg)
	}

	/*for msg := range c { //Deadlocktan bu şekildede manuel olmadan kurtulabiliriz
		fmt.Println(msg)
	}*/

}

func count(thing string, c chan string) {

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		c <- thing //Channelı Dolduruyor

	}

	close(c) // Deadlocktan kurtarır channelı kapatırki daha data gelmesin eğer receiversanız kesinlikle kapatmayın sender iseniz kapatın

	//Manuel olarak kapatmak yerine

}
